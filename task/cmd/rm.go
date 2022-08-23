package cmd

import (
	"fmt"
	"log"
	"strconv"
	"task/task/db"

	"github.com/spf13/cobra"
)

// removeCmd represents the do command
var removeCmd = &cobra.Command{
	Use:   "rm",
	Short: "removes tasks from the tasklist",
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
				fmt.Printf("> failed to remove task '%s' with id %d. Error: %s\n", tasks[id-1].Value, id, err.Error())
			} else {
				fmt.Printf("> removed task '%s' with id %d \n", tasks[id-1].Value, id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(removeCmd)

}
