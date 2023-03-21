package gitlab

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/sn3d/excav/pkg/log"
)

const GraphQLEndpoint = "/api/graphql"

// Client provide you access to GitLab API v4
// with your identity
type Client struct {
	PersonalAccessToken string
	URL                 string
	HTTPClient          *http.Client
}

func New(host string, token string) *Client {
	client := &Client{
		URL:                 host,
		PersonalAccessToken: token,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
	return client
}

// Login create new client with given identity
func Login(host string, token string) (*Client, error) {
	client := New(host, token)

	// ping the GitLab if we can use token, we expect
	// positive answer
	_, code := client.query("query {currentUser {name}}", "")
	if code != 200 {
		return nil, fmt.Errorf("Login failed, the ping query returns %d", code)
	}

	return client, nil
}

// PlainLogin to GitLab without need to configure it.
// Configuration is loaded from env. variables like
// GITLAB_HOST and GITLAB_TOKEN
//
// This function is useful for testing and experimenting.
func PlainLogin() (*Client, error) {
	host := os.Getenv("GITLAB_HOST")
	token := os.Getenv("GITLAB_TOKEN")
	if host == "" || token == "" {
		return nil, fmt.Errorf("the GITLAB_HOST or GITLAB_TOKEN env. variable is not set")
	}
	return Login(host, token)
}

func (c *Client) Projects() *Projects {
	return &Projects{
		client: c,
	}
}

// send the GraphQL query and return result
func (c *Client) query(q string, vars string) (string, StatusCode) {

	// prepare request with Bearer auth. and
	// wrapped query. The query must be wrapped into
	// json '{ query: "q" }'. For that reason we also need
	// to escape all dohbe quotes '"' in query.
	q = strings.ReplaceAll(q, "\"", "\\\"")
	var body string
	if vars == "" {
		body = "{\"query\": \"" + q + "\"}"
	} else {
		body = "{\"query\": \"" + q + "\", \"variables\":" + vars + " }"
	}

	return c.post(body, GraphQLEndpoint)
}

// GETting data for given endpoint. The endpoint is without base URL.
func (c *Client) get(endpoint string, opts ...interface{}) (string, StatusCode) {
	return c.send("GET", "", endpoint, opts...)
}

// POSTing data to given endpoint. The endpoint is without base URL.
func (c *Client) post(data string, endpoint string, opts ...interface{}) (string, StatusCode) {
	return c.send("POST", data, endpoint, opts...)
}

// DELETE the resource ono given endpoint
func (c *Client) del(endpoint string, opts ...interface{}) (string, StatusCode) {
	return c.send("DELETE", "", endpoint, opts...)
}

func (c *Client) send(method string, data string, endpoint string, opts ...interface{}) (string, StatusCode) {
	url := c.URL + fmt.Sprintf(endpoint, opts...)

	req, _ := http.NewRequest(method, url, strings.NewReader(data))
	req.Header.Add("Authorization", "Bearer "+c.PersonalAccessToken)
	req.Header.Add("Content-Type", "application/json")

	// execute the query and process the result
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Errorw("error in request", err, "method", method, "endpoint", url, "data", data)
		return "", -1
	}

	respBody, _ := ioutil.ReadAll(resp.Body)

	log.Debug("response received", "method", method, "url", url, "code", resp.StatusCode, "body", string(respBody))
	return string(respBody), StatusCode(resp.StatusCode)
}
