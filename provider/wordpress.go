package provider

import (
	"encoding/json"
	"fmt"
	"go-markdown-crawler/util"
	"io/ioutil"
	"net/http"
	"strconv"
)

type wordpressProvider struct {
	protocal string
	domain   string
	basepath string
}

func NewWordpressProvider(protocal string, domain string) Provider {
	return &wordpressProvider{protocal, domain, "/wp-json/wp/v2"}
}

func (wp *wordpressProvider) AuthWithUsernameAndPassword(username string, password string) {

}

func (wp *wordpressProvider) GetAllCategories() (Categories, error) {
	totalCount := 0
	totalPage := 0
	pageIndex := 1
	categories := make(Categories, 0)
	for {
		u := util.GenFullRequestUrl(wp.protocal, wp.domain, wp.basepath, "/categories",
			map[string]string{"page": strconv.Itoa(pageIndex), "per_page": "10"})
		resp, err := http.Get(u)
		if err != nil {
			return nil, err
		}
		bys, err := ioutil.ReadAll(resp.Body)
		util.FastClose(resp.Body)
		if err != nil {
			return nil, err
		}

		// 总条数
		totalCount, _ = strconv.Atoi(resp.Header.Get("X-WP-Total"))
		// 总页数
		totalPage, _ = strconv.Atoi(resp.Header.Get("X-WP-TotalPages"))

		fmt.Printf("请求wordpress分类信息完成，总页数%v，总条数%v，当前页数%v\n", totalPage, totalCount, pageIndex)

		result := make([]wordpressCategory, 0)
		err = json.Unmarshal(bys, &result)
		if err != nil {
			return nil, err
		}
		for _, v := range result {
			categories = append(categories, Category{v.ID, v.Name, v.Slug, v.Description, v.Link, v.Parent})
		}

		if pageIndex >= totalPage {
			break
		}
		pageIndex++
	}
	return categories, nil
}

func (wp *wordpressProvider) GetAllTags() (Tags, error) {
	totalCount := 0
	totalPage := 0
	pageIndex := 1
	tags := make(Tags, 0)
	for {
		u := util.GenFullRequestUrl(wp.protocal, wp.domain, wp.basepath, "/tags",
			map[string]string{"page": strconv.Itoa(pageIndex), "per_page": "10"})
		resp, err := http.Get(u)
		if err != nil {
			return nil, err
		}
		bys, err := ioutil.ReadAll(resp.Body)
		util.FastClose(resp.Body)
		if err != nil {
			return nil, err
		}

		// 总条数
		totalCount, _ = strconv.Atoi(resp.Header.Get("X-WP-Total"))
		// 总页数
		totalPage, _ = strconv.Atoi(resp.Header.Get("X-WP-TotalPages"))

		fmt.Printf("请求wordpress标签信息完成，总页数%v，总条数%v，当前页数%v\n", totalPage, totalCount, pageIndex)

		result := make([]wordpressTag, 0)
		err = json.Unmarshal(bys, &result)
		if err != nil {
			return nil, err
		}
		for _, v := range result {
			tags = append(tags, Tag{v.ID, v.Description, v.Name})
		}

		if pageIndex >= totalPage {
			break
		}
		pageIndex++
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
