package targets

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pibblokto/backlokto/pkg/types"
	"os"
	"bytes"
	"io"
)

func S3Target(target types.Target, backupPath string) {
	var access_key string
	var secret_key string
	bucketName := target.S3BucketName
	bucketKey := target.S3BucketKey

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
		Credentials: credentials.NewStaticCredentials(access_key, secret_key, ""),
	})

	svc := s3.New(sess)

	file, err := os.Open(backupPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file:", err)
		return
	}
	defer file.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, file); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading file:", err)
	}

	_, err = svc.PutObjet(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key: aws.String(Key),
		Body: bytes.NewReader(buf.Bytes())
	})

	if err != nil {
		fmt.Println("Error uploading file:", err)
		return
	}
}
