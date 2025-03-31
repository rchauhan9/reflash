package main

import (
	"context"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/pkg/errors"
	"github.com/rchauhan9/reflash/monolith/services/hello"
	"github.com/rchauhan9/reflash/monolith/services/study"
	"github.com/rchauhan9/reflash/monolith/utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func realMain() int {
	ctx := context.Background()

	router := utils.NewRouter()
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))

	appContext := &utils.AppContext{
		Context: ctx,
		Router:  router,
		Logger:  logger,
	}
	hello.RegisterRoutes(appContext)
	err := study.RegisterRoutes(appContext)
	if err != nil {
		level.Error(logger).Log("err", errors.Wrap(err, "error registering study routes"))
		return 1
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	defer func() {
		if err := srv.Shutdown(ctx); err != nil {
			level.Error(logger).Log("err", errors.Wrap(err, "error shutting down http router"))
		}
	}()

	errs := make(chan error, 3)
	go func() {
		level.Info(logger).Log("transport", "http", "address", ":8080", "msg", "listening")
		if err := srv.ListenAndServe(); err != nil {
			errs <- err
		}
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- errors.New((<-c).String())
	}()
	logger.Log("terminated", <-errs)

	return 0
}

func main() {
	os.Exit(realMain())
}
