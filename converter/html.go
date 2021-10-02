package converter

import (
	"bytes"
	"container/list"
	"golang.org/x/net/html"
	"io/ioutil"
	"strings"
)

type HtmlConverter struct {
}

func (hc *HtmlConverter) Convert(data []byte) ([]byte, error) {
	// 解析html
	doc, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	processor := newHtmlMarkdownProcessor()
	processor.process(doc)
	all, _ := ioutil.ReadAll(processor.buffer)

	return all, nil
}

// htmlMarkdownProcessor html -> markdown
type htmlMarkdownProcessor struct {
	context  *list.List
	buffer   *bytes.Buffer
	lastChar string
}

func newHtmlMarkdownProcessor() *htmlMarkdownProcessor {
	return &htmlMarkdownProcessor{
		context: list.New(),
		buffer:  bytes.NewBufferString(""),
	}
}

func (hp *htmlMarkdownProcessor) push(ele *html.Node) {
	if ele == nil || ele.Type != html.ElementNode {
		return
	}
	hp.context.PushBack(ele.Data)
}

func (hp *htmlMarkdownProcessor) pop() (string, bool) {
	last := hp.context.Back()
	if last != nil {
		hp.context.Remove(last)
		return last.Value.(string), true
	}
	return "", false
}

func (hp *htmlMarkdownProcessor) gCtxLinePrefix() string {
	prefix := ""
	for i := hp.context.Front(); i != nil; i = i.Next() {
		switch i.Value {
		case "q", "blockquote":
			prefix += "> "
		case "li":
			prefix += "  "
		}
	}
	return prefix
}

func (hp *htmlMarkdownProcessor) writeString(str string) {
	hp.buffer.WriteString(str)

	// 写入并更新记录的最后一个字符
	if str != "" {
		hp.lastChar = str[len(str)-1:]
	}
}

func (hp *htmlMarkdownProcessor) writeLine() {
	hp.writeString("\n")
}

func (hp *htmlMarkdownProcessor) writeBlockString(str string) {
	current := strings.TrimLeft(str, "\n")

	current = hp.gCtxLinePrefix() + current
	// 判断最后一个字符是否是换行符
	// 如果是换行符，就不用增添新的换行符了
	if hp.lastChar != "" && hp.lastChar != "\n" {
		current = "\n" + current
	}

	hp.writeString(current)
}

func (hp *htmlMarkdownProcessor) processChild(node *html.Node) {
	// 开始处理时加入到上下文中
	hp.push(node)
	// 结束处理时从上下文中移除
	defer hp.pop()
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			data := strings.TrimSpace(c.Data)
			if data != "" {
				hp.writeString(data)
			}
		} else {
			hp.process(c)
		}
	}
}

func (hp *htmlMarkdownProcessor) process(node *html.Node) {
	if node.Type != html.ElementNode && node.Type != html.DocumentNode {
		return
	}

	switch node.Data {
	case "h1":
		hp.writeBlockString("#")
		hp.processChild(node)
		hp.writeLine()
	case "h2":
		hp.writeBlockString("## ")
		hp.processChild(node)
		hp.writeLine()
	case "h3":
		hp.writeBlockString("### ")
		hp.processChild(node)
		hp.writeLine()
	case "h4":
		hp.writeBlockString("#### ")
		hp.processChild(node)
		hp.writeLine()
	case "h5":
		hp.writeBlockString("##### ")
		hp.processChild(node)
		hp.writeLine()
	case "h6":
		hp.writeBlockString("###### ")
		hp.processChild(node)
		hp.writeLine()
	case "span":
		hp.processChild(node)
	case "em", "i":
	case "br":
		hp.writeLine()
	case "strong", "b":
		hp.writeString("**")
		hp.processChild(node)
		hp.writeString("**")
	case "q", "blockquote":
		hp.processChild(node)
	case "p":
		hp.writeBlockString("")
		hp.processChild(node)
	case "del", "s":
		hp.writeString("~~")
		hp.writeString(node.FirstChild.Data)
		hp.writeString("~~")
	case "ins", "u":
		hp.writeString("<u>")
		hp.processChild(node)
		hp.writeString("</u>")
	case "code":
		hp.writeString("`")
		hp.writeString(node.FirstChild.Data)
		hp.writeString("`")
	case "pre":
		code := node.FirstChild
		lang := ""
		val := getAttrVal(code, "class")
		if val != "" {
			vals := strings.Split(val, "language-")
			if len(vals) > 1 {
				lang = vals[1]
			}
		}
		hp.writeBlockString("```" + lang + "\n")
		if node.FirstChild.Type == html.ElementNode && node.FirstChild.Data == "code" {
			hp.writeString(code.FirstChild.Data)
		}
		hp.writeBlockString("```")
	case "hr":
		hp.writeBlockString("---")
	case "a":
		href := getAttrVal(node, "href")
		if node.FirstChild != nil {
			hp.writeString("[" + node.FirstChild.Data + "](" + href + ")")
		}
	case "img":
		src := getAttrVal(node, "src")
		alt := getAttrVal(node, "alt")
		hp.writeString("![" + alt + "](" + src + ")")
	case "ul", "ol":
		hp.writeBlockString("")
		hp.processChild(node)
	case "li":
		if hp.context.Back().Value == "ol" {
			hp.writeBlockString("1. ")
			hp.processChild(node)
		} else if hp.context.Back().Value == "ul" {
			hp.writeBlockString("- ")
			hp.processChild(node)
		}
	default:
		hp.processChild(node)
	}
}

func getAttrVal(node *html.Node, name string) string {
	val := ""
	for _, a := range node.Attr {
		if a.Key == name {
			val = a.Val
		}
	}
	return val
}
