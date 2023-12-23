package targets

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pibblokto/backlokto/pkg/types"
	"os"
)

func S3Target(target types.Target, backupPath string) {
	var access_key string
	var secret_key string

	if target.AccessKey == nil {
		access_key = os.Getenv("AWS_ACCESS_KEY_ID")
	} else {
		access_key = target.AccessKey
	}

	if target.SecretKey == nil {
		secret_key = os.Getenv("AWS_SECRET_ACCESS_KEY")
	} else {
		secret_key = target.SecretKey
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(target.Region),
		Credentials: credentials.NewStaticCredentials(access_key, secret_key),
	})
}
