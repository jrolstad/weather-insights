package repositories

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/jrolstad/weather-insights/internal/config"
	"github.com/jrolstad/weather-insights/internal/core"
	"github.com/jrolstad/weather-insights/internal/models"
	"strings"
)

type DaylightRepository interface {
	Save(data *models.DaylightTimes) error
}

func NewDaylightRepository(appConfig *config.AppConfig) DaylightRepository {
	return &S3DaylightRepository{
		bucketName:   appConfig.ObservationBucketName,
		awsRegion:    appConfig.AwsRegion,
		fileUploader: initS3Uploader(appConfig.AwsRegion),
	}
}

type S3DaylightRepository struct {
	bucketName   string
	awsRegion    string
	fileUploader *s3manager.Uploader
}

func (r *S3DaylightRepository) Save(data *models.DaylightTimes) error {

	input := &s3manager.UploadInput{
		Bucket:      aws.String(r.bucketName),
		Key:         aws.String(r.getFileKey(data)),
		Body:        strings.NewReader(core.MapToJson(data)),
		ContentType: aws.String("application/json"),
	}
	_, err := r.fileUploader.UploadWithContext(context.Background(), input)
	if err != nil {
		return err
	}

	return nil
}

func (r *S3DaylightRepository) getFileKey(item *models.DaylightTimes) string {
	path := r.resolveItemPath(item)
	identifier := r.resolveItemId(item)

	return fmt.Sprintf("%v/%v", path, identifier)
}

func (r *S3DaylightRepository) resolveItemPath(item *models.DaylightTimes) string {
	if item == nil || item.LocationIdentifier == "" {
		return "unknown"
	}

	return fmt.Sprintf("daylight/data/%s", strings.ToLower(item.LocationIdentifier))
}

func (r *S3DaylightRepository) resolveItemId(item *models.DaylightTimes) string {
	return core.MapUniqueIdentifier(item.LocationIdentifier, item.Date.String())
}
