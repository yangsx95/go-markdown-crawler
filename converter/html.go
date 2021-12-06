package converter

import md "github.com/JohannesKaufmann/html-to-markdown"

type HtmlConverter struct {
}

// Convert 采用三方库 https://github.com/JohannesKaufmann/html-to-markdown 进行转换
func (hc *HtmlConverter) Convert(data []byte) (markdown []byte, err error) {
	// 解析html为markdown
	converter := md.NewConverter("", true, nil)
	return converter.ConvertBytes(data)
}
