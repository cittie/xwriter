package main

import (
	"errors"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/cittie/xwriter/util"
	"github.com/cittie/xwriter/writer"
)

func main() {
	app := &cli.App{
		Name:  "xls writer",
		Usage: "write data to xls files",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        util.Data,
				Aliases:     []string{"d"},
				Value:       util.DefaultPath,
				Usage:       "set dir name of data",
				DefaultText: util.DefaultPath,
			},
			&cli.StringFlag{
				Name:        util.Source,
				Aliases:     []string{"s"},
				Value:       util.DefaultPath,
				Usage:       "set dir name of source xls files",
				DefaultText: util.DefaultPath,
			},
			&cli.StringFlag{
				Name:        util.Target,
				Value:       util.DefaultPath,
				Aliases:     []string{"t"},
				Usage:       "set dir name of where appended xls files to save",
				DefaultText: util.DefaultPath,
			},
			&cli.StringFlag{
				Name:        util.Extension,
				Value:       util.JSON,
				Aliases:     []string{"e"},
				Usage:       "extend name of source data files, json or bin",
				DefaultText: util.JSON,
			},
		},
		Action: func(c *cli.Context) error {
			// 检查下目录是否存在
			for _, param := range []string{util.Data, util.Source, util.Target} {
				if !util.FileExists(c.String(param)) {
					return errors.New("dir not exists")
				}
			}

			return writer.Run(c)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
