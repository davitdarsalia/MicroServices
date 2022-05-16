package entities

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type MainServer struct {
	rootServer *http.Server
}

func (s *MainServer) Run(port string, handler http.Handler) error {
	s.rootServer = &http.Server{
		Handler:        handler,
		Addr:           fmt.Sprintf(":%s", port),
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
	}

	return s.rootServer.ListenAndServe()
}

func (s *MainServer) ShutDown(ctx context.Context) error {
	err := s.rootServer.Shutdown(ctx)
	return err
}
