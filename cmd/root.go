package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd 代表没有子命令的情况下，调用的基本命令
var rootCmd = &cobra.Command{
	Use:   "md-crawler",
	Short: "",
	Long:  ``,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
