package pkg

import (
	"context"
	"net/http"
	"net/http/cookiejar"

	"golang.org/x/time/rate"
)

type VSCAppClient struct {
	client      *http.Client
	Ratelimiter *rate.Limiter
}

func (c *VSCAppClient) Do(req *http.Request) (*http.Response, error) {
	ctx := context.Background()
	err := c.Ratelimiter.Wait(ctx)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "vsc-app")
	resp, err := c.client.Do(req)
	c.client.Jar.SetCookies(req.URL, resp.Cookies())
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *VSCAppClient) Jar() *http.CookieJar {
	return &c.client.Jar
}

// NewVSCAppClient return http client with a ratelimiter
func NewVSCAppClient(rl *rate.Limiter) (c *VSCAppClient) {
	jar, _ := cookiejar.New(&cookiejar.Options{})
	c = &VSCAppClient{
		client: &http.Client{
			Jar: jar,
		},
		Ratelimiter: rl,
	}

	return
}
