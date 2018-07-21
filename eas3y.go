package eas3y

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

// S3Writer defines the bucket and path where the struct should be uploaded
type S3Writer interface {
	SaveConfig() *Config
}

// Marshal takes an item and marshals it as per the config
func Marshal(cfg *Config, item S3Writer) ([]byte, error) {
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

// ContentTypeOrDefault takes a configuration and if the content-type is not set defaults to text/json
func ContentTypeOrDefault(cfg *Config) string {
	if cfg.ContentType == "" {
		return "text/json"
	}
	return cfg.ContentType
}
