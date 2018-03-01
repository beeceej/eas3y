package eas3y

import (
	"bytes"
	"testing"
)

type ThingJSON struct {
	Name string
}

func (t *ThingJSON) SaveConfig() *Config {
	return NewConfig(
		WithBucket("eas3y"),
		WithKey(t.Name),
		WithContentType("text/json"),
	)
}

func Test_Eas3y(t *testing.T) {
	testMarshalexpectBytes := func() {

		thing := &ThingJSON{Name: "ThisIsATest"}
		expectedBytes := []byte{123, 34, 78, 97, 109, 101, 34, 58, 34, 84, 104, 105, 115, 73, 115, 65, 84, 101, 115, 116, 34, 125}
		b, _ := marshal(thing.SaveConfig(), thing)
		if !bytes.Equal(b, expectedBytes) {
			t.Fail()
		}
	}

	testMarshalexpectBytes()
}
