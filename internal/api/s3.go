package api

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

//nolint:lll // go generate is ugly.
//go:generate mockgen -destination=mocks/s3_mock.go -package=mocks github.com/omegion/s3-secret-manager/internal/api APIInterface
// APIInterface is an interface for API.
type APIInterface interface {
	GetObject(options *GetObjectOptions) (*s3.GetObjectOutput, error)
	ListObjectVersions(options *ListObjectVersionsOptions) (*s3.ListObjectVersionsOutput, error)
	ListObjects(options *ListObjectOptions) (*s3.ListObjectsV2Output, error)
	PutObject(options *PutObjectOptions) (*s3.PutObjectOutput, error)
	DeleteObject(options *DeleteObjectOptions) (*s3.DeleteObjectOutput, error)
}

// API is main struct of S3.
type API struct {
	Config *aws.Config
	Client *s3.Client
}

// PutObjectOptions is options for API call.
type PutObjectOptions struct {
	Bucket,
	Path string
	Value       io.Reader
	EncodedTags string
}

// GetObjectOptions is options for API call.
type GetObjectOptions struct {
	Bucket,
	Path string
	VersionID *string
}

// ListObjectVersionsOptions is options for API call.
type ListObjectVersionsOptions struct {
	Bucket,
	Path string
}

// ListObjectOptions is options for API call.
type ListObjectOptions struct {
	Bucket,
	Path string
}

// DeleteObjectOptions is options for API call.
type DeleteObjectOptions struct {
	Bucket,
	Path string
}

// NewAPI inits new API.
func NewAPI() (APIInterface, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	return &API{
		Config: &cfg,
		Client: client,
	}, nil
}

// GetObject gets object from S3.
func (a API) GetObject(options *GetObjectOptions) (*s3.GetObjectOutput, error) {
	return a.Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket:    aws.String(options.Bucket),
		Key:       aws.String(options.Path),
		VersionId: options.VersionID,
	})
}

// ListObjectVersions gets object versions from S3.
func (a API) ListObjectVersions(options *ListObjectVersionsOptions) (*s3.ListObjectVersionsOutput, error) {
	return a.Client.ListObjectVersions(context.TODO(), &s3.ListObjectVersionsInput{
		Bucket: aws.String(options.Bucket),
		Prefix: aws.String(options.Path),
	})
}

// ListObjects gets object from S3.
func (a API) ListObjects(options *ListObjectOptions) (*s3.ListObjectsV2Output, error) {
	resp, err := a.Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(options.Bucket),
		Prefix: aws.String(options.Path),
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// PutObject puts object to S3.
func (a API) PutObject(options *PutObjectOptions) (*s3.PutObjectOutput, error) {
	return a.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:  aws.String(options.Bucket),
		Key:     aws.String(options.Path),
		Body:    options.Value,
		Tagging: aws.String(options.EncodedTags),
	})
}

// DeleteObject deletes object from S3.
func (a API) DeleteObject(options *DeleteObjectOptions) (*s3.DeleteObjectOutput, error) {
	return a.Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(options.Bucket),
		Key:    aws.String(options.Path),
	})
}
