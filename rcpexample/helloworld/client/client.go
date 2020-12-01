package main

import (
	"context"
	"fmt"
	"go_study/rcpexample/helloworld/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		fmt.Println("grpc.Dial failed:", err)
		return
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	reply, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "Haolipeng"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", reply.Message)
}
