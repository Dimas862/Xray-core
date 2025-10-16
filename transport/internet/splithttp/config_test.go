package splithttp_test

import (
	"testing"

	. "github.com/Dimas862/xray-core/transport/internet/splithttp"
)

func Test_GetNormalizedPath(t *testing.T) {
	c := Config{
		Path: "/?world",
	}

	path := c.GetNormalizedPath()
	if path != "/" {
		t.Error("Unexpected: ", path)
	}
}

