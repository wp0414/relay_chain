package main

import (
	fab_proto "github.com/golang/protobuf/proto"
	g_proto "google.golang.org/protobuf/proto"
	"log"
	"relay-chain/module/types/pb"
)

func main() {
	a := pb.Prof{
		Key:    "1",
		Result: "2",
	}
	c, err := fab_proto.Marshal(&a)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(c)

	d, err := g_proto.Marshal(&a)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(d)
}
