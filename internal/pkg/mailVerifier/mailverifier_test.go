package mailverifier

import (
	"testing"

	"github.com/sirupsen/logrus"
)

var (
	mv IMailVerifierProvider
)

func TestMain(m *testing.M) {
	log := logrus.New()
	mv = NewMailVerifier(log)

	m.Run()
}

func TestValidateEmailAddress(t *testing.T) {
	var tests = []struct {
		name  string
		email string
		want  bool
	}{
		{"empty address", "", false},
		{"just domain", "example.com", false},
		{"@ & a domain", "@example.com", false},
		{"without top-level domain", "pstano@example", false},
		{"without domain", "pstano", false},
		{"without 2nd level domain", "pstano@.com", false},
		{"proper address #1", "pstano@example.com", true},
		{"proper address #2", "p.stano@example.eu", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := mv.ValidateEmailAddress(test.email)
			if res != test.want {
				t.Errorf("got %t, want %t", res, test.want)
			}
		})
	}
}

func TestVerifyDomain(t *testing.T) {
	var tests = []struct {
		name   string
		domain string
		want   bool
	}{
		{"empty address", "", false},
		{"without top-level domain", "examplecom", false},
		{"just top-level domain", ".com", false},
		{"proper domain #1", "outlook.com", true},
		{"proper domain #2", "gmail.com", true},
		{"proper domain #3", "ron.mil.pl", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, _ := mv.VerifyDomain(test.domain)
			if res.IsValid != test.want {
				t.Errorf("got %t, want %t", res.IsValid, test.want)
			}
		})
	}
}
