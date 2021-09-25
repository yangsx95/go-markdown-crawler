package provider

import (
	"encoding/json"
	"go-markdown-crawler/util"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
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
	basepath string
	path     string
	// 总条数
	totalCount int
	// 每页条数
	perPage int
	// 当前读取到的页数
	currentPage int
}

func newWordpressPageIterator(protocol string, domain string, basepath string, path string, perPage int) (wordpressPageIterator, error) {
	u := util.GenFullRequestUrl(protocol, domain, basepath, path, map[string]string{"per_page": "1"})
	resp, err := http.Get(u)
	if err != nil {
		return wordpressPageIterator{}, err
	}
	defer util.FastClose(resp.Body)
	// 总条数
	totalCount, _ := strconv.Atoi(resp.Header.Get("X-WP-Total"))
	return wordpressPageIterator{protocol, domain, basepath, path, totalCount, perPage, 0}, nil
}

func (wpi *wordpressPageIterator) NextPage() ([]byte, error) {
	wpi.currentPage = wpi.currentPage + 1
	u := util.GenFullRequestUrl(wpi.protocol, wpi.domain, wpi.basepath, wpi.path,
		map[string]string{
			"per_page": strconv.Itoa(wpi.perPage),
			"page":     strconv.Itoa(wpi.currentPage),
		})
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer util.FastClose(resp.Body)
	bys, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bys, nil
}

func (wpi *wordpressPageIterator) HasNext() bool {
	return wpi.currentPage < wpi.totalCount/wpi.perPage
}

func (wp *wordpressProvider) GetAllCategories() (Categories, error) {
	categories := make(Categories, 0)
	iterator, err := newWordpressPageIterator(wp.protocol, wp.domain, wp.basepath, "/categories", 10)
	if err != nil {
		return nil, err
	}
	for iterator.HasNext() {
		bys, err := iterator.NextPage()
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
	iterator, err := newWordpressPageIterator(wp.protocol, wp.domain, wp.basepath, "/tags", 100)
	if err != nil {
		return nil, err
	}
	for iterator.HasNext() {
		bys, err := iterator.NextPage()
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

func (wp wordpressProvider) GetArticleIterator() (ArticleIterator, error) {
	return NewWordpressArticleIterator(wp.protocol, wp.domain, wp.basepath)
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
	Tags       []int            `json:"tags"`
}

type wordpressTitle struct {
	Rendered string `json:"rendered"`
}
type wordpressContent struct {
	Rendered  string `json:"rendered"`
	Protected bool   `json:"protected"`
}

type WordpressArticleIterator struct {
	iterator *wordpressPageIterator
}

func NewWordpressArticleIterator(protocol string, domain string, basepath string) (WordpressArticleIterator, error) {
	iterator, err := newWordpressPageIterator(protocol, domain, basepath, "/posts", 1)
	if err != nil {
		return WordpressArticleIterator{}, err
	}
	return WordpressArticleIterator{iterator: &iterator}, nil
}

func (wai WordpressArticleIterator) HasNext() bool {
	return wai.iterator.HasNext()
}

func (wai WordpressArticleIterator) Next() (Article, error) {
	bys, err := wai.iterator.NextPage()
	if err != nil {
		return Article{}, err
	}
	wArticles := make([]wordpressArticle, 0)
	err = json.Unmarshal(bys, &wArticles)
	if err != nil {
		return Article{}, err
	}
	if len(wArticles) == 0 {
		return Article{}, err
	}

	articles := make(Articles, 0)
	for _, wArticle := range wArticles {
		aDate, _ := time.Parse("2006-01-02T15:04:05", wArticle.Date)
		article := Article{wArticle.Id,
			aDate, wArticle.Link,
			wArticle.Title.Rendered,
			wArticle.Content.Rendered,
			"html",
			wArticle.Categories,
			wArticle.Tags}
		articles = append(articles, article)
	}
	return articles[0], nil
}
