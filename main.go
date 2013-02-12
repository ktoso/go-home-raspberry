package main

import (
	"bytes"
	"log"
	"os/exec"
	"strings"

	"code.google.com/p/gorest"
	"net/http"
)

var port = ":5555"

func main() {
	gorest.RegisterService(new(HelloService)) //Register our service
	http.Handle("/", gorest.Handle())
	http.ListenAndServe(port, nil)
}

//Service Definition
type HelloService struct {
	gorest.RestService `root:"/gohomepi/"`
	enable             gorest.EndPoint `method:"GET" path:"/enable/{name:string}" output:"string"`
	enableAfter        gorest.EndPoint `method:"GET" path:"/enable/{name:string}/after/{s:int}" output:"string"`

	disable      gorest.EndPoint `method:"GET" path:"/disable/{name:string}" output:"string"`
	disableAfter gorest.EndPoint `method:"GET" path:"/disable/{name:string}/after/{s:int}" output:"string"`
}

func (serv HelloService) Enable(name string) string {
	log.Print("Will [enable] device [", name, "]")

	c := make(chan string)
	go switchState("--on", name, c)
	go logOnSuccess(name, c)

	return name
}

func (serv HelloService) EnableAfter(name string, s int) string {
	log.Print("Will [enable] device [", name, "] in [", s, "] seconds...")

	c := make(chan string)
	go switchState("--on", name, c)
	go logOnSuccess(name, c)

	return "enable scheduled"
}

func (serv HelloService) Disable(name string) string {
	log.Print("Will [disable] device [", name, "]")

	c := make(chan string)
	go switchState("--on", name, c)
	go logOnSuccess(name, c)

	return name
}

func (serv HelloService) DisableAfter(name string, s int) string {
	log.Print("Will [disable] device [", name, "] in [", s, "] seconds...")

	c := make(chan string)
	go switchState("--on", name, c)
	go logOnSuccess(name, c)

	return "disable scheduled"
}

func logOnSuccess(name string, c chan string) {
	for msg := range c {
		log.Print(name, " reported: ", msg)
	}
}

func switchState(state string, name string, c chan string) string {
	cmd := exec.Command("tdtool", state, name)
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	var outString = out.String()
	c <- outString

	return outString
}
