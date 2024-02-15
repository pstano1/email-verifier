package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fasthttp/router"
	"github.com/pstano1/emailVerifier/internal/api"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

type HTTPInstanceAPI struct {
	log  logrus.FieldLogger
	api  *api.InstanceAPI
	port uint16
}

type HTTPConfig struct {
	Logger      logrus.FieldLogger
	InstanceAPI *api.InstanceAPI
	Port        uint16
}

func NewHTTPInstanceAPI(conf *HTTPConfig) *HTTPInstanceAPI {
	return &HTTPInstanceAPI{
		log:  conf.Logger,
		api:  conf.InstanceAPI,
		port: conf.Port,
	}
}

func (i *HTTPInstanceAPI) Run() {
	r := router.New()

	API := r.Group("/api")
	API.POST("/verify", i.verifyEmailAddresses)

	i.log.Debugf("Starting server at port :%v", i.port)
	log.Fatal(fasthttp.ListenAndServe(fmt.Sprintf(":%v", i.port), r.Handler))
}

func (i *HTTPInstanceAPI) verifyEmailAddresses(ctx *fasthttp.RequestCtx) {
	i.log.Debugf("got request for verifying e-mail addresses")
	var emailAddresses []string
	payload := ctx.Request.Body()
	err := json.Unmarshal(payload, &emailAddresses)
	if err != nil {
		i.log.Errorf("error while unmarshaling payload, error: %v", err)
		ctx.Response.SetBodyString(fmt.Sprintf("error while unmarshaling payload, error: %v", err))
		ctx.Response.SetStatusCode(http.StatusBadRequest)
		return
	}
	verifiedAddresses, err := i.api.VerifyEmailAddress(emailAddresses)
	if err != nil {
		i.log.Errorf("error while veryfing addresses, error: %v", err)
		ctx.Response.SetBodyString(fmt.Sprintf("error while veryfing addresses, error: %v", err))
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		return
	}
	body, err := json.Marshal(verifiedAddresses)
	if err != nil {
		i.log.Errorf("error while marshaling response, error: %v", err)
		ctx.Response.SetBodyString(fmt.Sprintf("error while marshaling response, error: %v", err))
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		return
	}
	ctx.Response.SetBody(body)
	ctx.Response.SetStatusCode(http.StatusOK)
}
