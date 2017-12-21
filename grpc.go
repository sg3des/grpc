//go:generate protoc -I ./pb --go_out=plugins=grpc:./pb ./pb/pb.proto

package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/sg3des/grpc/pb"

	"google.golang.org/grpc"
)

var reply = &pb.HelloReply{Message: "Hello"}

func main() {

}

type server struct{}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return reply, nil
}

func Server(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	return s.Serve(l)
}

//Client connect to GRPC server, send message, and print result
func Client(addr string) error {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "CLIENT"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Greeting: %s", r.Message)
	return nil
}
