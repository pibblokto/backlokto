package main

import (
	"flag"
	"fmt"

	"github.com/pibblokto/backlokto/pkg/confparse"
	"github.com/pibblokto/backlokto/pkg/providers"
	"github.com/pibblokto/backlokto/pkg/types"
)

var ProvidersMap = map[string]func(*types.BackupJob){
	"postgres.pg_dump": providers.PostgresPgDump,
}

func main() {

	configpathPtr := flag.String("config", "", "Path to config file with backup jobs declaration")
	flag.Parse()

	var jobs types.Jobs = confparse.ParseConfig(*configpathPtr)

	// Running providers
	for _, job := range jobs.Jobs {
		fmt.Printf("Running %s...", job.Id)
		ProvidersMap[job.Provider](&job)
	}

}
