package writer

import (
	"errors"
)

var (
	xlsExts = []string{"xlsx", "xls", "xlsm"} // 这里定义了所有xls文件的扩展名，用于遍历xls文件
)

var (
	ErrFilenameInvalid     = errors.New("invalid filename")
	ErrDuplicateXlsName    = errors.New("duplicate xls filename")
	ErrTargetFileNotExists = errors.New("target file not exists")  // 文件不存在
	ErrTargetSheetNotExits = errors.New("target sheet not exists") // 表不存在
	ErrInvalidSheet        = errors.New("invalid sheet")           // 表结构非法
	ErrInvalidCellData     = errors.New("invalid cell data")       // 单元格数据有误
)
