package main

import (
	"context"
	"flag"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"github.com/rchauhan9/reflash/monolith/common/configutil"
	"github.com/rchauhan9/reflash/monolith/config"
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

	configPath := flag.String("config-dir", "./config", "Directory containing config.yml")
	flag.Parse()

	router := utils.NewRouter()
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))

	var conf *config.Config
	err := configutil.LoadConfig(*configPath, logger, &conf)
	if err != nil {
		return 1
	}

	appContext := &utils.AppContext{
		Context: ctx,
		Router:  router,
		Logger:  logger,
	}
	hello.RegisterRoutes(appContext)
	studyServiceCleanup := study.InitialiseService(appContext, conf)
	defer studyServiceCleanup()

	srv := &http.Server{
		Addr:    conf.Server.HTTPAddress,
		Handler: router,
	}

	defer func() {
		if err := srv.Shutdown(ctx); err != nil {
			level.Error(logger).Log("err", errors.Wrap(err, "error shutting down http router"))
		}
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
		signal.Notify(c, syscall.SIGINT)
		errs <- errors.New((<-c).String())
	}()
	logger.Log("terminated", <-errs)

	return 0
}

func main() {
	os.Exit(realMain())
}
