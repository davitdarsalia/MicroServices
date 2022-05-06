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

func (s *MainServer) Run(port string) {
	s.rootServer = &http.Server{
		Addr:           fmt.Sprintf(":%s", port),
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
	}
}

func (s *MainServer) ShutDown(ctx context.Context) {
	s.rootServer.Shutdown(ctx)
}
