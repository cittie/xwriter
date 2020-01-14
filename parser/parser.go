package parser

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"

	"github.com/cittie/xwriter/protocol/pb"
	"github.com/cittie/xwriter/util"
)

var (
	ErrNilData           = errors.New("parser data is nil")  // 对应没有初始化内置data的情况
	ErrInvalidEncodeType = errors.New("invalid encode type") // encode类型不支持
)

type XlsDataParser struct {
	data       *pb.XlsData
	encodeType string
}

type XlsDataCell struct {
	id      string
	seq     int
	content string
}

func (xdc *XlsDataCell) Set(id string, seq int, content string) {
	xdc.id = id
	xdc.seq = seq
	xdc.content = content
}

func (xdc *XlsDataCell) Get() (id string, seq int, content string) {
	return xdc.id, xdc.seq, xdc.content
}

// SetInfo 设置对应信息
// project, branch目前版本不使用，只是用于记录
// xFileBaseName, xSheetName对大小写敏感，务必填写完全匹配的名称
// xFileBaseName 只需要填文件basename，譬如对应Item.xlsm，填"Item"即可
func (xdp *XlsDataParser) SetInfo(project, branch, xFilename, xSheetName string, idRowIdx int) {
	if xdp.data == nil {
		xdp.data = new(pb.XlsData)
	}
	xdp.data.Info = &pb.DataInfo{
		Project:      project,
		Branch:       branch,
		XlsFileName:  xFilename,
		XlsSheetName: xSheetName,
		IdRowIdx:     int32(idRowIdx),
	}
}

func (xdp *XlsDataParser) SetEncodeType(typ string) {
	xdp.encodeType = typ
}

// AddRow 添加一行数据，无需排序，如果重复，后面的数据会覆盖前面数据
// 这里的行并不对应实际的数据行，而是对应这行哪些列需要写入数据，因此不需要插入空格子
// 譬如，添加第2格和第7格，会生成长度7的行，仅2和7有内容
func (xdp *XlsDataParser) AddRow(cells []IXlsDataCell) error {
	row := &pb.RowData{}
	for i := 0; i < len(cells); i++ {
		id, seq, content := cells[i].Get()
		c := &pb.CellData{
			Id:      id,
			IdIdx:   int32(seq),
			Content: content,
		}
		row.Data = append(row.Data, c)
	}
	xdp.data.Data = append(xdp.data.Data, row)
	return nil
}

func (xdp *XlsDataParser) ToXlsDataFile(filename string) error {
	if xdp.data == nil {
		return ErrNilData
	}

	var raw []byte
	var err error
	switch xdp.encodeType {
	case util.JSON:
		raw, err = json.MarshalIndent(xdp.data, "", "  ")
	case util.BIN:
		raw, err = proto.Marshal(xdp.data)
	default:
		return ErrInvalidEncodeType
	}

	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, raw, os.ModePerm)
}
