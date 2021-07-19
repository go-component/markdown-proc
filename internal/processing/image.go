package processing

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/go-component/markdown-proc/internal/fileutil"
	"github.com/go-component/markdown-proc/types"
	"golang.org/x/sync/errgroup"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

type Image struct {
	Command *types.Command
}


func (i *Image) parseImageDir() string {

	return filepath.Join(i.Command.Output, i.Command.ImageDirname)
}

func (i *Image) tryCreateDir() error {

	var err error

	err = fileutil.MkdirAll(i.Command.Output)
	if err != nil {
		return err
	}

	err = fileutil.MkdirAll(i.parseImageDir())

	return err
}

func (i *Image) imagePathFormat(url string, index int) string {
	ext := fileutil.Ext(url)

	return filepath.Join(i.parseImageDir(), fmt.Sprintf("%d%s", index, ext))
}

func (i *Image) imageRelativePathFormat(url string, index int) string {

	ext := fileutil.Ext(url)

	return filepath.Join(filepath.Base(i.parseImageDir()), fmt.Sprintf("%d%s", index, ext))
}

func (i *Image) outputPathFormat() string {

	return filepath.Join(i.Command.Output, filepath.Base(i.Command.Filename))
}

func (i *Image) parseImageUrl() (list []string, err error) {

	handler, err := os.Open(i.Command.Filename)
	if err != nil {
		return list, err
	}
	defer handler.Close()

	b, err := io.ReadAll(handler)

	if err != nil {
		return list, err
	}

	reg := regexp.MustCompile("!\\[\\]\\((http.*)\\)")

	result := reg.FindAllStringSubmatch(string(b), -1)

	for _, v := range result {
		list = append(list, v[1])
	}

	return list, nil
}

func (i *Image) Process() error {

	var err error

	if err = i.tryCreateDir(); err != nil {
		return err
	}

	list, err := i.parseImageUrl()

	if err != nil {
		return err
	}

	eg := new(errgroup.Group)

	for k, url := range list {
		i.crawl(eg, url, k)
	}

	if err = eg.Wait(); err != nil {
		return err
	}

	for k, url := range list {
		i.crawl(eg, url, k)
	}

	b, err := i.replaceImagePath(list)

	if err != nil {
		return err
	}

	return i.write(b)
}

func (i *Image) write(content []byte) error {

	outputPath := i.outputPathFormat()
	handler, err := os.OpenFile(outputPath, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer handler.Close()

	_, err = bufio.NewWriter(handler).Write(content)

	log.Printf("processing success: %s", outputPath)

	return err
}

func (i *Image) replaceImagePath(list []string) (b []byte, err error) {

	handler, err := os.Open(i.Command.Filename)
	if err != nil {
		return b, err
	}
	defer handler.Close()

	b, err = io.ReadAll(handler)

	if err != nil {
		return b, err
	}

	for k, url := range list {
		imagePath := i.imageRelativePathFormat(url, k+1)

		b = bytes.Replace(b, []byte(url), []byte(imagePath), -1)
	}

	return b, nil
}

func (i *Image) crawl(eg *errgroup.Group, url string, k int) {

	eg.Go(func() error {

		response, err := http.Get(url)
		if err != nil {
			return err
		}

		defer response.Body.Close()

		if response.StatusCode != 200 {
			log.Printf("status code error of image: %s, statusCode: %d, index: %d\n", url, response.StatusCode, k+1)
		}

		imagePath := i.imagePathFormat(url, k+1)

		imageHandler, err := os.OpenFile(imagePath, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			return err
		}

		defer imageHandler.Close()

		_, err = io.Copy(imageHandler, response.Body)

		return err
	})

}
