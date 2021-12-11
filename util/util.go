package util

import (
	"log"
	"strconv"

	"google.golang.org/grpc"

	jsoniter "github.com/json-iterator/go"

	"github.com/joho/godotenv"
)

var Json = jsoniter.ConfigCompatibleWithStandardLibrary

func Env(path ...string) (err error) {
	err = godotenv.Load(path...)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return
}

func Dial(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", addr, err)
	}
	return conn
}

func StoI(input string) int32 {
	i, err := strconv.Atoi(input)
	if err != nil {
		log.Println("error convert string to int", err)
	}
	return int32(i)
}
