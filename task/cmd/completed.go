package cmd

import (
	"fmt"
	"log"
	"os"
	"task/task/db"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "list of all your completed tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ReadCompleted()
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("> no tasks completed")
		} else {
			fmt.Println("> tasks:")
			for i, task := range tasks {
				fmt.Printf("%d. %s  [%s] [%s]\n", i+1, task.Value, task.Created.Format("2006-01-02 15:04:05"), task.Completed.Format("2006-01-02 15:04:05"))
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(completedCmd)

}
