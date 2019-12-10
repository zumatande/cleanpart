package services

import (
	"net/http"

	"github.com/rs/zerolog"

	"github.com/zumatande/cleanpart/services/additionalinfo"
	"github.com/zumatande/cleanpart/services/book"
	"github.com/zumatande/cleanpart/services/cancel"
	"github.com/zumatande/cleanpart/services/search"
	"github.com/zumatande/cleanpart/services/status"
)

// Provider defines umbrella interface for services provided by partner.
// The naming of this interface is unimportant. It could be renamed to API,
// ServiceProvider, Partner, etc.
type Provider interface {
	search.Service
	additionalinfo.Service
	book.Service
	status.Service
	cancel.Service
}

// API implements services Provider interface.
type API struct {
	Client *http.Client
	Logger *zerolog.Logger
}
