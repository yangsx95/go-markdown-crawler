package exporter

import (
	"fmt"
	"github.com/yangsx95/md-crawler/provider"
	"testing"
)

func TestCategoryFileExporter_export(t *testing.T) {
	p := provider.NewWordpressProvider("http", "yangsx95.com")
	exporter, err := NewFileExporter(p, "./test")
	if err != nil {
		panic(err)
	}
	err = exporter.Export()
	if err != nil {
		panic(err)
	}
	fmt.Println(exporter.categoryIdFilePathMap)
}
