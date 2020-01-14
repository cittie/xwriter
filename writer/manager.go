package writer

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/tealeg/xlsx"

	"github.com/cittie/xwriter/protocol/pb"
	"github.com/cittie/xwriter/util"
)

var (
	dataMap    map[string]map[string][]*pb.XlsData // [target file][target sheet][]*xlsData
	sources    map[string]*XlsTarget               // [basename]XlsTarget
	duplicates map[string]struct{}                 // [basename]{}
)

func init() {
	dataMap = make(map[string]map[string][]*pb.XlsData)
	sources = make(map[string]*XlsTarget)
	duplicates = make(map[string]struct{})
}

type XlsTarget struct {
	f   *xlsx.File
	ext string
}

// AddData 添加数据文件
func AddData(filename string) error {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// 解码
	xd := new(pb.XlsData)

	switch ext {
	case util.JSON:
		err = json.Unmarshal(raw, xd)
	case util.BIN:
		err = proto.Unmarshal(raw, xd)
	default:
		return nil
	}

	if err != nil {
		return err
	}

	// load
	fn := xd.Info.XlsFileName
	_, ok := dataMap[fn]
	if !ok {
		dataMap[fn] = make(map[string][]*pb.XlsData)
	}

	sn := xd.Info.XlsSheetName
	_, ok = dataMap[fn][sn]
	if !ok {
		dataMap[fn][sn] = make([]*pb.XlsData, 0, 4)
	}
	dataMap[fn][sn] = append(dataMap[fn][sn], xd)

	return nil
}

// AddSource 添加xls源文件
func AddSource(filename string) error {
	f, err := xlsx.OpenFile(filename)
	if err != nil {
		return err
	}

	base, ext := filepath.Base(filename), filepath.Ext(filename)
	if base == "." {
		return nil
	}

	// base 再去掉扩展名，作为唯一标示
	bBase := strings.TrimSuffix(base, ext)
	if _, ok := duplicates[bBase]; ok {
		return ErrDuplicateXlsName
	}

	duplicates[bBase] = struct{}{}
	sources[bBase] = &XlsTarget{
		f:   f,
		ext: ext,
	}

	return nil
}
