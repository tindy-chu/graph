package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "Moovup map demo",
	Long:  "Moovup programming-test\nref: https://github.com/moovup/programming-test/blob/master/web.md",
}

type tResp map[string]interface{}

func response(resp *tResp) {
	jsonString, _ := json.Marshal(resp)

	fmt.Println(string(jsonString))
}

func init() {
	rootCmd.AddCommand(possibleCmd)
	rootCmd.AddCommand(shortestCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
