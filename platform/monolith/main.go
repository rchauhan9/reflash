package main

import (
	"context"
	"github.com/go-kit/log/level"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"github.com/rchauhan9/reflash/monolith/config"
	"github.com/rchauhan9/reflash/monolith/logging"
	"github.com/rchauhan9/reflash/monolith/services/hello"
	"github.com/rchauhan9/reflash/monolith/services/study"
	"github.com/rchauhan9/reflash/monolith/utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func realMain() int {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	router := utils.NewRouter()
	logger := logging.GetLogger()

	conf := config.GetConfig()

	appContext := &utils.AppContext{
		Context: ctx,
		Router:  router,
		Logger:  logger,
	}
	hello.RegisterRoutes(appContext)
	studyServiceCleanup := study.InitialiseService(appContext, conf)

	srv := &http.Server{
		Addr:    conf.Server.HTTPAddress,
		Handler: router,
	}

	defer func() {
		level.Info(logger).Log("msg", "cleaning up monolith")
		if err := srv.Shutdown(ctx); err != nil {
			level.Error(logger).Log("err", errors.Wrap(err, "error shutting down http router"))
		}
		studyServiceCleanup()
	}()

	errs := make(chan error, 3)
	go func() {
		level.Info(logger).Log("transport", "http", "address", conf.Server.HTTPAddress, "msg", "listening")
		if err := srv.ListenAndServe(); err != nil {
			errs <- err
		}
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		sig := <-c
		level.Info(logger).Log("msg", "received signal", "signal", sig)
		errs <- nil
	}()
	if err := <-errs; err != nil {
		level.Error(logger).Log("terminated", err)
	} else {
		level.Info(logger).Log("msg", "terminated via signal")
	}

	return 0
}

func main() {
	os.Exit(realMain())
}
