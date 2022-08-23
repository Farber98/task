package cmd

import (
	"fmt"
	"log"
	"strings"
	"task/task/db"
	"time"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		err := db.CreateTask(&db.Task{Value: task, Created: time.Now()})
		if err != nil {
			log.Printf(err.Error())
			return
		}
		fmt.Printf("> added \"%s\" to your task list\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
