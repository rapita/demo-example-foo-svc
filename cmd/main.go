package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/oklog/run"
	"github.com/rapita/demo-example-foo-svc/internal/app"
	pb "github.com/rapita/demo-example-foo-svc/pkg/api/example/v1/foo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	grpcAddr string
	httpAddr string
)

func init() {
	fs := flag.NewFlagSet("", flag.ExitOnError)
	fs.StringVar(&grpcAddr, "grpc-addr", ":6565", "grpc address")
	fs.StringVar(&httpAddr, "http-addr", ":8080", "http address")
	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// proto server
	fooServer := app.NewFooServer()

	// grpc server
	grpcServer := grpc.NewServer()

	// gateway server
	ctx := context.Background()
	mux := runtime.NewServeMux()
	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// register grpc parts
	pb.RegisterFooServer(grpcServer, fooServer)
	// register gateway parts
	if err := pb.RegisterFooHandlerFromEndpoint(ctx, mux, grpcAddr, dialOpts); err != nil {
		log.Fatal(err)
	}

	var g run.Group
	{
		grpcListener, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			log.Fatal(err)
		}
		g.Add(func() error {
			log.Printf("Serving grpc address %s", grpcAddr)
			return grpcServer.Serve(grpcListener)
		}, func(error) {
			grpcListener.Close()
		})
	}
	{
		httpListener, err := net.Listen("tcp", httpAddr)
		if err != nil {
			log.Fatal(err)
		}
		g.Add(func() error {
			log.Printf("Serving http address %s", httpAddr)
			return http.Serve(httpListener, mux)
		}, func(err error) {
			httpListener.Close()
		})
	}
	{
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}

	if err := g.Run(); err != nil {
		log.Fatal(err)
	}
}
