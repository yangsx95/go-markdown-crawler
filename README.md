# markdown-tools

使用Golang编写的Markdown工具，可以很方便的操作Markdown文件。

## 安装

```shell
go install github.com/yangsx95/markdown-tools@latest
```

## 使用

清理markdown文件夹中的没有被引用的资源：

```shell
markdown-tools clean markdown_file_path markdown_resource_path
```

演习清理：

```shell
markdown-tools clean markdown_file_path markdown_resource_path -e
```
