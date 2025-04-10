package consumer

import (
	"github.com/nycu-ucr/openapi/nrf/NFManagement"
	"github.com/nycu-ucr/udr/pkg/app"
)

type Consumer struct {
	app.App

	*NrfService
}

func NewConsumer(udr app.App) *Consumer {
	configuration := NFManagement.NewConfiguration()
	configuration.SetBasePath(udr.Context().NrfUri)
	nrfService := &NrfService{
		nfMngmntClients: make(map[string]*NFManagement.APIClient),
	}

	return &Consumer{
		App:        udr,
		NrfService: nrfService,
	}
}
