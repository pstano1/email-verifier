package api

import (
	"errors"
	"strings"

	"github.com/pstano1/emailVerifier/internal/pkg"
	mailverifier "github.com/pstano1/emailVerifier/internal/pkg/mailVerifier"
	"github.com/sirupsen/logrus"
)

type InstanceAPI struct {
	log          logrus.FieldLogger
	mailVerifier mailverifier.IMailVerifierProvider
}

type APIConfig struct {
	Logger       logrus.FieldLogger
	MailVerifier mailverifier.IMailVerifierProvider
}

func NewInstanceAPI(conf *APIConfig) *InstanceAPI {
	return &InstanceAPI{
		log:          conf.Logger,
		mailVerifier: conf.MailVerifier,
	}
}

func (a *InstanceAPI) VerifyEmailAddress(emailAddresses []string) ([]pkg.DomainVerifierResponse, error) {
	a.log.Debugf("starting email addresses verification: %d", len(emailAddresses))

	var response = make([]pkg.DomainVerifierResponse, len(emailAddresses))
	var err error

	for index, address := range emailAddresses {
		if ok := a.mailVerifier.ValidateEmailAddress(address); !ok {
			response[index].IsValid = false
		}
		emailArray := strings.SplitN(address, "@", 2)
		domain := ""
		if len(emailArray) > 1 {
			domain = emailArray[1]
		} else {
			return response, errors.New("something went wrong when splitting an address")
		}
		response[index], err = a.mailVerifier.VerifyDomain(domain)
		if err != nil {
			response[index].IsValid = false
		}
		response[index].EmailAddress = address
	}

	return response, err
}
