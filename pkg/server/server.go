package server

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/madhukirans/replayed/pkg/types"
	"io/ioutil"

	"github.com/golang/glog"
	"sync"
	"net/http"
)

var config *types.ReplayedConfig
var serverBuf *bytes.Buffer
var clientRequestBufferSize int
var mutex sync.RWMutex
var serverBufferSize int

func StartServer(c *types.ReplayedConfig) *gin.Engine {
	config = c
	serverBufferSize = c.BufferSizeInMB * 1024 * 1024
	clientRequestBufferSize = c.ClientRequestBufferSizeInKB * 1024

	serverBuf = new(bytes.Buffer)
	r := gin.Default()
	v1 := r.Group("/")
	{
		v1.GET("/", GetData)
		v1.POST("/", PostData)
	}

	return r
}

func PostData(c *gin.Context) {
	//glog.Info("Reading client \n")
	x, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		glog.Errorf("Reading error %v", err)
	} else {
		mutex.Lock()
		if serverBuf.Len() + len(x) < serverBufferSize {
			serverBuf.Write(x)
			c.JSON(http.StatusAccepted, gin.H{"status": "Ok"})
		} else {
			c.JSON(http.StatusInsufficientStorage, gin.H{"status": "StatusInsufficientStorage"})
		}
		mutex.Unlock()
	}
}

func GetData(c *gin.Context) {
	mutex.RLock()
	c.String(200,  serverBuf.String())
	mutex.RUnlock()
}
