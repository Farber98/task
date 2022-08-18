package cmd

import (
	"fmt"
	"log"
	"strconv"
	"task/task/db"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "marks task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			arg, err := strconv.Atoi(arg)
			if err != nil {
				log.Println("failed to parse argument ", arg)
			} else {
				ids = append(ids, arg)
			}
		}
		tasks, err := db.ReadTasks()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("> invalid task number: ", id)
				continue
			}
			err := db.DeleteTask(tasks[id-1].Key)
			if err != nil {
				fmt.Printf("> failed to complete task %d. Error: %s\n", id, err.Error())
			} else {
				fmt.Printf("> marked %d as completed \n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
