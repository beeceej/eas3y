package eas3y

import (
	"fmt"
	"testing"
)

type Thing struct {
	Name string
}

func (t *Thing) S3Path() (string, string) {
	return "eas3y", fmt.Sprintf("%s.json", t.Name)
}

func Test(t *testing.T) {
	Save(&Thing{Name: "ThisIsATest"})
}
