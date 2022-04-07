package godabble

import (
	"testing"
)

func TestApi(t *testing.T) {
	api := ConstructApi(1)
	h, err := api.Home()
	t.Errorf("Error: %s, h: %#v", err, h)
}
