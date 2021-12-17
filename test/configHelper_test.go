package test

import (
	"fmt"
	"github.com/simonalong/tools/config"
	"testing"
)

func TestLoad(t *testing.T) {
	config.LoadConfig()

	fmt.Println(config.GetValueString("a.b"))
	fmt.Println(config.GetValueBool("a.e"))
	fmt.Println(config.GetValueIntDefault("a.f", 33))
	fmt.Println(config.GetValueObject("a.b"))
}
