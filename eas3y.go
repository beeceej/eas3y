package eas3y

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/s3"
)

// S3Writer defines the bucket and path where the struct should be uploaded
type S3Writer interface {
	SaveConfig() *Config
}

// Save will save the item to s3 based on the configuration provided by the result of the S3Config() method
func Save(svc *s3.S3, item S3Writer) (*s3.PutObjectOutput, error) {
	return putItem(svc, item)
}

func putItem(svc *s3.S3, item S3Writer) (*s3.PutObjectOutput, error) {
	var (
		b   []byte
		err error
	)

	cfg := item.SaveConfig()

	cfg.Key = cfg.formatKey()

	if b, err = marshal(cfg, item); err != nil {
		return nil, err
	}

	params := &s3.PutObjectInput{
		Bucket:      &cfg.Bucket,
		Key:         &cfg.Key,
		Body:        bytes.NewReader(b),
		ContentType: aws.String(contentTypeOrDefault(cfg)),
	}

	return svc.PutObject(params)
}

func marshal(cfg *Config, item S3Writer) ([]byte, error) {
	marshalTo := marshalAsOrDefault(cfg)
	switch marshalTo {
	case asJSON:
		return json.Marshal(item)
	case asXML:
		return xml.Marshal(item)
	default:
		return nil, errors.New("Unsupported serialization strategy")
	}
}
