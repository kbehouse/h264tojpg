package main

import (
	"io/ioutil"
	"log"

	"github.com/kbehouse/h264tojpg"
)

func main() {

	h264byte, err := ioutil.ReadFile("test.h264")

	if err != nil {
		log.Fatal(err)
	}

	h264tojpg.H264tojpg(h264byte, "test.jpg")

}
