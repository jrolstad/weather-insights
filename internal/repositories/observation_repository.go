package repositories

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/jrolstad/weather-insights/internal/config"
	"github.com/jrolstad/weather-insights/internal/core"
	"github.com/jrolstad/weather-insights/internal/logging"
	"github.com/jrolstad/weather-insights/internal/models"
	"strings"
)

type ObservationRepository interface {
	Save(data []*models.Observation) error
}

func NewObservationRepository(appConfig *config.AppConfig) ObservationRepository {
	return &S3ObservationRepository{
		bucketName:   appConfig.ObservationBucketName,
		awsRegion:    appConfig.AwsRegion,
		fileUploader: initS3Uploader(appConfig.AwsRegion),
	}
}

type S3ObservationRepository struct {
	bucketName   string
	awsRegion    string
	fileUploader *s3manager.Uploader
}

func (r *S3ObservationRepository) Save(data []*models.Observation) error {
	if data == nil || len(data) == 0 {
		return nil
	}

	saveErrors := make([]error, 0)

	for _, item := range data {
		err := r.saveObservation(item)
		if err != nil {
			saveErrors = append(saveErrors, err)
		}
	}

	return core.ConsolidateErrors(saveErrors)
}

func (r *S3ObservationRepository) saveObservation(item *models.Observation) error {

	input := &s3manager.UploadInput{
		Bucket:      aws.String(r.bucketName),
		Key:         aws.String(r.getFileKey(item)),
		Body:        strings.NewReader(core.MapToJson(item)),
		ContentType: aws.String("application/json"),
	}
	_, err := r.fileUploader.UploadWithContext(context.Background(), input)
	if err != nil {
		return err
	}

	return nil
}

func (r *S3ObservationRepository) getFileKey(item *models.Observation) string {
	path := r.resolveItemPath(item)
	identifier := r.resolveItemId(item)

	return fmt.Sprintf("%v/%v", path, identifier)
}

func (r *S3ObservationRepository) resolveItemPath(item *models.Observation) string {
	if item == nil || item.Station.MacAddress == "" {
		return "unknown"
	}

	return strings.ToLower(item.Station.MacAddress)
}

func (r *S3ObservationRepository) resolveItemId(item *models.Observation) string {
	return core.MapUniqueIdentifier(item.Station.MacAddress, item.At.String())
}

func initS3Uploader(awsRegion string) *s3manager.Uploader {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion)},
	)
	if err != nil {
		logging.LogError(err, "context", "failed to create AWS session")
	}

	uploader := s3manager.NewUploader(sess)
	return uploader
}
