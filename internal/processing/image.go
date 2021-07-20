package processing

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/go-component/markdown-proc/internal/baseutil"
	"github.com/go-component/markdown-proc/internal/fileutil"
	"github.com/go-component/markdown-proc/internal/types"
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

func (i *Image) imagePathFormat(filename string) string {

	return filepath.Join(i.parseImageDir(), filename)
}

func (i *Image) imageRelativePathFormat(filename string) string {

	return filepath.Join(filepath.Base(i.parseImageDir()), filename)
}

func (i *Image) outputPathFormat() string {

	return filepath.Join(i.Command.Output, filepath.Base(i.Command.Filename))
}

func (i *Image) parseImageUrl() (list []types.UrlAddress, err error) {

	handler, err := os.Open(i.Command.Filename)
	if err != nil {
		return list, err
	}
	defer handler.Close()

	b, err := io.ReadAll(handler)

	if err != nil {
		return list, err
	}

	reg := regexp.MustCompile("!\\[.*\\]\\((http.*)\\)")

	result := reg.FindAllStringSubmatch(string(b), -1)

	for _, v := range result {

		md5Str := baseutil.Md5(v[1])
		list = append(list,types.UrlAddress{
			Value:       v[1],
			Md5:         md5Str,
			Md5Filename: fmt.Sprintf("%s%s", md5Str, fileutil.Ext(v[1])),
		})
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

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
		},
	}

	for _,urlAddress := range list {
		i.crawl(eg, client, urlAddress)
	}

	if err = eg.Wait(); err != nil {
		return err
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

	bf := bufio.NewWriter(handler)
	_, err = bf.Write(content)

	if err != nil{
		return err
	}

	err = bf.Flush()

	if err == nil{
		log.Printf("processing success: %s", outputPath)
	}

	return err
}

func (i *Image) replaceImagePath(list []types.UrlAddress) (b []byte, err error) {

	handler, err := os.Open(i.Command.Filename)
	if err != nil {
		return b, err
	}
	defer handler.Close()

	b, err = io.ReadAll(handler)

	if err != nil {
		return b, err
	}

	for _, urlAddress := range list {
		imagePath := i.imageRelativePathFormat(urlAddress.Md5Filename)
		b = bytes.Replace(b, []byte(urlAddress.Value), []byte(imagePath), -1)
	}

	return b, nil
}

func (i *Image) crawl(eg *errgroup.Group, client *http.Client, urlAddress types.UrlAddress) {

	eg.Go(func() error {

		req, err := http.NewRequest("GET", urlAddress.Value, nil)
		response, err := client.Do(req)
		if err != nil {
			return err
		}

		defer response.Body.Close()

		if response.StatusCode != 200 {
			log.Printf("status code error of image: %s, statusCode: %d \n", urlAddress.Value, response.StatusCode)
		}

		imagePath := i.imagePathFormat(urlAddress.Md5Filename)

		imageHandler, err := os.OpenFile(imagePath, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			return err
		}

		defer imageHandler.Close()

		_, err = io.Copy(imageHandler, response.Body)

		return err
	})

}
