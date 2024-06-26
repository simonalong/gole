package test

import (
	"context"
	"fmt"
	goleHttp "github.com/simonalong/gole/http"
	"github.com/simonalong/gole/util"
	"net/http"
	"testing"
	"unsafe"
)

type DemoHttpHook struct {
}

func (*DemoHttpHook) Before(ctx context.Context, req *http.Request) context.Context {
	return ctx
}

func (*DemoHttpHook) After(ctx context.Context, rsp *http.Response, rspCode int, rspData any, err error) {

}

func TestGetSimple(t *testing.T) {
	_, _, data, err := goleHttp.GetSimple("http://10.30.30.78:29013/api/core/license/osinfo")

	if err != nil {
		fmt.Printf("error = %v\n", err)
		return
	}
	fmt.Println("结果： " + string(data.([]byte)))

	datas := util.ToInt(unsafe.Sizeof(data))

	fmt.Println("====" + util.ToString(datas))
}
