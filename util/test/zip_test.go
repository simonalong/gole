package test

import (
	"fmt"
	"github.com/simonalong/gole/util"
	"testing"
)

func TestZip(t *testing.T) {
	err := util.Zip("./resources", "resources.zip")
	if err != nil {
		fmt.Println("报错")
		return
	}
}

func TestUnZip(t *testing.T) {
	err := util.UnZip("resources.zip", "./resources2")
	if err != nil {
		fmt.Println("报错")
		return
	}
}
