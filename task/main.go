package main

import (
	"log"
	"os"
	"path/filepath"
	"task/task/cmd"
	"task/task/db"

	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	must(db.Init(filepath.Join(home, "tasks.db")))
	must(cmd.RootCmd.Execute())

}

func must(err error) {
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
}
