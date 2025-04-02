package study

import (
	"github.com/go-kit/log"
	"github.com/pkg/errors"
	"github.com/rchauhan9/reflash/monolith/common/clients/card_creator"
	"github.com/rchauhan9/reflash/monolith/common/database"
	"github.com/rchauhan9/reflash/monolith/config"
	"github.com/rchauhan9/reflash/monolith/utils"
)

func InitialiseService(appContext *utils.AppContext, config *config.Config) func() {
	migrationPath := config.StudyService.Database.MigrationPath
	databaseURL := config.StudyService.Database.URL

	cleanup := func() {}

	migrator, err := database.NewMigrator(databaseURL, migrationPath, appContext.Logger)
	if err != nil {
		panic(errors.Wrapf(err, "unable to create study service migrator"))
	}
	if err = migrator.MigrateDb(); err != nil {
		panic(errors.Wrapf(err, "unable to migrate study service database"))
	}
	if err = migrator.Close(); err != nil {
		panic(errors.Wrapf(err, "unable to close study service migrator"))
	}

	pool, err := NewDatabasePool(appContext.Context, databaseURL)
	if err != nil {
		panic(errors.Wrapf(err, "failed to create database pool"))
	}
	rep := NewRepository(pool)
	createCardClient := card_creator.NewClient()
	svc := NewService(rep, createCardClient)
	studyLogger := log.WithSuffix(appContext.Logger, "svc", "study")

	err = RegisterRoutes(svc, appContext.Router, studyLogger)
	if err != nil {
		panic(errors.Wrapf(err, "failed to register study service routes"))
	}
	studyLogger.Log("msg", "initialised service")

	cleanup = func() {
		studyLogger.Log("msg", "cleaning up study service")
		pool.Close()
	}

	return cleanup
}
