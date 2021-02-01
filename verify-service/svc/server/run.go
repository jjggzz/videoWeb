// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 84830579f0
// Version Date: Thu Dec 31 05:37:46 UTC 2020

package server

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"os"

	// 3d Party
	"google.golang.org/grpc"

	// This Service
	"videoWeb/verify-service/handlers"
	pb "videoWeb/verify-service/proto"
	"videoWeb/verify-service/svc"
)

var DefaultConfig Config

func init() {
	flag.StringVar(&DefaultConfig.DebugAddr, "debug.addr", ":5060", "Debug and metrics listen address")
	flag.StringVar(&DefaultConfig.HTTPAddr, "http.addr", ":5050", "HTTP listen address")
	flag.StringVar(&DefaultConfig.GRPCAddr, "grpc.addr", ":5040", "gRPC (HTTP) listen address")

	// Use environment variables, if set. Flags have priority over Env vars.
	if addr := os.Getenv("DEBUG_ADDR"); addr != "" {
		DefaultConfig.DebugAddr = addr
	}
	if port := os.Getenv("PORT"); port != "" {
		DefaultConfig.HTTPAddr = fmt.Sprintf(":%s", port)
	}
	if addr := os.Getenv("HTTP_ADDR"); addr != "" {
		DefaultConfig.HTTPAddr = addr
	}
	if addr := os.Getenv("GRPC_ADDR"); addr != "" {
		DefaultConfig.GRPCAddr = addr
	}
}

// Config contains the required fields for running a server
type Config struct {
	HTTPAddr  string
	DebugAddr string
	GRPCAddr  string
}

func NewEndpoints(service pb.VerifyServer) svc.Endpoints {
	// Business domain.

	// Wrap Service with middlewares. See handlers/middlewares.go
	service = handlers.WrapService(service)

	// Endpoint domain.
	var (
		sendverifycodeEndpoint  = svc.MakeSendVerifyCodeEndpoint(service)
		checkverifycodeEndpoint = svc.MakeCheckVerifyCodeEndpoint(service)
	)

	endpoints := svc.Endpoints{
		SendVerifyCodeEndpoint:  sendverifycodeEndpoint,
		CheckVerifyCodeEndpoint: checkverifycodeEndpoint,
	}

	// Wrap selected Endpoints with middlewares. See handlers/middlewares.go
	endpoints = handlers.WrapEndpoints(endpoints)

	return endpoints
}

// Run starts a new http server, gRPC server, and a debug server with the
// passed config and logger
func Run(cfg Config) {
	service := handlers.NewService()
	endpoints := NewEndpoints(service)

	// Mechanical domain.
	errc := make(chan error)

	// Interrupt handler.
	go handlers.InterruptHandler(errc)

	// Debug listener.
	go func() {
		log.Println("transport", "debug", "addr", cfg.DebugAddr)

		m := http.NewServeMux()
		m.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
		m.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
		m.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
		m.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
		m.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))

		errc <- http.ListenAndServe(cfg.DebugAddr, m)
	}()

	// HTTP transport.
	go func() {
		log.Println("transport", "HTTP", "addr", cfg.HTTPAddr)
		h := svc.MakeHTTPHandler(endpoints)
		errc <- http.ListenAndServe(cfg.HTTPAddr, h)
	}()

	// gRPC transport.
	go func() {
		log.Println("transport", "gRPC", "addr", cfg.GRPCAddr)
		ln, err := net.Listen("tcp", cfg.GRPCAddr)
		if err != nil {
			errc <- err
			return
		}

		srv := svc.MakeGRPCServer(endpoints)
		s := grpc.NewServer()
		pb.RegisterVerifyServer(s, srv)

		errc <- s.Serve(ln)
	}()

	// Run!
	log.Println("exit", <-errc)
}
