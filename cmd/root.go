package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "tasket",
	Short:   "Tasket",
	Long:    `Tasket is a Golang based Todo App to keep track of your tasks`,
	Aliases: []string{"tasket"},
	Version: "0.1.0",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(AddCmd)
	RootCmd.AddCommand(ListCmd)
	RootCmd.AddCommand(DeleteCmd)
	RootCmd.AddCommand(DoneCmd)
	RootCmd.AddCommand(DoingCmd)
}
