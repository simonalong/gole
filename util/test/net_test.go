package test

import (
	"fmt"
	"github.com/simonalong/gole/util"
	"testing"
)

func TestHostIp(t *testing.T) {
	fmt.Println(util.GetIntranetIp())
	fmt.Println(util.GetPublicIP())
}

// [xx:xx:xx:xx:xx:xx xx:xx:xx:xx:xx:xx xx:xx:xx:xx:xx:xx xx:xx:xx:xx:xx:xx]
func TestMacAddress(t *testing.T) {
	fmt.Println(util.GetMacAddress())
}
