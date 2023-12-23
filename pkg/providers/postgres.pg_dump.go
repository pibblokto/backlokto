package providers

import (
	"github.com/habx/pg-commands"
	"github.com/pibblokto/backlokto/pkg/types"
	"os"
)

func PostgresPgDump(job types.BackupJob) {

	var pgPass string

	if job.Spec.Password == nil {
		pgPass = os.Getenv("PG_PASS")
	} else {
		pgPass = job.Spec.Password
	}

	dump, _ := pg.NewDump(&pg.Postgres{
		Host:     job.Spec.Host,
		Port:     job.Spec.Port,
		DB:       job.Spec.DB,
		Username: job.Spec.Username,
		Password: pgPass,
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
