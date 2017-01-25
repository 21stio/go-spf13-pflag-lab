package sub

import (
	"log"

	flag "github.com/spf13/pflag"
)

func init()  {
	var (
		flagset = flag.NewFlagSet("subInit", flag.ContinueOnError)
		subInitFlag = flagset.String("subInit", "", "")
	)


	log.Printf("subInit: %v", subInitFlag)
}

func SetFlags() {
	var (
		flagset = flag.NewFlagSet("sub", flag.ContinueOnError)
		subFlag = flagset.String("sub", "", "")
	)



	log.Printf("sub: %v", subFlag)
}
