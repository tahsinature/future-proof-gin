package tests

import (
	"testing"

	"github.com/tahsinature/future-proof-gin/pkg/application"
	"github.com/tahsinature/future-proof-gin/pkg/routes"
)

func TestPingRoute(t *testing.T) {
	new(application.Application).Setup()
	routes.Setup()

	// router.ServeHTTP()
}
