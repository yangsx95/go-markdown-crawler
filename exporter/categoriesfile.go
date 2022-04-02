package exporter

import (
	"fmt"
	"github.com/yangsx95/markdown-tools/converter"
	"github.com/yangsx95/markdown-tools/provider"
	"github.com/yangsx95/markdown-tools/util"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type CategoryFileExporter struct {
	// 导出的目标路径
	path string
	// categoryIdFilePathMap 分类与本地文件夹的映射关系
	categoryIdFilePathMap map[int]string
	// 数据源
	provider provider.Provider
}

func NewFileExporter(provider provider.Provider, path string) (*CategoryFileExporter, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return &CategoryFileExporter{}, err
	}

	err = util.MkDirForce(absPath)
	if err != nil {
		return &CategoryFileExporter{}, err
	}

	m := make(map[int]string)
	return &CategoryFileExporter{provider: provider, path: absPath, categoryIdFilePathMap: m}, err
}

func (fe *CategoryFileExporter) Export() (err error) {
	defer func() {

	}()
	var categories provider.Categories
	if categories, err = fe.provider.GetAllCategories(); err != nil {
		return
	}
	fmt.Printf("读取wordpress分类信息成功\n")

	for _, category := range categories {
		fe.mkdir(categories, category)
	}
	fmt.Printf("根据wordpress分类创建目录成功 \n")

	var aIterator provider.ArticleIterator
	if aIterator, err = fe.provider.GetArticleIterator(); err != nil {
		return
	}
	fmt.Printf("准备进行文章读取~\n")

	for aIterator.HasNext() {
		next, err := aIterator.Next()
		fmt.Printf("开始处理wordpress文章，目标文章 %v, 目标地址 %v \n", next.Title, next.Link)
		if err != nil {
			fmt.Printf("处理wordpress文章失败，目标文章 %v, 目标地址 %v, 错误信息 %v\n", next.Title, next.Link, err)
		}
		path, ok := fe.categoryIdFilePathMap[next.CategoryIds[0]]
		if ok { // 找到路径的文章
			// 将目标转换为markdown文本
			htmlConv := converter.NewConverter(converter.GetContentTypeByString(next.ContentType), &converter.Option{
				SaveImgToLocal:  true,
				MarkdownPath:    path,
				ImgRelativePath: strings.TrimSpace(next.Title) + ".assets",
			})
			markdown, err := htmlConv.Convert([]byte(next.Content))
			if err != nil {
				return err
			}
			// 写入markdown文本
			err = ioutil.WriteFile(filepath.Join(path, next.Title+".md"), markdown, 0777)
		}
		if err != nil {
			fmt.Printf("处理wordpress文章失败，目标文章 %v, 目标地址 %v, 错误信息 %v\n", next.Title, next.Link, err)
		}
		fmt.Printf("处理wordpress文章成功，目标文章 %v, 目标地址 %v \n", next.Title, next.Link)

	}
	return err
}

// 递归，类似java classloader双亲委派
// 判断文件夹是否有父文件夹，有的话先创建父文件夹，最后再递归回来创建子文件件
// 如果文件夹已经存在，说明他的父文件夹一定已经被创建
// 记录文件夹与category的对应关系
func (fe *CategoryFileExporter) mkdir(categories provider.Categories, category provider.Category) {
	_, ok := fe.categoryIdFilePathMap[category.Id]

	// 文件夹已经创建
	if ok {
		return
	}

	// 文件夹没有创建

	// 是根文件夹，直接创建
	if category.Parent == 0 {
		fe.categoryIdFilePathMap[category.Id] = filepath.Join(fe.path, category.Name)
		_ = util.MkDirForce(fe.categoryIdFilePathMap[category.Id])
		return
	}
	// 如果不是根文件夹，创建父文件夹
	parentCategory, hasParent := categories.GetCategoryById(category.Parent)
	if hasParent {
		fe.mkdir(categories, parentCategory)
		// 创建成功后，创建当前文件夹
		fe.categoryIdFilePathMap[category.Id] = filepath.Join(fe.categoryIdFilePathMap[category.Parent], category.Name)
		_ = util.MkDirForce(fe.categoryIdFilePathMap[category.Id])
	}
}
