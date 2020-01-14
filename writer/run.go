package writer

import (
	"log"
	"path/filepath"

	"github.com/urfave/cli/v2"

	"github.com/cittie/xwriter/util"
)

var (
	data string
	src  string
	tar  string
	ext  string
)

func Run(c *cli.Context) error {
	data, src, tar, ext = c.String(util.Data), c.String(util.Source), c.String(util.Target), c.String(util.Extension)

	return run()
}

func run() error {
	// 遍历data
	s, err := util.GetSortedFilenames(data, ext)
	if err != nil {
		return err
	}

	for _, sName := range s {
		if err := AddData(filepath.Join(data, sName)); err != nil {
			return err
		}
	}

	// 遍历source
	for _, xe := range xlsExts {
		names, err := util.GetSortedFilenames(src, xe)
		if err != nil {
			return err
		}
		for _, tName := range names {
			if err := AddSource(filepath.Join(src, tName)); err != nil {
				return err
			}
		}
	}

	// link and add
	for fn, m := range dataMap {
		f, ok := sources[fn]
		if !ok {
			log.Printf("%v, no xls file %s", ErrFilenameInvalid, fn)
			continue
		}

		for sn, d := range m {
			sh, ok := f.f.Sheet[sn]
			if !ok {
				log.Printf("%v, no sheet %s", ErrInvalidSheet, sn)
				continue
			}
			for _, dd := range d {
				if err := NewWriteObject(dd, sh).Write2xls(); err != nil {
					return err
				}
			}
		}

		// 写回去
		if f != nil && f.f != nil {
			err := f.f.Save(filepath.Join(tar, fn+f.ext))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
