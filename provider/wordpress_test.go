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

func TestWordpressProvider_GetArticleIterator(t *testing.T) {
	provider := NewWordpressProvider("http", "yangsx95.com")
	iterator, err := provider.GetArticleIterator()
	if err != nil {
		panic(err)
	}
	if iterator.HasNext() {
		multi, err := iterator.NextMulti(3)
		if err != nil {
			panic(err)
		}
		fmt.Println(multi)
		if len(multi) != 3 {
			_ = fmt.Errorf("返回文章数量%v不是预期值%v\n", len(multi), 3)
		}
	}

}
