package main

import (
	"context"
	"testing"
	"time"

	"github.com/sg3des/grpc/pb"

	"google.golang.org/grpc"
)

const testAddr = "127.0.0.1:12345"

func TestServer(t *testing.T) {
	go func(t *testing.T) {
		err := Server(testAddr)
		if err != nil {
			t.Fatal(err)
		}
	}(t)
}

func TestClient(t *testing.T) {
	err := Client(testAddr)
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkGRPC(b *testing.B) {
	b.StopTimer()
	go Server(testAddr)
	time.Sleep(1 * time.Second)

	conn, err := grpc.Dial(testAddr, grpc.WithInsecure())
	if err != nil {
		b.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	req := &pb.HelloRequest{Name: "CLIENT"}

	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		c.SayHello(context.Background(), req)
	}
}

func BenchmarkGRPC_Connect(b *testing.B) {
	b.StopTimer()
	go Server(testAddr)
	time.Sleep(1 * time.Second)

	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		conn, _ := grpc.Dial(testAddr, grpc.WithInsecure())

		c := pb.NewGreeterClient(conn)
		req := &pb.HelloRequest{Name: "CLIENT"}

		c.SayHello(context.Background(), req)
		conn.Close()
	}
}
