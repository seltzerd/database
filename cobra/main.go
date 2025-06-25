package main

import (
	"fmt"
	"log"

	"fuk/fukkk"

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

	var input = &cobra.Command{
		Use:   "myapp",
		Short: "examp",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Args:", args)
			err := fukkk.Logs(db, fmt.Sprint(args), "good")
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	if err := input.Execute(); err != nil {
		log.Fatal(err)
	}
}
