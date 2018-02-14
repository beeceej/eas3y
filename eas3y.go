package eas3y

import (
	"bytes"
	"encoding/json"
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
	e.s3 = _s3()
}

func _s3() *s3.S3 {
	return s3.New(sess)
}

func Save(item Eas3yer) error {
	bucket, key := item.S3Path()
	b, err := json.Marshal(item)
	params := &s3.PutObjectInput{
		Bucket:      &bucket,
		Key:         &key,
		Body:        bytes.NewReader(b),
		ContentType: aws.String("text/json"),
	}
	_, err = e.s3.PutObject(params)
	return err
}

type eas3y struct {
	item Eas3yer
	s3   *s3.S3
}

func (e *eas3y) save() error {
	bucket, key := e.item.S3Path()
	fmt.Println(bucket, key)
	return nil
}

// Eas3yer defines the bucket and path where the struct should be uploaded
type Eas3yer interface {
	S3Path() (bucket string, key string)
}
