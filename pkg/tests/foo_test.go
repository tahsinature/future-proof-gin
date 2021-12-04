package tests

import (
	"testing"

	"github.com/tahsinature/future-proof-gin/pkg/application"
)

func TestPingRoute(t *testing.T) {
	application := new(application.Application)
	application.Setup()
}
