package eas3y

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	testFormatKey := func(contentType, key, expectedSuffix string) {
		c := Config{ContentType: contentType, Key: key}
		var expectedKey string
		if expectedSuffix != "" {
			expectedKey = fmt.Sprintf("%s.%s", key, expectedSuffix)
		} else {
			expectedKey = key
		}

		actualKey := c.formatKey()
		if expectedKey != actualKey {
			t.Fail()
		}
	}
	testFormatKey("application/json", "a/b/c/d", "json")
	testFormatKey("text/json", "a/b/c/d.json", "")
	testFormatKey("text/xml", "a/b/c/d.xml", "")
	testFormatKey("application/xml", "a/b/c/d", "xml")
}
