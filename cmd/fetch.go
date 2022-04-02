package cmd

import (
	"errors"
	"github.com/yangsx95/markdown-tools/exporter"
	"github.com/yangsx95/markdown-tools/provider"

	"github.com/spf13/cobra"
)

// fetchCmd 抓取数据命令
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "抓取一个url地址的网站，并将内容转换为markdown",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		t := cmd.Flag("type").Value.String()
		o := cmd.Flag("output").Value.String()
		var p provider.Provider
		switch t {
		case "wordpress":
			p = provider.NewWordpressProvider("http", args[0])
		default:
			return errors.New("该类型网站不支持抓取，目前支持的网站有：wordpress")
		}
		var e exporter.Exporter
		if e, err = exporter.NewFileExporter(p, o); err != nil {
			return
		}

		return e.Export()
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	fetchCmd.Flags().StringP("type", "t", "", "网站类型：wordpress")
	fetchCmd.Flags().StringP("output", "o", ".", "输出路径")
	_ = fetchCmd.MarkFlagRequired("type")

}
