// Package eas3y allows painless upload to S3 of structs in Golang.
// Given a valid aws S3 client / session, and a struct which implements the S3Writer interface
// all a user needs to do; to write that struct to S3 is call eas3y.Save(s3 :: *s3.S3, item :: S3Writer)
package eas3y
