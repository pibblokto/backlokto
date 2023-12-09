package confparse

import (
	"log"
	"os"

	"github.com/pibblokto/backlokto/pkg/types"
)

func ParseConfig(configPath string) types.Jobs {
	var jobs types.Jobs

	f, err := os.ReadFile(configPath)

	if err != nil {
		log.Fatalln(err)
	}

	if err := yaml.Unmarshal(f, &jobs); err != nil {
		log.Fatalln(err)
	}

	return jobs
}
