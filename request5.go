package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/KeThichDua/ex5go/db"
	"github.com/KeThichDua/ex5go/rpc"
	"google.golang.org/grpc"
)

var usv UserPartner5

// Run5 Tài liệu [grpc](https://grpc.io/docs/what-is-grpc/core-concepts/)
// Tạo 1 service gen code. Tạo 1 grpc server với message `UserPartner`. Nhằm getlist, create, update
// Tạo 1 grpc client để thực hiện
func Run5() {
	// ket noi mysql
	var db db.Database
	err := db.Connect("mysql", "root:1@tcp(0.0.0.0:3306)/ex5go")
	ThrowError(err)
	defer db.Engine.Close()
	ctx = context.Background()
	usv = UserPartner5{Db: db}

	go server.Start()
	if err = RunGrpcServer(ctx, usv); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func RunGrpcServer(ctx context.Context, usv UserPartner5) error {
	listen, err := net.Listen("tcp", ":3001")
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	rpc.RegisterUserPartnerService5Server(server, &usv)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println("starting gRPC server...")
	return server.Serve(listen)
}
