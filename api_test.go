package godabble

import (
	"testing"
)

// TestApi calls each endpoint in Api, and confirms JSON parsing is valid and no other errors occur.
func TestApi(t *testing.T) {
	api := Construct()

	_, err := api.Me()
	if err != nil {
		t.Errorf("failed to call Me, %s", err.Error())
	}

	_, err = api.Categories()
	if err != nil {
		t.Errorf("failed to call Categories, %s", err.Error())
	}

	_, err = api.CategoryPage("technology")
	if err != nil {
		t.Errorf("failed to call CategoryPage, %s", err.Error())
	}

	_, err = api.Chart("c4s114u73jrtd86rhrd0", fiveDays)
	if err != nil {
		t.Errorf("failed to call Chart, %s", err.Error())
	}

	_, err = api.Comments("c8la7im73jrvpjmpo42g", "portfolio", 0, 5)
	if err != nil {
		t.Errorf("failed to call Comments, %s", err.Error())
	}

	_, err = api.NewsPage("elon-musk-says-domino-effect-led-to-twitter-button-from-him-selling-first-firm-in-the-90s-c97bc8e73jrtmrh0sgfg")
	if err != nil {
		t.Errorf("failed to call NewsPage, %s", err.Error())
	}

	_, err = api.PortfolioPage("bitcoin-only-100-bitcoin-c8fuq8m73jru622m1f50")
	if err != nil {
		t.Errorf("failed to call PortfolioPage, %s", err.Error())
	}

	_, err = api.Rankings()
	if err != nil {
		t.Errorf("failed to call Rankings, %s", err.Error())
	}

	_, err = api.Stock("AAPL")
	if err != nil {
		t.Errorf("failed to call Stock, %s", err.Error())
	}

	tags := api.Tags()
	if len(tags) == 0 {
		t.Errorf("no tags returned when calling /tags")
	}

	_, err = api.Home()
	if err != nil {
		t.Errorf("failed to call Home, %s", err.Error())
	}

	_, err = api.User("luc")
	if err != nil {
		t.Errorf("failed to call User, %s", err.Error())
	}

	_, err = api.Ticker("/crypto/x-adausd")
	if err != nil {
		t.Errorf("failed to call Ticker, %s", err.Error())
	}

	_, err = api.Crypto("x-adausd")
	if err != nil {
		t.Errorf("failed to call Ticker, %s", err.Error())
	}
}
