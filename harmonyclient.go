package harmonyclient

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/univedo/api2go/jsonapi"
)

func init() {

}

// Client the main Harmony Client
type Client struct {
	// The client configuration
	Config Config
	_data  string
}

// NewHarmonyClient returns a new Harmony Client instance
func NewHarmonyClient(conf Config) (c *Client, err error) {
	c = &Client{
		Config: conf,
		_data:  "chicken",
	}

	err = nil
	return
}

// Containers gets a list of all the containers
func (C *Client) Containers() (*[]Container, error) {
	m := map[string]string{}

	payload := new(map[string]interface{})
	if err := C.get("/containers", m, payload); err != nil {
		return nil, err
	}

	if err := (*payload)["errors"]; err != nil {
		err := err.([]interface{})
		e := err[0].(map[string]interface{})

		return nil, fmt.Errorf("[%d] %s", int(e["status"].(float64)), e["title"])
	}

	// fmt.Printf("\n\nHERE: %+v\n\n\n", payload)
	var containers []Container
	err := jsonapi.Unmarshal(*payload, &containers)

	return &containers, err
}

func (C *Client) get(url string, params map[string]string, payload interface{}) error {

	// build and execute the request the resource from the api server
	buf, err := C.request("GET", url, params, nil)

	// ensure that we didnt fail making the request
	if err != nil {
		return fmt.Errorf("Failed requesting %s: %s", url, err)
	}

	// decode the json, checking for errors if they exist
	err = json.Unmarshal(buf, payload)
	if err != nil {
		return fmt.Errorf("json unmarshal error: %s", err)
	}

	return nil
}

func (C *Client) request(requestType, urlSuffix string, params map[string]string, body io.Reader) ([]byte, error) {

	// set the http transport config
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: (!C.Config.APIVerifySSL)},
	}

	// instianiate a new http client
	client := &http.Client{
		// CheckRedirect: redirectPolicyFunc,
		Transport: tr,
	}

	// build the query string
	queryBuffer := bytes.NewBufferString("?")
	for k, v := range params {
		queryBuffer.WriteString(fmt.Sprintf("%s=%s&", k, v))
	}

	// remove the trailing amp if it exists
	queryString := strings.TrimRight(queryBuffer.String(), "&")

	if queryString == "?" {
		queryString = ""
	}

	// initialize a new HTTP request with the given requestType (ie GET|POST|PATCH|etc)
	url := fmt.Sprintf("%s/%s%s%s", C.Config.APIHost, C.Config.APIVersion, urlSuffix, queryString)
	req, err := http.NewRequest(requestType, url, body)

	// add required http headers to the request
	req.Header.Add("Content-Type", "application/json")

	// execute the request
	resp, err := client.Do(req)

	// check for errors (go figure)
	if err != nil {
		return nil, fmt.Errorf("Failed executing request: %s", err)
	}

	// close the response body buffer of the request after method return
	defer resp.Body.Close()

	// read and return the response body buffer
	return ioutil.ReadAll(resp.Body)
}
