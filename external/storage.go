package external

import (
	"image"
	"image/png"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/pkg/errors"
)

// GetS3Uploader set s3 config.
func GetS3Uploader() *s3manager.Uploader {
	conf := &aws.Config{
		Region: aws.String(endpoints.ApNortheast1RegionID),
	}
	sess := session.New(conf)
	svc := s3manager.NewUploader(sess)
	return svc
}

// UploadImg upload image to s3.
func (r *Repository) UploadImg(img image.Image, Name string) (string, error) {
	fileName := strings.Split(Name, "/")[1]
	filePath := fileName + ".png"
	file, err := os.Create(filePath)
	if err != nil {
		return "", errors.Wrap(err, "failed to create file")
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return "", errors.Wrap(err, "failed to encode to file")
	}
	file, err = os.Open(filePath)
	if err != nil {
		return "", errors.Wrap(err, "failed to open file")
	}

	svc := GetS3Uploader()
	result, err := svc.Upload(&s3manager.UploadInput{
		Bucket: aws.String("github-link-card"),
		Key:    aws.String(Name + ".png"),
		Body:   file,
	})
	if err != nil {
		return "", errors.Wrap(err, "failed to upload file")
	}
	err = os.Remove(fileName + ".png")
	if err != nil {
		return "", errors.Wrap(err, "failed to delete file")
	}
	return result.Location, nil
}
