package entities

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, h http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           fmt.Sprintf(":%s", port),
		Handler:        h,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServeTLS(os.Getenv("SERVER_CERT_PATH"), os.Getenv("SERVER_KEY_PATH"))
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
