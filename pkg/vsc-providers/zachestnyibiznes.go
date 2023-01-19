package providers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/EnergoStalin/vsc-app/pkg"
	"github.com/EnergoStalin/vsc-app/pkg/vsc-providers/responces"
	log "github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
)

type ZaChestnyiBiznesProvider struct {
	client *pkg.VSCAppClient
	logger *log.Logger
	token  string
}

func (z *ZaChestnyiBiznesProvider) appendBase(path string) string {
	return z.GetUrl() + path
}

func (z *ZaChestnyiBiznesProvider) insertToken(req *http.Request) *http.Request {
	req.Header.Add("X-Csrf-Token", z.token)
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Referer", z.GetUrl())
	return req
}

func (z *ZaChestnyiBiznesProvider) authenticate(ctx context.Context) (err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, z.appendBase("/site/ajax-login"), nil)
	req.Header.Add("X-Csrf-Token", "random_bytes")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	res, err := z.client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	dc := json.NewDecoder(res.Body)

	var token responces.LoginResponce

	err = dc.Decode(&token)
	if err != nil {
		return
	}

	z.token = token.Token

	return
}

func (z *ZaChestnyiBiznesProvider) autocomplete(term string, ctx context.Context) (r *responces.AutocompleteResponce, err error) {
	if z.token == "" {
		err = z.authenticate(ctx)
		if err != nil {
			return
		}
	}

	url, _ := url.Parse(z.appendBase("/site/get-autocomplete-api"))
	q := url.Query()
	q.Add("query", term)
	q.Add("index", "ul")
	url.RawQuery = q.Encode()
	u := url.String()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, u, nil)
	if err != nil {
		return
	}

	res, err := z.client.Do(z.insertToken(req))
	if err != nil {
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	r = new(responces.AutocompleteResponce)
	err = json.Unmarshal(body, r)
	return
}

func (z *ZaChestnyiBiznesProvider) Search(term string, ctx context.Context) (result responces.SearchResponce, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	data, err := z.autocomplete(term, ctx)
	if err != nil {
		return
	}

	fns := data.Docs[0].Source.FNS

	return responces.SearchResponce{
		Title:   fns.URLONG,
		County:  "na",
		Address: "na",
		Site:    "na",
		OKVED:   []string{},
		INN:     fns.INN,
		OGRN:    fns.OGRN,
	}, nil
}

func (z *ZaChestnyiBiznesProvider) GetName() string {
	return "ZaChestnyiBiznes"
}

func (z *ZaChestnyiBiznesProvider) GetUrl() string {
	return "https://zachestnyibiznes.ru"
}

func NewZaChestnyiBiznesProvider() (z *ZaChestnyiBiznesProvider) {
	z = &ZaChestnyiBiznesProvider{
		client: pkg.NewVSCAppClient(rate.NewLimiter(rate.Every(50*time.Millisecond), 1)),
		logger: log.New(),
	}

	z.logger.SetOutput(os.Stdout)
	z.logger.SetLevel(log.TraceLevel)
	z.logger.SetFormatter(&log.JSONFormatter{})

	return
}
