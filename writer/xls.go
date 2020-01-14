package writer

import (
	"strings"

	"github.com/tealeg/xlsx"

	"github.com/cittie/xwriter/protocol/pb"
)

type WriteObject struct {
	source *pb.XlsData
	sheet  *xlsx.Sheet
}

// NewWriteObject 创建一个写入对象，包括源数据和写入表信息
func NewWriteObject(src *pb.XlsData, sheet *xlsx.Sheet) *WriteObject {
	if src == nil || sheet == nil {
		return nil
	}

	return &WriteObject{
		source: src,
		sheet:  sheet,
	}
}

// removeBlankRows 删除空行
func (wb *WriteObject) removeBlankRows() error {
	length := len(wb.sheet.Rows)

	for l := length - 1; l > int(wb.source.Info.IdRowIdx); l-- {
		// 从最后一行，到定义ID的行为止，检查该行是不是空的
		// 如果cell个数不为0，遍历每个cell
		if len(wb.sheet.Rows[l].Cells) > 0 {
			isBlank := true
			for i := 0; i < len(wb.sheet.Rows[l].Cells); i++ {
				if wb.sheet.Rows[l].Cells[i].Value != "" {
					isBlank = false
					break
				}
			}
			if isBlank {
				err := wb.sheet.RemoveRowAtIndex(l)
				if err != nil {
					return err
				}
			} else {
				break
			}
		} else {
			err := wb.sheet.RemoveRowAtIndex(l)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// write2xls 将一张表的数据写入
func (wb *WriteObject) Write2xls() error {
	err := wb.removeBlankRows()
	if err != nil {
		return err
	}

	// 表头
	if int(wb.source.Info.IdRowIdx) >= len(wb.sheet.Rows) {
		return ErrInvalidSheet
	}

	idRow := wb.sheet.Rows[wb.source.Info.IdRowIdx]
	rowLen := len(idRow.Cells)

	// 建立id和col对应的map
	idMap := make(map[string][]int)
	for idx, cell := range idRow.Cells {
		if cell == nil || strings.TrimSpace(cell.Value) == "" {
			continue
		}
		v := strings.TrimSpace(cell.Value)
		idMap[v] = append(idMap[v], idx)
	}

	// 按行写入
	for i := 0; i < len(wb.source.Data); i++ {
		row := wb.sheet.AddRow()
		rowData := wb.source.Data[i]
		row.Cells = make([]*xlsx.Cell, rowLen) // 按最大长度先建好
		for _, cellData := range rowData.Data {
			// 先检查有没有对应的id
			seq, ok := idMap[cellData.Id]
			if !ok {
				continue
			}

			// 再检查长度
			if int(cellData.IdIdx) >= len(seq) {
				continue
			}

			// 写入内容
			cell := xlsx.NewCell(row)
			cell.SetString(cellData.Content)

			row.Cells[seq[cellData.IdIdx]] = cell
		}

		// 需要给空格子赋值，否则会panic
		for i := 0; i < len(row.Cells); i++ {
			if row.Cells[i] == nil {
				row.Cells[i] = xlsx.NewCell(row)
			}
		}
	}

	return nil
}
