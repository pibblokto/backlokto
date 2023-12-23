package main

import (
	"fmt"
	"github.com/pibblokto/backlokto/pkg/confparse"
	"github.com/pibblokto/backlokto/pkg/types"
)

func main() {

	var jobs types.Jobs = confparse.ParseConfig("examples/configs/config.yaml")

	fmt.Println(jobs)

}
