package provider

import (
	"encoding/json"
	"go-markdown-crawler/util"
	"io/ioutil"
	"net/http"
	"strconv"
)

type wordpressProvider struct {
	protocol string
	domain   string
	basepath string
}

func NewWordpressProvider(protocol string, domain string) Provider {
	return &wordpressProvider{protocol, domain, "/wp-json/wp/v2"}
}

func (wp *wordpressProvider) AuthWithUsernameAndPassword(username string, password string) {

}

type wordpressPageIterator struct {
	protocol string
	domain   string
	baseurl  string
	path     string
	// 总条数
	totalCount int
	// 当前读取到的条数
	currentIndex int
}

func newWordpressPageIterator(protocol string, domain string, baseurl string, path string) (wordpressPageIterator, error) {
	u := util.GenFullRequestUrl(protocol, domain, baseurl, path, map[string]string{"per_page": "1"})
	resp, err := http.Get(u)
	if err != nil {
		return wordpressPageIterator{}, err
	}
	defer util.FastClose(resp.Body)
	// 总条数
	totalCount, _ := strconv.Atoi(resp.Header.Get("X-WP-Total"))
	return wordpressPageIterator{protocol, domain, baseurl, path, totalCount, 0}, nil
}

func (wpi *wordpressPageIterator) Next() ([]byte, error) {
	return wpi.NextMulti(1)
}

func (wpi *wordpressPageIterator) NextMulti(count int) ([]byte, error) {
	u := util.GenFullRequestUrl(wpi.protocol, wpi.domain, wpi.baseurl, wpi.path,
		map[string]string{"offset": strconv.Itoa(wpi.currentIndex), "per_page": strconv.Itoa(count)})
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer util.FastClose(resp.Body)
	bys, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	wpi.currentIndex = wpi.currentIndex + count
	return bys, nil
}

func (wpi *wordpressPageIterator) HasNext() bool {
	return wpi.currentIndex < wpi.totalCount
}

func (wp *wordpressProvider) GetAllCategories() (Categories, error) {
	categories := make(Categories, 0)
	iterator, err := newWordpressPageIterator(wp.protocol, wp.domain, wp.basepath, "/categories")
	if err != nil {
		return nil, err
	}
	for iterator.HasNext() {
		bys, err := iterator.NextMulti(10)
		if err != nil {
			return nil, err
		}
		result := make([]wordpressCategory, 0)
		err = json.Unmarshal(bys, &result)
		if err != nil {
			return nil, err
		}
		for _, v := range result {
			categories = append(categories, Category{v.ID, v.Name, v.Slug, v.Description, v.Link, v.Parent})
		}
	}
	return categories, nil
}

func (wp *wordpressProvider) GetAllTags() (Tags, error) {
	tags := make(Tags, 0)
	iterator, err := newWordpressPageIterator(wp.protocol, wp.domain, wp.basepath, "/tags")
	if err != nil {
		return nil, err
	}
	for iterator.HasNext() {
		bys, err := iterator.NextMulti(10)
		if err != nil {
			return nil, err
		}
		result := make([]wordpressTag, 0)
		err = json.Unmarshal(bys, &result)
		if err != nil {
			return nil, err
		}
		for _, v := range result {
			tags = append(tags, Tag{v.ID, v.Description, v.Name})
		}
	}
	return tags, nil
}

func (wp wordpressProvider) GetArticleIterator() ArticleIterator {
	return nil
}

type wordpressCategory struct {
	ID          int           `json:"id"`
	Count       int           `json:"count"`
	Description string        `json:"description"`
	Link        string        `json:"link"`
	Name        string        `json:"name"`
	Slug        string        `json:"slug"`
	Taxonomy    string        `json:"taxonomy"`
	Parent      int           `json:"parent"`
	Meta        []interface{} `json:"meta"`
}

type wordpressTag struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

// wordpressArticle 文章
type wordpressArticle struct {
	Id         int              `json:"id"`
	Date       string           `json:"date"`
	Status     string           `json:"status"`
	Link       string           `json:"link"`
	Title      wordpressTitle   `json:"title"`
	Content    wordpressContent `json:"content"`
	Categories []int            `json:"categories"`
	Tags       []interface{}    `json:"tags"`
}

type wordpressTitle struct {
	Rendered string `json:"rendered"`
}
type wordpressContent struct {
	Rendered  string `json:"rendered"`
	Protected bool   `json:"protected"`
}

type WordpressArticleIterator struct {
}
