package providers

import (
	"github.com/pibblokto/backlokto/pkg/targets"
)

var TargetsMap = map[string]func(){
	"s3": targets.S3Target,
	// "local": target.LocalTarget
}
