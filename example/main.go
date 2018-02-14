package main

import (
	"fmt"

	"github.com/beeceej/eas3y"
)

type ExampleCustom struct {
	Name string
}

type ExampleDefault struct {
	Name string
}

func (e *ExampleCustom) SaveConfig() *eas3y.Config {
	return eas3y.NewConfig(
		eas3y.WithBucket("eas3y"),
		eas3y.WithKey(fmt.Sprintf("%s.xml", e.Name)),
		eas3y.WithContentType("text/xml"),
		eas3y.AsXML(),
	)
}

func (e *ExampleDefault) SaveConfig() *eas3y.Config {
	return eas3y.NewDefaultConfig("eas3y", fmt.Sprintf("%s.json", e.Name))
}
func main() {
	var err error
	eDefault := &ExampleDefault{Name: "Default"}
	err = eas3y.Save(eDefault)
	if err != nil {
		fmt.Println(err.Error())
	}

	eCustom := &ExampleCustom{Name: "Custom"}
	err = eas3y.Save(eCustom)
	if err != nil {
		fmt.Println(err.Error())
	}
}
