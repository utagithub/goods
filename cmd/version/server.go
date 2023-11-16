/*
 * @Author: 尤_Ta
 * @Date:  23:05
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  23:05
 */

package version

import (
	"fmt"
	"github.com/spf13/cobra"
	"goods/common/global"
)

var (
	StartCmd = &cobra.Command{
		Use:     "version",
		Short:   "Get version info",
		Example: "goods version",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func run() error {
	fmt.Println(global.Version)
	return nil
}
