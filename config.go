package eas3y

// Config structure for eas3y, used to specify the specifics of how to save to s3
type Config struct {
	Bucket      string
	Key         string
	ContentType string
	MarshalAs   uint
}

// ConfigFunc represents a closure which sets some configuration option to the config passed in
type ConfigFunc func(*Config)

const defaultContentType = "text/json"
const defaultSerialization = "JSON"
const (
	asJSON = iota + 1
	asXML
)

// NewConfig takes a list of ConfigFuncs, applys configurations (in order provided)
// and returns to the caller a configuration which can then be used to save a struct to s3.
// Example Usage:
//
// eas3y.NewConfig(
// 		eas3y.WithBucket("eas3y"),
// 		eas3y.WithKey("key"),
// 		eas3y.WithContentType("text/json"),
// 		eas3y.AsJSON(),
// )
func NewConfig(cfgs ...ConfigFunc) *Config {
	c := new(Config)
	for _, f := range cfgs {
		f(c)
	}
	return c
}

// NewDefaultConfig will return a basic configuration with the bucket and key provided saved as JSON
func NewDefaultConfig(bucket, key string) *Config {
	return NewConfig(
		WithContentType("text/json"),
		AsJSON(),
		WithBucket(bucket),
		WithKey(key),
	)
}

// WithBucket is a configuration Option which sets the bucket to save to
func WithBucket(bucket string) ConfigFunc {
	return func(c *Config) {
		c.Bucket = bucket
	}
}

// WithKey is a configuration Option which sets the key to save to
func WithKey(key string) ConfigFunc {
	return func(c *Config) {
		c.Key = key
	}
}

// WithContentType is a configuration Option which sets the content-type of the file being saved
func WithContentType(contentType string) ConfigFunc {
	return func(c *Config) {
		c.ContentType = contentType
	}
}

// AsJSON sets MarshalAs to json, and as a side affect sets content-type to text/json
func AsJSON() ConfigFunc {
	return func(c *Config) {
		c.MarshalAs = asJSON
		c.ContentType = "text/json"
	}
}

// AsXML sets MarshalAs to XML, and as a side affect sets content-type to text/xml
func AsXML() ConfigFunc {
	return func(c *Config) {
		c.MarshalAs = asXML
		c.ContentType = "text/xml"
	}
}

func contentTypeOrDefault(cfg *Config) string {
	if cfg.ContentType == "" {
		return "text/json"
	}
	return cfg.ContentType
}

func marshalAsOrDefault(cfg *Config) uint {
	if cfg.MarshalAs == 0 {
		return asJSON
	}
	return cfg.MarshalAs
}
