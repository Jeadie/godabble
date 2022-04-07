package godabble

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DabbleApi interface {
	Me() (*Me, error)
	Categories() ([]Category, error)
	CategoryPage() (*CategoryPage, error)
	Crypto(crypto string) (*Ticker, error)
	Home() (*Home, error)
	Comments(rId, rType string, start, limit uint) (*Comments, error)
	NewsPage(slug string) (*NewsPage, error)
	PortfolioPage(slug string) (*PortfolioPage, error)
	Rankings() (*Rankings, error)
	Stock(stockKey string) (*Ticker, error)
	Ticker(slug string) (*Ticker, error)
	Tags() []Tag
	User(username string) (*UserPage, error)
}

type Api struct {
	client  *http.Client
	version uint
}

// Construct an Api to interact with dabble.com.
func Construct() *Api {
	return &Api{
		client:  &http.Client{},
		version: 1,
	}
}

// Chart of a Ticker entity for a given ChartFidelity
func (api *Api) Chart(chartId string, fidelity ChartFidelity) (*Chart, error) {
	c := &Chart{}
	err := api.getDecode(fmt.Sprintf("/charts/ticker?id=%s&fidelity=%s", chartId, fidelity), c)
	if err == nil && len(c.ErrorMessage) > 0 {
		return c, fmt.Errorf("received error in dabble.com response, %s", c.ErrorMessage)
	}
	return c, err
}

// Me endpoint. Sparse fields if unauthenticated.
func (api *Api) Me() (*Me, error) {
	me := &Me{}
	return me, api.getDecode("/me", me)
}

// Home page used by dabble.com
func (api *Api) Home() (*Home, error) {
	h := &Home{}
	return h, api.getDecode("/pages/home", h)
}

// Categories found on the the categories page. Importantly, not the category page itself (which is /pages/categories)
func (api *Api) Categories() ([]Category, error) {
	c := &Categories{}
	err := api.getDecode("/categories", c)
	if err != nil {
		return []Category{}, err
	}
	if len(c.ErrorMessage) > 0 {
		return []Category{}, fmt.Errorf("received error in dabble.com response, %s", c.ErrorMessage)
	}
	return c.Categories, nil
}

// CategoryPage for a single dabble category. Importantly slug is a single `/%s`. Subcategories, i.e. `/%s/%s` Don't
// have the same page structure.
func (api *Api) CategoryPage(slug string) (*CategoryPage, error) {
	n := &CategoryPage{}
	err := api.getDecode(fmt.Sprintf("/pages/category?slug=/category/%s", slug), n)
	if err == nil && len(n.ErrorMessage) > 0 {
		return n, fmt.Errorf("received error in dabble.com response, %s", n.ErrorMessage)
	}
	return n, err
}

// Comments from an arbitrary reference.
func (api *Api) Comments(rId, rType string, start, limit uint) (*Comments, error) {
	c := &Comments{}
	err := api.getDecode(fmt.Sprintf(
		"/comments?reference_id=%s&reference_type=%s&cursor=%d&limit=%d", rId, rType, start, limit,
	), c)
	if err == nil && len(c.ErrorMessage) > 0 {
		return c, fmt.Errorf("received error in dabble.com response, %s", c.ErrorMessage)
	}
	return c, err
}

// Crypto is a page for a cryptocurrency entity. Importantly is not the full slug /crypto/%s, just %s.
func (api *Api) Crypto(crypto string) (*Ticker, error) {
	return api.Ticker(fmt.Sprintf("/crypto/%s", crypto))
}

// NewsPage is a page for a single news article with related content data (i.e. related tags, tickers, etc).
func (api *Api) NewsPage(slug string) (*NewsPage, error) {
	n := &NewsPage{}
	err := api.getDecode(fmt.Sprintf("/pages/news?slug=/news/%s", slug), n)
	if err == nil && len(n.ErrorMessage) > 0 {
		return n, fmt.Errorf("received error in dabble.com response, %s", n.ErrorMessage)
	}
	return n, err
}

// PortfolioPage for a specific portfolio.
func (api *Api) PortfolioPage(slug string) (*PortfolioPage, error) {
	p := &PortfolioPage{}
	err := api.getDecode(fmt.Sprintf("/pages/portfolio?slug=/portfolio/%s", slug), p)
	if err == nil && len(p.ErrorMessage) > 0 {
		return p, fmt.Errorf("received error in dabble.com response, %s", p.ErrorMessage)
	}
	return p, err
}

// Rankings of the largest gaining and losing portfolios
func (api *Api) Rankings() (*Rankings, error) {
	r := &Rankings{}
	return r, api.getDecode("/pages/rankings", r)
}

// Ticker page for a specific stock. stockKey is not a slug, but rather the end stock tag. e.g. AAPL.
func (api *Api) Stock(stockKey string) (*Ticker, error) {
	return api.Ticker(fmt.Sprintf("/stocks/%s", stockKey))
}

// Tags associated with the home page. Tags on a category page in specific are found in CategoryPage.Tags
func (api *Api) Tags() []Tag {
	t := &Tags{}
	err := api.getDecode("/tags", t)
	if err != nil {
		return []Tag{}
	}
	return t.Tags
}

// Ticker is a page for an arbitrary Ticker.
func (api *Api) Ticker(slug string) (*Ticker, error) {
	s := &Ticker{}
	return s, api.getDecode(fmt.Sprintf("/pages/ticker?slug=%s", slug), s)
}

func (api Api) User(username string) (*UserPage, error) {
	p := &UserPage{}
	err := api.getDecode(fmt.Sprintf("/pages/user?slug=/users/%s", username), p)
	if err == nil && len(p.ErrorMessage) > 0 {
		return p, fmt.Errorf("received error in dabble.com response, %s", p.ErrorMessage)
	}
	return p, err
}

func (api *Api) getRaw(uri string) ([]byte, error) {
	res, err := api.client.Get(fmt.Sprintf(
		"https://api.dabble.com/v%d%s", api.version, uri,
	))
	if err != nil {
		return []byte{}, errors.New("GET request bad" + err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, errors.New("badd read all")
		//err
	}
	return body, nil
}

func (api *Api) getDecode(uri string, obj any) error {
	b, err := api.getRaw(uri)
	if err != nil {
		return errors.New("bad getRaw" + err.Error())
	} //err }
	return json.Unmarshal(b, obj)
}
