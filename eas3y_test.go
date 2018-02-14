package eas3y

import (
	"testing"
)

type Thing struct {
	Name string
}

func (t *Thing) SaveConfig() *Config {
	return NewConfig(
		WithBucket("eas3y"),
		WithKey(t.Name),
		WithContentType("text/json"),
	)
}

func Test(t *testing.T) {
	Save(&Thing{Name: "ThisIsATest"})
}
