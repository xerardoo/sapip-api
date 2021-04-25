package models

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io"
	"os"
)

func GetAwsSession() (err error, sess *session.Session) {
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_S3_BUCKET_REGION")
	token := ""

	creds := credentials.NewStaticCredentials(accessKey, secretKey, token)
	_, err = creds.Get()
	if err != nil {
		return
	}

	sess, err = session.NewSession(&aws.Config{Region: aws.String(region), Credentials: creds})
	if err != nil {
		return
	}
	return
}

func UploadFileS3(folder string, filename string, file io.Reader) (err error) {
	bucket := os.Getenv("AWS_S3_BUCKET")
	// bucket = bucket + "/" + folder
	if folder == "" {
		fmt.Println("")
	}

	err, sess := GetAwsSession()
	if err != nil {
		return
	}

	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		// ACL:    aws.String("public-read"),
		Key:  aws.String(filename),
		Body: file,
	})
	if err != nil {
		return
	}
	return
}

func GetFilePathS3(filename string) (path string) {
	if filename == "" {
		return ""
	}
	bucket := os.Getenv("AWS_S3_BUCKET")
	region := os.Getenv("AWS_S3_BUCKET_REGION")
	path = fmt.Sprintf("https://%s.s3-%s.amazonaws.com/%s", bucket, region, filename)
	return
}
