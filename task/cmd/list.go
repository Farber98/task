package cmd

import (
	"fmt"
	"log"
	"os"
	"task/task/db"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list of all your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ReadTasks()
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("> no tasks to complete")
		} else {
			fmt.Println("> tasks:")
			for i, task := range tasks {
				fmt.Printf("%d. %s\n", i+1, task.Value)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

}
