package targets

//import (
//	"github.com/aws/aws-sdk-go/aws"
//	"github.com/aws/aws-sdk-go/aws/credentials"
//	"github.com/aws/aws-sdk-go/aws/session"
//	"github.com/aws/aws-sdk-go/service/s3"
//	"github.com/pibblokto/backlokto/pkg/types"
//	"os"
//)

// to be deleted
func S3Target() {}

//func S3Target(target types.Target, backupPath string) {
//	var access_key string
//
//	var secret_key string
//
//	if target.AccessKey == "" {
//		access_key = os.Getenv("AWS_ACCESS_KEY_ID")
//	} else {
//		access_key = target.AccessKey
//	}
//
//	if target.SecretKey == "" {
//		secret_key = os.Getenv("AWS_SECRET_ACCESS_KEY")
//	} else {
//		secret_key = target.SecretKey
//	}
//
//	sess, err := session.NewSession(&aws.Config{
//		Region:      aws.String(target.Region),
//		Credentials: credentials.NewStaticCredentials(access_key, secret_key),
//	})
//}
