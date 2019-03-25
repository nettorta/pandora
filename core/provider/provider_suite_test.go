package provider

import (
	"testing"

	"github.com/nettorta/pandora/core"
	"github.com/nettorta/pandora/lib/ginkgoutil"
)

func TestProvider(t *testing.T) {
	ginkgoutil.RunSuite(t, "AmmoQueue Suite")
}

func testDeps() core.ProviderDeps {
	return core.ProviderDeps{ginkgoutil.NewLogger()}
}
