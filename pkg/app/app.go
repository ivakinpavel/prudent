package app

import (
	"fmt"
	"github.com/ivakinpavel/prudent/pkg/config"
	"github.com/ivakinpavel/prudent/pkg/s3"
	"github.com/ivakinpavel/prudent/pkg/shell"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

func Start(config *config.Config) {
	backupFileName := getBackupFileName(config.PostgresDBName, "db")
	backupFilePath := "/tmp/" + backupFileName
	dumpDB(config.PostgresHost, config.PostgresDBName, config.PostgresUsername, config.PostgresPassword, backupFilePath)

	awsSession := s3.GetSession(config.AWSRegion, config.AWSAccessKeyID, config.AWSSecretKey)
	s3.UploadFileToS3(awsSession, config.AWSBucket, backupFilePath, backupFileName)

	defer removeBackupFile(backupFilePath)
}

func dumpDB(host, dbname, username, password, outputFilePath string) {
	log.WithFields(log.Fields{"host": host, "dbname": dbname}).Info("Running pg_dump")
	cmd := fmt.Sprintf("PGPASSWORD=%s pg_dump -Fc --host=%s --dbname=%s --username=%s --file=%s",
		password, host, dbname, username, outputFilePath)

	err, stdout, stderr := shell.ExecuteShellCmd(cmd)
	if err != nil {
		log.WithFields(log.Fields{"error": err, "stdout": stdout, "stderr": stderr}).Fatal("pg_dump cmd failed")
	}
}

func getBackupFileName(prefixes ...string) string {
	var toJoin []string
	for _, pref := range prefixes {
		toJoin = append(toJoin, pref)
	}
	toJoin = append(toJoin, time.Now().Format(time.RFC3339)+".gz")
	return strings.Join(toJoin, "-")
}

func removeBackupFile(backupFileName string) {
	err := os.Remove(backupFileName)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Can't remove backup file")
	}
}
