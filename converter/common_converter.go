package converter

import (
	"fmt"
	"github.com/yangsx95/markdown-tools/util"
	"os"
	"path/filepath"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

type CommonHtml2MdConverter struct {
	Option *Option
}

// Convert 采用三方库 https://github.com/JohannesKaufmann/html-to-markdown 进行转换
func (hc *CommonHtml2MdConverter) Convert(data []byte) (markdown []byte, err error) {
	// 解析html为markdown
	converter := md.NewConverter("", true, nil)
	if !hc.Option.SaveImgToLocal {
		return converter.ConvertBytes(data)
	}
	converter.AddRules(md.Rule{
		Filter: []string{"img"},
		Replacement: func(content string, selec *goquery.Selection, options *md.Options) *string {
			src, srcExist := selec.Attr("src")
			if !srcExist {
				return md.String("![无效的图片]()")
			}
			alt, altExist := selec.Attr("alt")
			if !altExist {
				alt = util.RandStr(10)
			}

			var f *os.File
			if f, err = util.DownloadFile(src, filepath.Join(hc.Option.MarkdownPath, hc.Option.ImgRelativePath), util.UrlStrategy); err != nil {
				fmt.Printf("下载文件出错 src=%v，错误信息 %v\n", src, err)
				return md.String("![" + alt + "](" + src + ")")
			}

			return md.String("![" + alt + "](" + filepath.Join(hc.Option.ImgRelativePath, filepath.Base(f.Name())) + ")")
		}})
	return converter.ConvertBytes(data)
}
