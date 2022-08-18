package cmd

import (
	"fmt"
	"log"
	"strings"
	"task/task/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		id, err := db.CreateTask(task)
		if err != nil {
			log.Printf(err.Error())
			return
		}
		fmt.Printf("> added \"%s\" to your task list with id %d\n", task, id)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
