package main

import (
	"context"
	"flag"
	"github.com/Victor90001/prod/internal/db" //
	"github.com/Victor90001/prod/internal/handlers" //
	"github.com/Victor90001/prod/internal/repository" //
	"github.com/Victor90001/prod/internal/services" //
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	var (
		dbUrl    string
		listen   string
		logLevel string
	)

	flag.StringVar(&dbUrl, "db", "postgres://shopus:shopus@postgres_shopus1:5432/shopus", "database connection url")
	flag.StringVar(&listen, "listen", ":8000", "server listen interface")
	flag.StringVar(&logLevel, "log-level", "debug", "log level: panic, fatal, error, warning, info, debug, trace")

	flag.Parse()

	ctx := context.Background()

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.Panicf("unable to get log level: %v", err)
	}
	logrus.SetLevel(level)

	dbc, err := db.NewPostgresPool(dbUrl)
	if err != nil {
		logrus.Panicf("unable get postgres pool: %v", err)
	}

	networkRepo, err := repository.NewPostgresNetworkRepository(dbc)
	if err != nil {
		logrus.Panicf("unable build network repo: %v", err)
	}

	dealerRepo, err := repository.NewPostgresDealerRepository(dbc)
	if err != nil {
		logrus.Panicf("unable build dealer repo: %v", err)
	}

	listRepo, err := repository.NewPostgresListRepository(dbc)
	if err != nil {
		logrus.Panicf("unable build list repo: %v", err)
	}

	userRepo, err := repository.NewPostgresUserRepository(dbc)
	if err != nil {
		logrus.Panicf("unable build user repo: %v", err)
	}

	networkService, err := services.NewNetworkService(networkRepo)
	if err != nil {
		logrus.Panicf("unable build network service: %v", err)
	}

	dealerService, err := services.NewDealerService(dealerRepo)
	if err != nil {
		logrus.Panicf("unable build dealer service: %v", err)
	}

	listService, err := services.NewListService(listRepo)
	if err != nil {
		logrus.Panicf("unable build list service: %v", err)
	}

	userService, err := services.NewUserService(userRepo)
	if err != nil {
		logrus.Panicf("unable build user service: %v", err)
	}

	g := gin.New()

	_, err = handlers.NewNetworkHandlers(g, networkService)
	if err != nil {
		logrus.Panicf("unable build slug handlers: %v", err)
	}

	_, err = handlers.NewDealerHandlers(g, dealerService)
	if err != nil {
		logrus.Panicf("unable build dealer handlers: %v", err)
	}

	_, err = handlers.NewListHandlers(g, listService)
	if err != nil {
		logrus.Panicf("unable build list handlers: %v", err)
	}

	_, err = handlers.NewUserHandlers(g, userService)
	if err != nil {
		logrus.Panicf("unable build slug handlers: %v", err)
	}

	doneC := make(chan error)

	go func() { doneC <- g.Run(listen) }()

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGABRT, syscall.SIGHUP, syscall.SIGTERM)

	childCtx, cancel := context.WithCancel(ctx)
	go func() {
		sig := <-signalChan
		logrus.Debugf("exiting with signal: %v", sig)
		cancel()
	}()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				doneC <- ctx.Err()
			}
		}
	}(childCtx)

	<-doneC

}
