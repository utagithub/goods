/*
 * @Author: 尤_Ta
 * @Date:  23:13
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  23:13
 */

package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"goods/cmd/migrate"
	"goods/cmd/server"
	"goods/cmd/version"
	"goods/common/global"
	"goods/common/tools"
	"os"
)

var rootCmd = &cobra.Command{
	Use:          "goods",
	Short:        "goods",
	SilenceUsage: true,
	Long:         `goods`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New(tools.Red("requires at least one arg"))
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `欢迎使用 ` + tools.Green(`goods `+global.Version) + ` 可以使用 ` + tools.Red(`-h`) + ` 查看命令`
	usageStr1 := `也可以参考 https://doc.goods.dev/guide/ksks 的相关内容`
	fmt.Printf("%s\n", usageStr)
	fmt.Printf("%s\n", usageStr1)
}

func init() {
	rootCmd.AddCommand(server.StartCmd)
	rootCmd.AddCommand(migrate.StartCmd)
	rootCmd.AddCommand(version.StartCmd)
	//rootCmd.AddCommand(config.StartCmd)
	//rootCmd.AddCommand(app.StartCmd)
}

// Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
