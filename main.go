package main

import (
	"flag"
	"github.com/go-component/markdown-proc/internal"
	"github.com/go-component/markdown-proc/internal/conf"
	"github.com/go-component/markdown-proc/internal/option"
	"log"
)

var (
	mode     int
	output   string
	filename string
)

const (
	defaultMode = conf.Image
)

func init() {
	flag.IntVar(&mode, "m", defaultMode, `processing mode 
0: image 
1: word

default 0
`)
	flag.StringVar(&output, "o", "", "output path")
	flag.StringVar(&filename, "f", "", "filepath of markdown")
}

func main() {
	flag.Parse()

	if err := conf.CheckMode(mode); err != nil {
		log.Fatalln(err)
	}

	var opt option.CommandOption

	switch mode {
	case conf.Image:
		opt = option.WithImageModeOption()
	case conf.Word:
		opt = option.WithWordModeOption()
	}

	commandOption, err := option.NewCommandOption(filename, output, opt)

	if err != nil {
		log.Fatalln(err)
	}

	if err = internal.Run(commandOption); err != nil {
		log.Fatalln(err)
	}

}
