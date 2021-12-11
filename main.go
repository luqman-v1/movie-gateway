package main

import (
	"fmt"
	"movie-gateway/route"
	"movie-gateway/util"
	"os"
)

func main() {
	util.Env()
	e := route.Init()
	data, err := util.Json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		panic(fmt.Sprint(err))
	}
	fmt.Println(string(data))
	e.Start(":" + os.Getenv("port"))
}
