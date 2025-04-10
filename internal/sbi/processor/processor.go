package processor

import (
	"github.com/nycu-ucr/udr/internal/database"
	"github.com/nycu-ucr/udr/pkg/app"
)

type Processor struct {
	app.App
	database.DbConnector
}

func NewProcessor(udr app.App) *Processor {
	return &Processor{
		App:         udr,
		DbConnector: database.NewDbConnector(udr.Config().Configuration.DbConnectorType),
	}
}
