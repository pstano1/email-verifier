package mailverifier

import (
	"bufio"
	"net"
	"os"
	"regexp"
	"strings"

	"github.com/pstano1/emailVerifier/internal/pkg"
	"github.com/sirupsen/logrus"
)

type MailVerifier struct {
	log     logrus.FieldLogger
	scanner *bufio.Scanner
}

type IMailVerifierProvider interface {
	ValidateEmailAddress(emailAddress string) bool
	VerifyDomain(domain string) (pkg.DomainVerifierResponse, error)
}

func NewMailVerifier(
	log logrus.FieldLogger,
) IMailVerifierProvider {
	scanner := bufio.NewScanner(os.Stdin)
	return &MailVerifier{
		log:     log,
		scanner: scanner,
	}
}

func (mv MailVerifier) ValidateEmailAddress(emailAddress string) bool {
	isValid := false
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	mv.log.Debugf("starting validation of an e-mail address: %s", emailAddress)
	if emailRegex.MatchString(emailAddress) {
		isValid = true
	}

	return isValid
}

func (mv MailVerifier) VerifyDomain(domain string) (pkg.DomainVerifierResponse, error) {
	mv.log.Debugf("veryfing domain: %s", domain)

	response := &pkg.DomainVerifierResponse{
		IsValid:     false,
		HasMX:       false,
		HasSPF:      false,
		HasDMARC:    false,
		SpfRecord:   "",
		DmarcRecord: "",
	}
	var err error

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		mv.log.Errorf("got error when checking MX records, error: %v", err)
		return *response, err
	}
	if len(mxRecords) > 0 {
		response.HasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		mv.log.Errorf("got error when checking SPF records, error: %v", err)
		return *response, err
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			response.HasSPF = true
			response.SpfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		mv.log.Errorf("got error when checking DMARC records, error: %v", err)
		return *response, err
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			response.HasDMARC = true
			response.DmarcRecord = record
			break
		}
	}

	if response.HasMX && response.HasSPF && response.HasDMARC {
		response.IsValid = true
	}

	return *response, err
}
