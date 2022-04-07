package godabble

import (
	"testing"
)

func TestApi(t *testing.T) {
	api := ConstructApi(1)
	_, err := api.Home()
	if err != nil {
		t.Errorf(err.Error())
	}
}
