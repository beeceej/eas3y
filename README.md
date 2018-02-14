# eas3y

Example usage:

```go
type Example struct {
  Field string
}

func (e *Example) S3Path() bucket, key{
  return YOUR_S3_BUCKET_HERE, YOUR_DESIRED_KEY_HERE
}

eas3y.Save(&Example{Field: "SomeField"})

```