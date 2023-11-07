package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/WORUS/grpc_golang/gen/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	client := pb.NewTestApiClient(conn)

	resp, err := client.Echo(context.Background(), &pb.ResponseRequest{Msg: "Hello everyone"})
	if err != nil {
		log.Println(err)
	}

	fmt.Println(resp.Msg)
}