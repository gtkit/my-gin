package cmd

import (
	"fmt"

	"github.com/gtkit/encry/rsa"
	"github.com/gtkit/logger"
	"github.com/spf13/cobra"
)

var action string

// newrsaCmd 初始化生成 rsa 公钥 和 私钥
var newrsaCmd = &cobra.Command{
	Use: "newrsa",
	Run: func(cmd *cobra.Command, args []string) {
		switch action {
		case "n":
			newrsa()
		case "e":
			enrsa()
		case "d":
			dersa()
		}
	},
}

func init() {
	rootCmd.AddCommand(newrsaCmd)
	newrsaCmd.Flags().StringVarP(&action, "action", "a", "e", "seclect action")
}

func newrsa() {
	fmt.Println("-----new rsa ------")
	err := rsa.GenerateRsaKey(1024, "/Users/xiaozhaofu/go/src/officekey/cmd/rsapem/")
	logger.LogIf(err)
}

func enrsa() {
	fmt.Println("------ enrsa ----")
}

func dersa() {
	fmt.Println("-----dersa ---")
}
