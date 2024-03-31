package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io/fs"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var ignoredDirs = []string{
	".git",
	".trash",
	".idea",
	".vscode",
	".obsidian",
}

var ignoredFiles = []string{
	".DS_Store",
	".gitignore",
}

// cleanCmd 清理命令
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "清理本地的markdown文件中的冗余信息",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("按照规范指定markdown扫描路径与清理路径")
		}
		if len(args) > 2 {
			return errors.New("按照规范指定markdown扫描路径与清理路径")
		}
		stat, err := os.Stat(args[0])
		if os.IsExist(err) {
			return errors.New("扫描路径不存在")
		}
		// 如果目标是个文件 && 文件不是md文件，不处理
		if !stat.IsDir() && strings.LastIndex(stat.Name(), ".md") < 0 {
			return errors.New("目标不是以.md结尾的文件，不支持清理")
		}
		stat, err = os.Stat(args[1])
		if os.IsExist(err) {
			return errors.New("清理路径不存在")
		}
		if !stat.IsDir() {
			return errors.New("清理路径不能是一个文件")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		//imgFlag := cmd.Flag("image").Value.String()
		exerciseFlag := cmd.Flag("exercise").Value.String()
		stat, _ := os.Stat(args[0])
		// 从扫描路径中找出所有的md文件，以及非md文件
		allMdPath := make([]string, 0)
		allResPath := make([]string, 0)
		if !stat.IsDir() {
			allMdPath = append(allMdPath, args[0])
		} else {
			mdPath, resPath, err := getAllMdAndResourceFromDir(args[0])
			if err != nil {
				return err
			}
			allMdPath = append(allMdPath, mdPath...)
			allResPath = append(allResPath, resPath...)
		}
		// 从清理路径中找出所有的md文件以及非md文件
		mdPath, resPath, err := getAllMdAndResourceFromDir(args[1])
		if err != nil {
			return err
		}
		allMdPath = append(allMdPath, mdPath...)
		allResPath = append(allResPath, resPath...)

		// 找出所有的资源文件，不包含md
		allRefRes := make([]string, 0)
		// 找出所有的md文件中引用的资源
		for _, mdPath := range allMdPath {
			refRes, err := getMdFileResource(mdPath)
			if err != nil {
				return err
			}
			allRefRes = append(allRefRes, refRes...)
		}

		// 找出所有未被引用的资源文件
		allUnRefRes := make([]string, 0)
		for _, aRes := range allResPath {
			if !ResourcePathContains(allRefRes, aRes) {
				allUnRefRes = append(allUnRefRes, aRes)
			}
		}

		fmt.Printf("一共查询出%d个Markdown文件, %d个资源(非md)文件, 其中md文件中共引用的资源有%d个，未被引用的资源有%d个\n", len(allMdPath), len(allResPath), len(allRefRes), len(allUnRefRes))
		fmt.Println("这些资源为：")
		for _, p := range allUnRefRes {
			fmt.Println("\t" + p)
		}

		if exerciseFlag != "true" {
			fmt.Println("清理中..")
			for _, unused := range allUnRefRes {
				_ = os.Remove(unused)
			}
			fmt.Println("清理成功...")
		}

		return err
	},
}

func ResourcePathContains(target []string, ele string) bool {
	for _, t := range target {
		trimT := strings.TrimSpace(t)
		trimEle := strings.TrimSpace(ele)
		if trimT == trimEle || url.PathEscape(trimT) == trimEle {
			return true
		}
	}
	return false
}

func init() {
	rootCmd.AddCommand(cleanCmd)
	cleanCmd.SetUsageTemplate("md-crawler clean <scan_path> <clean_path> [flags]")
	cleanCmd.Flags().BoolP("exercise", "e", false, "演习清理")
}

func getMdFileResource(mdPath string) ([]string, error) {
	file, err := os.ReadFile(mdPath)
	if err != nil {
		fmt.Println("读取文件失败" + mdPath + "0--")
		return nil, err
	}
	r := regexp.MustCompile(`\[.*\]\((.+)\)`)
	resList := r.FindAllStringSubmatch(string(file), -1)
	result := make([]string, 1)
	for _, res := range resList {
		resPath := res[1]
		// 如果文件链接为url，跳过处理
		if strings.HasPrefix(resPath, "http") {
			fmt.Printf("发现url资源 %s %s \n", mdPath, resPath)
			continue
		}
		resPath, _ = url.QueryUnescape(resPath)
		// 如果文件链接不是绝对路径，将其转换为绝对路径
		if !filepath.IsAbs(res[1]) {
			resPath = filepath.Join(filepath.Dir(mdPath), resPath)
		}
		result = append(result, resPath)
	}
	return result, nil
}

var processedFiles = make(map[string]bool)

func getAllMdAndResourceFromDir(dir string) (mdPath []string, resPath []string, err error) {
	err = filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {

		if err != nil {
			fmt.Printf("Error accessing path: %v\n", err)
			return err
		}

		// 跳过忽略的的文件夹
		if d.IsDir() && isIgnoreDir(d.Name()) {
			return filepath.SkipDir
		}

		// 文件夹不做任何处理
		if d.IsDir() {
			return err
		}

		// 忽略的文件不处理
		if isIgnoredFile(d.Name()) {
			return err
		}

		// 检查文件是否已经处理过
		if processedFiles[path] {
			return err
		} else {
			processedFiles[path] = true
		}

		path, err = filepath.Abs(path)
		if err != nil {
			return err
		}
		info, _ := os.Stat(path)
		if strings.LastIndex(info.Name(), ".md") > 0 {
			mdPath = append(mdPath, path)
		} else {
			resPath = append(resPath, path)
		}
		return err
	})
	if err != nil {
		return nil, nil, err
	}
	return mdPath, resPath, err
}

func isIgnoreDir(name string) bool {
	for _, dir := range ignoredDirs {
		if name == dir {
			return true
		}
	}
	return false
}

func isIgnoredFile(name string) bool {
	for _, file := range ignoredFiles {
		if strings.ToLower(name) == strings.ToLower(file) {
			return true
		}
	}
	return false
}
