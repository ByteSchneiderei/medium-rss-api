package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"

	"byteschneiderei.com/intern/medium/api/pkg/config"
	"byteschneiderei.com/intern/medium/api/pkg/medium"
	"byteschneiderei.com/intern/medium/api/pkg/version"
)

// VERSION indictes the current version
var VERSION string

func main() {
	env := config.EnvVar{}
	if err := env.Parse(); err != nil {
		log.Fatal("Error on parsing environment variable: ", err)
	}

	c := cache.New(config.CacheDefaultExpiration, config.CacheCleanupInterval)

	logLevel := log.DebugLevel
	if env.GoEnv == "production" {
		logLevel = log.InfoLevel
	}
	log.SetLevel(logLevel)
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = config.LogTimestampFormat
	log.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = config.LogFulltimestamp

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat(config.HeartbeatEndpoint))
	r.Use(middleware.Timeout(config.APITimeoutDuration))
	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	versionPkg := version.New(VERSION)
	mediumPkg := medium.New(&env, c)

	r.Get("/version", versionPkg.Handler())
	r.Get("/medium", mediumPkg.Handler())

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", env.Port),
		Handler: r,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Error on starting server: ", err)
		}
	}()
	log.Info(fmt.Sprintf("Server started on port %d with version %s", env.Port, VERSION))

	<-done
	log.Info("Server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), config.ServerGracefulShutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Error on server shutdown: ", err)
	}
	log.Info("Server exited gracefully")
}
