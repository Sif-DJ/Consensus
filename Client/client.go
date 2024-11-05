package Client

import (
	proto "Consensus/grpc"
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type NodeServiceServer struct {
	proto.UnimplementedNodeServiceServer
	hasBaton bool
}

func (srv *NodeServiceServer) PassBaton(ctx context.Context, req *proto.Empty) (*proto.Empty, error) {
	srv.hasBaton = true
	return &proto.Empty{}, nil
}

func startServer(address string, neighbour string, baton bool) {
	srv := &NodeServiceServer{
		hasBaton: baton,
	}

	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":"+address)
	if err != nil {
		log.Fatalln("Exception Error")
	}
	proto.RegisterNodeServiceServer(grpcServer, srv)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Exception Error after Registration")
	}

	conn, err := grpc.NewClient("localhost:"+neighbour, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Exception Error")
	}

	client := proto.NewNodeServiceClient(conn)

	for {
		if srv.hasBaton {
			srv.hasBaton = false
			fmt.Println(address + " passing baton to " + neighbour)
			client.PassBaton(context.Background(), &proto.Empty{})
			time.Sleep(5 * time.Second)
		} else {
			fmt.Println(address + " is sleeping")
			time.Sleep(5 * time.Second)
		}
	}
}
