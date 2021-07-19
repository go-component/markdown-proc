package main

import (
	"flag"
	"github.com/go-component/markdown-proc/conf"
	"github.com/go-component/markdown-proc/option"
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
1: image 
2: word`)
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

	err = commandOption.Processing.Process()

	if err != nil {
		log.Fatalln(err)
	}

}
