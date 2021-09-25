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
	fmt.Println(len(cts))
}

func TestWordpressProvider_GetAllTags(t *testing.T) {
	provider := NewWordpressProvider("http", "yangsx95.com")
	tgs, _ := provider.GetAllTags()
	fmt.Println(tgs)
}

func TestWordpressProvider_GetArticleIterator(t *testing.T) {
	provider := NewWordpressProvider("http", "yangsx95.com")
	iterator, err := provider.GetArticleIterator()
	if err != nil {
		panic(err)
	}
	if iterator.HasNext() {
		atc, err := iterator.Next()
		if err != nil {
			panic(err)
		}
		fmt.Println(atc)
	}

}
