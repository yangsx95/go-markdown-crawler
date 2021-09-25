package provider

import "time"

// Provider 数据提供接口
type Provider interface {

	// AuthWithUsernameAndPassword 使用用户名密码认证
	AuthWithUsernameAndPassword(username string, password string)

	// GetAllCategories 获取所有的分类信息
	GetAllCategories() (Categories, error)

	// GetAllTags 获取所有文章标签信息
	GetAllTags() (Tags, error)

	// GetArticleIterator 获取文章迭代器
	GetArticleIterator() (ArticleIterator, error)
}

// Category 分类
type Category struct {
	// 分类Id
	Id int
	// 分类展示名称
	Name string
	// 分类别名
	Alias string
	// 分类描述
	Description string
	// 分类下的URL
	Link string
	// 父分类Id
	Parent int
}

// Categories 所有分类
type Categories []Category

// GetCategoryById 根据Id获取Category信息
func (cs *Categories) GetCategoryById(id int) (Category, bool) {
	for _, c := range *cs {
		if c.Id == id {
			return c, true
		}
	}
	return Category{}, false
}

// Tag 标签信息
type Tag struct {
	Id          int
	Description string
	Name        string
}

// Tags 所有标签
type Tags []Tag

// GetTagById 根据Id查询标签
func (ts *Tags) GetTagById(id int) (Tag, bool) {
	for _, t := range *ts {
		if t.Id == id {
			return t, true
		}
	}
	return Tag{}, false
}

// Article 文章
type Article struct {
	Id          int
	Date        time.Time
	Link        string
	Title       string
	Content     string
	ContentType string
	CategoryIds []int
	TagIds      []int
}

// Articles 多篇文章
type Articles []Article

// ArticleIterator 文章迭代器
type ArticleIterator interface {
	HasNext() bool
	Next() (Article, error)
}
