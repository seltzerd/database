package main

import (
	"fmt"
	"log"
	"strings"

	"fuk/fukkk"
	"fuk/serv"

	"github.com/spf13/cobra"
)

func main() {
	db, err := fukkk.Dbinit()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = fukkk.CreateTable(db)
	if err != nil {
		log.Fatal(err)
	}

	var argsStr string

	var input = &cobra.Command{
		Use:   "myapp",
		Short: "examp",
		Run: func(cmd *cobra.Command, args []string) {
			argsStr = strings.Join(args, " ")
			fmt.Println(argsStr)

			err := fukkk.Logs(db, argsStr, "good")
			if err != nil {
				log.Fatal(err)
			}
			serv.Server(&argsStr, db)
		},
	}

	if err := input.Execute(); err != nil {
		log.Fatal(err)
	}
}
