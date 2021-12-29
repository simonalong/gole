package test

import (
	"fmt"
	"github.com/simonalong/gole/config"
	"io/ioutil"
	"testing"
)

func TestLoad(t *testing.T) {
	config.LoadConfig()

	fmt.Println(config.GetValueString("a.b"))
	fmt.Println(config.GetValueBool("a.e"))
	fmt.Println(config.GetValueIntDefault("a.f", 33))
	fmt.Println(config.GetValue("a.b"))

	files, err := ioutil.ReadDir("./../resources/")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
