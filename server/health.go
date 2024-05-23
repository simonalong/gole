package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/simonalong/gole/config"

	"github.com/gin-gonic/gin"
	h2 "github.com/simonalong/gole/http"
	t2 "github.com/simonalong/gole/time"
)

var procId = os.Getpid()
var startTime = time.Now().Format(t2.FmtYMdHms)

const defaultVersion = "unknown"

func healthSystemStatus(c *gin.Context) {
	c.Data(http.StatusOK, h2.ContentTypeJson, []byte(fmt.Sprintf(`{"status":"ok","running":true,"pid":%d,"startupAt":"%s","version":"%s"}`, procId, startTime, getVersion())))
}

func healthSystemInit(c *gin.Context) {
	c.Data(http.StatusOK, h2.ContentTypeText, []byte(`{"status":"ok"}`))
}

func healthSystemDestroy(c *gin.Context) {
	c.Data(http.StatusOK, h2.ContentTypeText, []byte(`{"status":"ok"}`))
}

func getVersion() string {
	return config.GetValueStringDefault("base.server.version", defaultVersion)
}
