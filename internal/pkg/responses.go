package pkg

type DomainVerifierResponse struct {
	IsValid      bool   `json:"isValid"`
	HasMX        bool   `json:"hasMX"`
	HasSPF       bool   `json:"hasSPF"`
	HasDMARC     bool   `json:"hasDMARC"`
	SpfRecord    string `json:"spfRecord"`
	DmarcRecord  string `json:"dmarcRecord"`
	EmailAddress string `json:"emailAddress"`
}
