package providers

import (
	"fmt"
	pg "github.com/habx/pg-commands"
	"github.com/pibblokto/backlokto/pkg/types"
	"os"
	"time"
)

func PostgresPgDump(job *types.BackupJob) {

	var pgPass string

	if job.Spec.Password == "" {
		pgPass = os.Getenv("PG_PASS")
	} else {
		pgPass = job.Spec.Password
	}

	dump, err := pg.NewDump(&pg.Postgres{
		Host:     job.Spec.Host,
		Port:     job.Spec.Port,
		DB:       job.Spec.Database,
		Username: job.Spec.Username,
		Password: pgPass,
	})

	if err != nil {
		fmt.Println(err)
	}

	dump.SetFileName(fmt.Sprintf(`%v_%v.sql`, dump.DB, time.Now().Unix()))
	dump.SetupFormat("p")

	dumpExec := dump.Exec(pg.ExecOptions{StreamPrint: true})
	if dumpExec.Error != nil {
		fmt.Println(dumpExec.Error.Err)
		fmt.Println(dumpExec.Output)
	} else {
		fmt.Println("Dump success")
		fmt.Println(dumpExec.File)
	}
}
