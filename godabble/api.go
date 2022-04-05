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
	Home() (*Home, error)
	Comments(rId, rType string, start, limit uint) (*Comments, error)
	NewsPage(slug string) (*NewsPage, error)
	PortfolioPage(slug string) (*PortfolioPage, error)
	Rankings() (*Rankings, error)
	Stock(stockKey string) (*Stock, error)
	Tags() []Tag
}

type Api struct {
	client  *http.Client
	version uint
}

// Construct an api object to interact with dabble.com. Version number should be at least 1.
func ConstructApi(v uint) *Api {
	return &Api{
		client:  &http.Client{},
		version: v,
	}
}

func (api *Api) GetRaw(uri string) ([]byte, error) {
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

func (api *Api) GetDecode(uri string, obj any) error {
	b, err := api.GetRaw(uri)
	if err != nil {
		return errors.New("bad GetRaw" + err.Error())
	} //err }
	return json.Unmarshal(b, obj)
}

func (api *Api) Me() (*Me, error) {
	me := &Me{}
	return me, api.GetDecode("/me", me)
}
func (api *Api) Home() (*Home, error) {
	h := &Home{}
	return h, api.GetDecode("/pages/home", h)
}

func (api *Api) Categories() ([]Category, error) {
	c := &Categories{}
	err := api.GetDecode("/categories", c)
	if err != nil {
		return []Category{}, err
	}
	if len(c.ErrorMessage) > 0 {
		return []Category{}, fmt.Errorf("received error in dabble.com response, %s", c.ErrorMessage)
	}
	return c.Categories, nil
}

func (api *Api) CategoryPage(slug string) (*CategoryPage, error) {
	n := &CategoryPage{}
	err := api.GetDecode(fmt.Sprintf("/pages/category?slug=/category/%s", slug), n)
	if err == nil && len(n.ErrorMessage) > 0 {
		return n, fmt.Errorf("received error in dabble.com response, %s", n.ErrorMessage)
	}
	return n, err
}

func (api *Api) Comments(rId, rType string, start, limit uint) (*Comments, error) {
	c := &Comments{}
	err := api.GetDecode(fmt.Sprintf(
		"/comments?reference_id=%s&reference_type=%s&cursor=%d&limit=%d", rId, rType, start, limit,
	), c)
	if err == nil && len(c.ErrorMessage) > 0 {
		return c, fmt.Errorf("received error in dabble.com response, %s", c.ErrorMessage)
	}
	return c, err
}

func (api *Api) NewsPage(slug string) (*NewsPage, error) {
	n := &NewsPage{}
	err := api.GetDecode(fmt.Sprintf("/pages/news?slug=/news/%s", slug), n)
	if err == nil && len(n.ErrorMessage) > 0 {
		return n, fmt.Errorf("received error in dabble.com response, %s", n.ErrorMessage)
	}
	return n, err
}

func (api *Api) PortfolioPage(slug string) (*PortfolioPage, error) {
	p := &PortfolioPage{}
	err := api.GetDecode(fmt.Sprintf("/pages/portfolio?slug=/portfolio/%s", slug), p)
	if err == nil && len(p.ErrorMessage) > 0 {
		return p, fmt.Errorf("received error in dabble.com response, %s", p.ErrorMessage)
	}
	return p, err
}

func (api *Api) Rankings() (*Rankings, error) {
	r := &Rankings{}
	return r, api.GetDecode("/pages/rankings", r)
}

func (api *Api) Stock(stockKey string) (*Stock, error) {
	s := &Stock{}
	return s, api.GetDecode(fmt.Sprintf("/pages/ticker?slug=/stocks/%s", stockKey), s)
}

func (api *Api) Tags() []Tag {
	t := &Tags{}
	err := api.GetDecode("/tags", t)
	if err != nil {
		return []Tag{}
	}
	return t.Tags
}
