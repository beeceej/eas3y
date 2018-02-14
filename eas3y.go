package eas3y

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var sess *session.Session
var e *eas3y

func init() {
	sess = session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-1")}))
	e = new(eas3y)
	e.s3 = s3.New(sess)
}

type eas3y struct {
	item Eas3yer
	s3   *s3.S3
}

// Eas3yer defines the bucket and path where the struct should be uploaded
type Eas3yer interface {
	SaveConfig() *Config
}

// Save will save the item to s3 based on the configuration provided by the result of the S3Config() method
func Save(item Eas3yer) (err error) {
	return putItem(item)
}

func putItem(item Eas3yer) (err error) {
	var (
		b []byte
		r *s3.PutObjectOutput
	)
	cfg := item.SaveConfig()
	fmt.Println(cfg)
	b, err = marshal(cfg, item)
	if err == nil {
		fmt.Println("serialized!")
		params := &s3.PutObjectInput{
			Bucket:      &cfg.Bucket,
			Key:         &cfg.Key,
			Body:        bytes.NewReader(b),
			ContentType: aws.String(contentTypeOrDefault(cfg)),
		}
		r, err = e.s3.PutObject(params)
		fmt.Println(r)
	}
	return err
}

func marshal(cfg *Config, item Eas3yer) ([]byte, error) {
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
