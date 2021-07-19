package main

import (
	"flag"
	"fmt"
	"github.com/go-component/markdown-proc/conf"
	"github.com/go-component/markdown-proc/option"
	"log"
)

var (
	mode         int
	output       string
	imageDirname string
	filename     string
)

const (
	defaultMode = conf.Image
)

func init() {
	flag.IntVar(&mode, "m", defaultMode, `processing mode 
1:image 
2:word`)
	flag.StringVar(&output, "o", "", "output dir")
	flag.StringVar(&imageDirname, "d", "", "dirname of image, default same as markdown name")
	flag.StringVar(&filename, "f", "", "filename of markdown")
}

func main() {
	flag.Parse()

	if err := conf.CheckMode(mode); err != nil {
		log.Fatalln(err)
	}

	var opt option.CommandOption

	switch mode {
	case conf.Image:
		opt = option.WithImageModeOption(imageDirname)
	case conf.Word:
		opt = option.WithWordModeOption()
	}

	commandOption, err := option.NewCommandOption(filename, output, opt)

	if err != nil{
		log.Fatalln(err)
	}

	fmt.Println(commandOption)

}
