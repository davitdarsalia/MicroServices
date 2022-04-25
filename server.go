package todo

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Server - Basic abstraction over http.server
type Server struct {
	httpServer *http.Server
}

// Run - Just a method of Server, which runs a server on a particular prt. Also, Run requires a basic handler struct
func (s *Server) Run(port string, h http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           fmt.Sprintf(":%s", port),
		Handler:        h,
		MaxHeaderBytes: 1 << 20,          // MaxHeaderBytes defines max amount of bytes per request in header. 1 << 20 is 1 MB
		ReadTimeout:    10 * time.Second, // Overall amount of time, that server is ready to read request body
		WriteTimeout:   10 * time.Second, // Overall amount of time of writing response for the particular request
	}
	return s.httpServer.ListenAndServe()
}

// ShutDown - Shuts down the server based on context. Note, that context must be a first parameter and named exactly: "ctx"
func (s *Server) ShutDown(ctx context.Context) {
	s.ShutDown(ctx)
}
