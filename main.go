package main

import "raj/tasket/cmd"

func main() {
	cmd.RootCmd.AddCommand(cmd.AddCmd)

	cmd.RootCmd.Execute()
}
