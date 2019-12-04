package main

import (
	"github.com/golang/glog"
	"github.com/madhukirans/replayed/pkg/server"
	"github.com/madhukirans/replayed/pkg/types"
	"strconv"
	_ "net/http/pprof"
)

var config *types.ReplayedConfig

func init() {
	//log.SetOutput(new(types.LogWriter))
	config = types.GetReplayedConfig()
	glog.Info("Config: v", config)
}

func main() {
	router := server.StartServer(config)

	glog.Info("Starting server")
	router.Run(":" + strconv.Itoa(config.Port))
}
