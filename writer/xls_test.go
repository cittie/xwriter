package writer

import (
	"strconv"
	"testing"

	"github.com/tealeg/xlsx"
)

func newXls() {
	f := xlsx.NewFile()
	sh, _ := f.AddSheet("TestSheet")

	// 3行空行
	for i := 0; i < 3; i++ {
		sh.AddRow()
	}

	// 第4行写入0-2
	row := sh.AddRow()

	for i := 0; i < 9; i++ {
		c := row.AddCell()
		c.SetString(strconv.Itoa(i % 3))
	}

	// 3行空行
	for i := 0; i < 3; i++ {
		sh.AddRow()
	}

	f.Save("../tests/s/NewTest.xlsx")
}

func TestRun(t *testing.T) {
	newXls()

	data = "../tests/d"
	src = "../tests/s"
	tar = "../tests/t"
	ext = "json"

	err := run()
	t.Logf("%v", err)
}
