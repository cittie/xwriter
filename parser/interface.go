package parser

// IXlsDataCell 单个格子
type IXlsDataCell interface {
	// Set 赋值此单元格对应的id，对应的序列（从0开始），以及内容
	Set(id string, seq int, content string)

	// Get 获取此单元格对应的id，对应的序列（从0开始），以及内容
	Get() (id string, seq int, content string)
}

// IXlsDataParser 聚合写入数据
type IXlsDataParser interface {
	// SetInfo 设置对应信息
	// project, branch目前版本不使用，只是用于记录
	// xFileBaseName, xSheetName对大小写敏感，务必填写完全匹配的名称
	// xFileBaseName 只需要填文件basename，譬如对应Item.xlsm，填"Item"即可
	SetInfo(project, branch, xFileBaseName, xSheetName string, idRowIdx int)

	// SetEncodeType 设置编码，当前仅json和bin两个选项可用
	SetEncodeType(typ string)

	// AddRow 添加一行数据，无需排序，如果重复，后面的数据会覆盖前面数据
	// 这里的行并不对应实际的数据行，而是对应这行哪些列需要写入数据，因此不需要插入空格子
	// 譬如，添加第2格和第7格，会生成长度7的行，仅2和7有内容
	AddRow(cells []IXlsDataCell) error

	// ToXlsDataFile 输出到文件
	ToXlsDataFile(filename string) error
}
