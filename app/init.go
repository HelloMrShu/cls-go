package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var Conf *Config

type HotPlate struct {
	Host       string `json:"host"`
	Path       string `json:"path"`
	Refer      string `json:"refer"`
	Categories string `json:"categories"`
}

type NewsPlate struct {
	Host  string `json:"host"`
	Path  string `json:"path"`
	Refer string `json:"refer"`
}

type ClsInfo struct {
	Hot  HotPlate  `json:"hot"`
	News NewsPlate `json:"news"`
}

type Config struct {
	Webhook     string  `json:"webhook"`
	Cls         ClsInfo `json:"cls"`
	Ua          string  `json:"ua"`
	ContentType string  `json:"content-type"`
	Charset     string  `json:"charset"`
}

func NewCycle() {
	filePath := "./conf/local.json"

	var (
		data []byte
		err  error
	)
	data, err = ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read conf file err", err)
		return
	}

	er := json.Unmarshal(data, &Conf)
	if er != nil {
		fmt.Println("json unMarsha1 conf file err", er)
		return
	}
}
