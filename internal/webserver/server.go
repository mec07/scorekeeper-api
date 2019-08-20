package webserver

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mec07/rununtil"
	"github.com/rs/zerolog/log"
)

// NewRunner returns a function that runs the http webserver
func NewRunner() rununtil.RunnerFunc {
	return func() rununtil.ShutdownFunc {
		r := chi.NewRouter()
		r.Get("/ping", pingHandler)

		httpServer := http.Server{Addr: ":8080", Handler: r}

		go runHTTPServer(&httpServer)

		return func() {
			if err := httpServer.Shutdown(context.Background()); err != nil {
				log.Error().Err(err).Msg("error shutting down http server")
			}
		}
	}
}

func runHTTPServer(srv *http.Server) {
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("ListenAndServe")
	}
}
