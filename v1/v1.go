package v1

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/beeceej/eas3y"
)

// Save will save the item to s3 based on the configuration provided by the result of the S3Config() method
func Save(svc s3iface.S3API, item eas3y.S3Writer) (*s3.PutObjectOutput, error) {
	return putItem(svc, item)
}

func putItem(svc s3iface.S3API, item eas3y.S3Writer) (*s3.PutObjectOutput, error) {
	var (
		b   []byte
		err error
	)

	cfg := item.SaveConfig()

	cfg.Key = cfg.FormatKey()

	if b, err = eas3y.Marshal(cfg, item); err != nil {
		return nil, err
	}

	params := &s3.PutObjectInput{
		Bucket:      &cfg.Bucket,
		Key:         &cfg.Key,
		Body:        bytes.NewReader(b),
		ContentType: aws.String(eas3y.ContentTypeOrDefault(cfg)),
	}

	return svc.PutObject(params)
}
