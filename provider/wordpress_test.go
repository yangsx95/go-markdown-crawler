package provider

import (
	"fmt"
	"testing"
)

func TestNewWordpressProvider(t *testing.T) {

}

func TestWordpressProvider_GetAllCategories(t *testing.T) {
	provider := NewWordpressProvider("http", "yangsx95.com")
	cts, _ := provider.GetAllCategories()
	fmt.Println(cts)
}

func TestWordpressProvider_GetAllTags(t *testing.T) {
	provider := NewWordpressProvider("http", "yangsx95.com")
	tgs, _ := provider.GetAllTags()
	fmt.Println(tgs)
}
