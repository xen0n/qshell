package cmd

import (
	"fmt"
	"github.com/qiniu/qshell/v2/iqshell/common/account/operations"
	"github.com/spf13/cobra"
)

var (
	accountOver bool
)

func init() {
	cmdAccount.Flags().BoolVarP(&accountOver, "overwrite", "w", false, "overwrite account or not when account exists in local db, by default not overwrite")
	RootCmd.AddCommand(cmdAccount)
}

var cmdAccount = &cobra.Command{
	Use:   "account [<AccessKey> <SecretKey> <Name>]",
	Short: "Get/Set current account's AccessKey and SecretKey",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 && len(args) != 3 {
			return fmt.Errorf("command account receives zero or three args, received %d\n", len(args))
		}
		return nil
	},
	Run: Account,
}

// 【account】获取本地保存的用户的AK/AK/Name信息； 设置保存用户AK/SK信息到本地
func Account(cmd *cobra.Command, params []string) {
	if len(params) == 0 {
		operations.Current()
	} else if len(params) == 3 {
		accessKey := params[0]
		secretKey := params[1]
		name := params[2]

		operations.Add(operations.AddInfo{
			Name:      name,
			AccessKey: accessKey,
			SecretKey: secretKey,
			Over:      accountOver,
		})
	}
}
