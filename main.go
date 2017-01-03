package main

import (
	"fmt"
	"os"

	"github.com/anyweez/adapt/cmd"
)

func main() {
	if err := cmd.Default.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
