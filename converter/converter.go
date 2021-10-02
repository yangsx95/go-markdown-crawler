package converter

import (
	"strings"
)

type Converter interface {

	// Convert 将指定的数据转换为markdown并返回，
	// 如果转换中出错，则会返回error
	Convert(data []byte) ([]byte, error)
}

type ContentType int

const (
	HTML ContentType = iota
)

func NewConverter(contentType ContentType) Converter {
	switch contentType {
	case HTML:
		return &HtmlConverter{}
	}
	return nil
}

func GetContentTypeByString(contentType string) ContentType {
	ct := strings.ToUpper(contentType)
	switch ct {
	case "HTML":
		return HTML
	default:
		return -1
	}
}
