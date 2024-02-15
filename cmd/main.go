package main

import (
	"context"
	"log"

	"github.com/pstano1/emailVerifier/config"
	"github.com/pstano1/emailVerifier/internal/api"
	"github.com/pstano1/emailVerifier/internal/http"
	mailverifier "github.com/pstano1/emailVerifier/internal/pkg/mailVerifier"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	server := createServerFromConfig(logger)
	server.Run()
}

func createServerFromConfig(logger *logrus.Logger) *http.HTTPInstanceAPI {
	cfg, err := config.LoadFromPath(context.Background(), "pkl/config.pkl")
	if err != nil {
		panic(err)
	}

	mailVerifier := mailverifier.NewMailVerifier(
		logger.WithField("component", "MailVerifier"),
	)
	if err != nil {
		log.Fatalf("error while configuring kubernetes %v", err)
		return nil
	}

	api := api.NewInstanceAPI(&api.APIConfig{
		Logger:       logger.WithField("component", "api"),
		MailVerifier: mailVerifier,
	})

	return http.NewHTTPInstanceAPI(&http.HTTPConfig{
		Logger:      logger.WithField("component", "http"),
		InstanceAPI: api,
		Port:        cfg.Port,
	})
}
