package providers

import (
	"github.com/habx/pg-commands"
	"github.com/pibblokto/backlokto/pkg/types"
)

func PostgresPgDump(job types.BackupJob) {
	dump, _ := pg.NewDump(&pg.Postgres{
		Host:     job.Spec.,
		Port:     5432,
		DB:       "dev_example",
		Username: "example",
		Password: "example",
	})

	dumpExec := dump.Exec(pg.ExecOptions{StreamPrint: false})
	if dumpExec.Error != nil {
		fmt.Println(dumpExec.Error.Err)
		fmt.Println(dumpExec.Output)
	} else {
		fmt.Println("Dump success")
		fmt.Println(dumpExec.Output)
	}
}
