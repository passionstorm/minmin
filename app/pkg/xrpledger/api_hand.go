package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Map map[string]interface{}

type Req struct{}



func (Req) Post(url string, param Xparam) {
	reqParam, err := json.Marshal(param)
	handErr(err)
	fmt.Println(string(reqParam))
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqParam))
	handErr(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	handErr(err)
	log.Println(prettyJson(body))
}

func prettyJson(data []byte) string {
	var buff bytes.Buffer
	fmt.Println(string(data))
	err := json.Indent(&buff, data, "", "\t")
	handErr(err)
	return buff.String()
}

func handErr(err error) {
	if err != nil {
		panic(err)
	}
}