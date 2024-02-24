package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/pibblokto/backlokto/pkg/confparse"
	"github.com/pibblokto/backlokto/pkg/providers"
	"github.com/pibblokto/backlokto/pkg/types"
)

var ProvidersMap = map[string]func(*types.BackupJob){
	"postgres.pg_dump": providers.PostgresPgDump,
}

func main() {

	configpathPtr := flag.String("config", "", "Path to config file with backup jobs declaration")
	tagsPtr := flag.String("tags", "", "List of tags to group jobs in the following format: KEY1=VALUE1,KEY2=VALUE2")
	flag.Parse()

	tagsMap := make(map[string]string)
	if *tagsPtr != "" {
		pairs := strings.Split(*tagsPtr, ",")

		// Iterate over key-value pairs
		for _, pair := range pairs {
			// Split pair by equal sign
			parts := strings.Split(pair, "=")
			if len(parts) == 2 {
				key := parts[0]
				value := parts[1]
				tagsMap[key] = value
			}
		}
	}

	var jobs types.Jobs = confparse.ParseConfig(*configpathPtr)

	for _, job := range jobs.Jobs {
		tags := job.Tags
		for k, v := range tags {
			fmt.Printf("key: %s, value: %s\n", k, v)
		}
	}

	// Running providers
	//for _, job := range jobs.Jobs {
	//	fmt.Printf("Running %s...", job.Id)
	//	ProvidersMap[job.Provider](&job)
	//}

}
