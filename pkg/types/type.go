package types

import (
	"encoding/json"
	"github.com/golang/glog"
	"io/ioutil"
)

type ReplayedConfig struct {
	BufferSizeInMB              int `json:"BufferSizeInMB"`
	ClientRequestBufferSizeInKB int `json:"ClientRequestBufferSizeInKB"`
	Port                        int `json:"Port"`
}

const (
	ConfigFile = "/etc/replayed/replayed.conf"
)

func GetReplayedConfig() *ReplayedConfig {
	config := &ReplayedConfig{
		BufferSizeInMB:              1000,
		ClientRequestBufferSizeInKB: 1000,
		Port:                        8080,
	}
	var jsonFile []byte
	var err error
	jsonFile, err = ioutil.ReadFile(ConfigFile)

	if err != nil {
		//glog.Fatalf("Config file reading error [ %v ]", err)
		glog.Errorf("Config file reading error [ %v ]\nLoading defaults", err)
	} else {
		err = json.Unmarshal(jsonFile, config)
		if err != nil {
			glog.Fatalf("Configfile parsing error: %v", err)
		}
	}
	glog.Info("Config loaded")
	return config
}
