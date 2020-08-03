package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// LoadFile 下载文件到本地
func LoadFile(url string, local string) {

	file, err := os.OpenFile(local, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal("open file failed, filename: ", local, " ", err)
		panic(err)
	}
	defer file.Close()

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("create request failed: ", url, " ", err)
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("get responce failed: ", url, " ", err)
		panic(err)
	}

	defer resp.Body.Close()
	rBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("read responce failed: ", url, " ", err)
		panic(err)
	}
	file.Write(rBody)
}
