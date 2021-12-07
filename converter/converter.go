package converter

import (
	"strings"
)

// Html2MdConverter 数据转换器，会将指定的数据转换为markdown字节流
type Html2MdConverter interface {
	// Convert 将指定的数据转换为markdown并返回，
	// 如果转换中出错，则会返回error
	Convert(data []byte) ([]byte, error)
}

type ContentType int

const (
	HTML ContentType = iota
)

func NewConverter(contentType ContentType, option *Option) Html2MdConverter {
	switch contentType {
	case HTML:
		return &CommonHtml2MdConverter{Option: option}
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
