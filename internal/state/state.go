package state

import (
	"github.com/wolv89/bootgator/internal/config"
	"github.com/wolv89/bootgator/internal/database"
)

type State struct {
	Config *config.Config
	DB     *database.Queries
}
