package s3

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
)

//nolint:lll // go generate is ugly.
//go:generate mockgen -destination=mocks/api_mock.go -package=mocks github.com/omegion/vault-unseal/internal/vault APIInterface
// APIInterface is an interface for API.
type APIInterface interface {
	GetObject(options *GetObjectOptions) (io.ReadCloser, error)
	PutObject(options *PutObjectOptions) (*s3.PutObjectOutput, error)
	DeleteObject(options *DeleteObjectOptions) (*s3.DeleteObjectOutput, error)
}

// API is main struct of S3.
type API struct {
	Config *aws.Config
	Client *s3.Client
}

type PutObjectOptions struct {
	Bucket,
	Path string
	Value       io.Reader
	EncodedTags string
}

type GetObjectOptions struct {
	Bucket,
	Path string
}

type DeleteObjectOptions struct {
	Bucket,
	Path string
}

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

func (a API) GetObject(options *GetObjectOptions) (io.ReadCloser, error) {
	resp, err := a.Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(options.Bucket),
		Key:    aws.String(options.Path),
	})
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func (a API) PutObject(options *PutObjectOptions) (*s3.PutObjectOutput, error) {
	return a.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:  aws.String(options.Bucket),
		Key:     aws.String(options.Path),
		Body:    options.Value,
		Tagging: aws.String(options.EncodedTags),
	})
}

func (a API) DeleteObject(options *DeleteObjectOptions) (*s3.DeleteObjectOutput, error) {
	return a.Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(options.Bucket),
		Key:    aws.String(options.Path),
	})
}
