package main

import (
	"fmt"
	"github.com/pibblokto/backlokto/pkg/confparse"
	"github.com/pibblokto/backlokto/pkg/providers"
	"github.com/pibblokto/backlokto/pkg/types"
)

var ProvidersMap = map[string]func(*types.BackupJob) types.Artifacts{
	"postgres.pg_dump": providers.PostgresPgDump,
}

func main() {

	var jobs types.Jobs = confparse.ParseConfig("examples/configs/test.yaml")

	// Running providers
	for _, job := range jobs.Jobs {
		fmt.Printf("Running %s...", job.Id)
		ProvidersMap[job.Provider](&job)
	}

}
