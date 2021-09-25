package exporter

import (
	"fmt"
	"go-markdown-crawler/provider"
	"testing"
)

func TestCategoryFileExporter_export(t *testing.T) {
	p := provider.NewWordpressProvider("http", "yangsx95.com")
	exporter, err := NewFileExporter(p, "./test")
	if err != nil {
		panic(err)
	}
	err = exporter.export()
	if err != nil {
		panic(err)
	}
	fmt.Println(exporter.categoryIdFilePathMap)
}
