package parser

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cittie/xwriter/util"
)

func TestXlsDataCell_GetSet(t *testing.T) {
	id := "TestID"
	seq := 0
	context := "whatever"

	c := new(XlsDataCell)
	c.Set(id, seq, context)
	i, s, ctx := c.Get()

	assert.Equal(t, id, i)
	assert.Equal(t, seq, s)
	assert.Equal(t, context, ctx)
}

func TestXlsDataParser(t *testing.T) {
	xpd := new(XlsDataParser)
	xpd.SetInfo("jws", "develop", "NewTest", "TestSheet", 3)
	xpd.SetEncodeType(util.JSON)

	row := make([]IXlsDataCell, 10)
	for i := 0; i < 10; i++ {
		c := new(XlsDataCell)
		c.Set(strconv.Itoa(i), i, strconv.Itoa(i))
		row[i] = c
	}

	xpd.AddRow(row)

	xpd.ToXlsDataFile("../tests/no.json")
}
