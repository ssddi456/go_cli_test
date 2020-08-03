package main

import (
	"os"
	"path"
	"testing"
)

func TestLoadFile(t *testing.T) {
	outputFileName := path.Join("D:/projects/go_cli_test", "test.html")
	t.Log("outputfile ", outputFileName)
	LoadFile("https://www.baidu.com/", outputFileName)

	t.Log("test finish")

	os.Remove(outputFileName)
}
