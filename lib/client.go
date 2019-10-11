package lib

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
)

var defaultURLStr = "https://api.zenhub.io/"

// Client is zenhub Client
type Client struct {
	URL         *url.URL
	HTTPClient  *http.Client
	AccessToken string
	Logger      *log.Logger
}

func (c *Client) newRequest(ctx context.Context, method, spath string, body io.Reader) (*http.Request, error) {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)
	req, err := http.NewRequest(method, u.String(), body)

	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	req.Header.Set("X-Authentication-Token", c.AccessToken)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (c *Client) GetIssue(ctx context.Context, repoID, issueNumber int64) {

}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

// NewClient is return Client
func NewClient(inputURL, accessToken string, logger *log.Logger) (*Client, error) {
	urlStr := inputURL

	if len(accessToken) == 0 {
		return nil, errors.New("missing access token")
	}

	if len(urlStr) == 0 {
		urlStr = defaultURLStr
	}

	parseURL, err := url.ParseRequestURI(urlStr)

	if err != nil {
		return nil, err
	}

	var discardLogger = log.New(ioutil.Discard, "", log.LstdFlags)
	if logger == nil {
		logger = discardLogger
	}

	client := Client{
		URL:         parseURL,
		HTTPClient:  &http.Client{},
		AccessToken: accessToken,
		Logger:      logger,
	}

	return &client, nil
}
