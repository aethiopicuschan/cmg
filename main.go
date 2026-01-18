package main

import (
	"github.com/aethiopicuschan/cmg/cmd"
	"github.com/aethiopicuschan/cmg/pkg/logs"
)

func main() {
	if err := cmd.Execute(); err != nil {
		logs.Fatal(err.Error())
	}
}
