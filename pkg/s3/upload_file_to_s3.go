package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	log "github.com/sirupsen/logrus"
	"os"
)

func GetSession(region, accessKeyID, secretAccessKey string) *session.Session {
	sess, err := session.NewSession(
		&aws.Config{
			Region:      aws.String(region),
			Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
		},
	)
	if err != nil {
		log.WithFields(log.Fields{"region": region, "accessKeyID": accessKeyID}).Fatal("Can't create aws session with given credentials")
	}
	return sess
}

func UploadFileToS3(session *session.Session, bucket, filepath, filename string) {
	log.WithFields(log.Fields{"bucket": bucket, "source": filepath, "dest": filename}).Info("Uploading file to S3")
	file, err := os.Open(filepath)
	if err != nil {
		log.WithFields(log.Fields{"filepath": filepath}).Fatal("Unable to open file")
	}

	defer file.Close()

	uploader := s3manager.NewUploader(session)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		log.WithFields(log.Fields{"bucket": bucket, "source": filepath, "dest": filename, "error": err}).Fatal("Upload to S3 failed")
	}

	log.WithFields(log.Fields{"bucket": bucket, "source": filepath, "dest": filename}).Info("Upload to S3 succeeded")
}
