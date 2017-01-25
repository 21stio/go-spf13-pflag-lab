package main

import (
	"log"

	flag "github.com/spf13/pflag"

	"github.com/21stio/go-spf13-pflag-lab/sub"
)

func init()  {
	var (
		mainFlag = flag.String("main", "", "")
	)

	log.Printf("main: %v", mainFlag)
}

func main() {
	sub.SetFlags()
}
