package Infrastructure

import (
	"github.com/elastic/go-elasticsearch/v8"
	"User"
	"Publication"
	"Test"
)

type WebServiceHandler struct {
	UserRepository User.UserHandler
	PublicationRepository Publication.PublicationHandler
	ESC elasticsearch.Client
	Test Test.TestHandler
	OAuthServerAddress string
}