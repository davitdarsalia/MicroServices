package internal

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type AuthServer struct {
	instance *http.Server
}

func (s *AuthServer) Run(port string, handler http.Handler) error {
	s.instance = &http.Server{
		Handler:        handler,
		Addr:           fmt.Sprintf(":%s", port),
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
	}

	return s.instance.ListenAndServe()
}

func (s *AuthServer) Kill(ctx context.Context) error {
	return s.instance.Shutdown(ctx)
}
