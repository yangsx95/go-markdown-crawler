package util

import (
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// GenFullRequestUrl 根据路径信息，生成全路径
func GenFullRequestUrl(protocal string, domain string, basepath string, path string, query map[string]string) string {
	ks := make([]string, 0)
	var rawQuery string
	if query != nil && len(query) > 0 {
		for k, v := range query {
			ks = append(ks, k, "=", v, "&")
		}
		rawQuery = strings.Join(ks[:len(ks)-1], "")
	}

	u := &url.URL{
		Scheme:   protocal,
		Host:     domain,
		Path:     filepath.Join(basepath, path),
		RawQuery: rawQuery,
	}
	return u.String()
}

// FastClose 快速关闭Closer，并忽略err
func FastClose(closer io.Closer) {
	_ = closer.Close()
}

// PageReaderIterator 分页读取迭代器
type PageReaderIterator interface {
	// NextPage 读取下一页
	NextPage() ([]byte, error)
	// HasNext 是否有下一页面
	HasNext() bool
}

type FileNameStrategy func(url string, resp *http.Response) (filename string, err error)

var ContentDispositionStrategy FileNameStrategy = func(url string, resp *http.Response) (filename string, err error) {
	var cdstr string
	if cdstr = resp.Header.Get("Content-Disposition"); cdstr == "" {
		return filename, errors.New("未找到Content-Disposition头")
	}
	var params map[string]string
	if _, params, err = mime.ParseMediaType(cdstr); err != nil {
		return
	}
	var ok bool
	if filename, ok = params["filename"]; !ok {
		return filename, errors.New("未找到文件名称")
	}
	return
}

var UrlStrategy FileNameStrategy = func(u string, resp *http.Response) (filename string, err error) {
	var uu *url.URL
	if uu, err = url.Parse(u); err != nil {
		return
	}
	filename = filepath.Base(uu.Path)
	return
}

func DownloadFile(url, path string, strategy FileNameStrategy) (f *os.File, err error) {
	var resp *http.Response
	if resp, err = http.Get(url); err != nil {
		return
	}
	defer func() { _ = resp.Body.Close() }()

	// 获取文件存储路径
	var filename string
	if filename, err = strategy(url, resp); err != nil {
		return
	}
	_ = os.MkdirAll(path, os.ModePerm)
	if f, err = os.Create(filepath.Join(path, strings.TrimSpace(filename))); err != nil {
		fmt.Println(err)
		return
	}
	if _, err = io.Copy(f, resp.Body); err != nil {
		return
	}
	return
}
