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

	"github.com/manyminds/api2go/jsonapi"
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

// Container gets a container by ID
func (C *Client) Container(ID string) (*Container, error) {
	m := map[string]string{}

	response := new(map[string]interface{})
	if err := C.get(fmt.Sprintf("/containers/%s", ID), m, response); err != nil {
		return nil, err
	}

	if err := (*response)["errors"]; err != nil {
		err := err.([]interface{})
		e := err[0].(map[string]interface{})

		return nil, fmt.Errorf("[%d] %s", int(e["status"].(float64)), e["title"])
	}

	var container Container
	err := jsonapi.Unmarshal(*response, &container)

	return &container, err
}

// Containers gets a list of all the containers
func (C *Client) Containers() (*[]Container, error) {
	m := map[string]string{}

	response := new(map[string]interface{})
	if err := C.get("/containers", m, response); err != nil {
		return nil, err
	}

	if err := (*response)["errors"]; err != nil {
		err := err.([]interface{})
		e := err[0].(map[string]interface{})

		return nil, fmt.Errorf("[%d] %s", int(e["status"].(float64)), e["title"])
	}

	// fmt.Printf("\n\nHERE: %+v\n\n\n", response)
	var containers []Container
	err := jsonapi.Unmarshal(*response, &containers)

	return &containers, err
}

// Machines gets a list of all the machines
func (C *Client) Machines() (*[]Machine, error) {
	m := map[string]string{}

	response := new(map[string]interface{})
	if err := C.get("/machines", m, response); err != nil {
		return nil, err
	}

	if err := (*response)["errors"]; err != nil {
		err := err.([]interface{})
		e := err[0].(map[string]interface{})

		return nil, fmt.Errorf("[%d] %s", int(e["status"].(float64)), e["title"])
	}

	// fmt.Printf("\n\nHERE: %+v\n\n\n", response)
	var machines []Machine
	err := jsonapi.Unmarshal(*response, &machines)

	return &machines, err
}

// ContainersAdd will create a new Container resource
func (C *Client) ContainersAdd(c *Container) (*Container, error) {
	// marshal the resource
	payload, err := jsonapi.MarshalToJSON(c)
	if err != nil {
		return nil, err
	}

	// execute the request
	m := map[string]string{}
	response := new(map[string]interface{})
	if err := C.post("/containers", m, payload, response); err != nil {
		return nil, err
	}

	// handle api errors
	if err := (*response)["errors"]; err != nil {
		err := err.([]interface{})
		e := err[0].(map[string]interface{})

		return nil, fmt.Errorf("[%d] %s", int(e["status"].(float64)), e["title"])
	}

	// unmarshal the response
	var container Container
	err = jsonapi.Unmarshal(*response, &container)

	return &container, err
}

// MachineByName will fetch a machine by its name
func (C *Client) MachineByName(name string) (*Machine, error) {
	// setup the filters
	m := map[string]string{}
	m["name"] = name

	// execute the request
	response := new(map[string]interface{})
	if err := C.get("/machines", m, response); err != nil {
		return nil, err
	}

	// handle api errors
	if err := (*response)["errors"]; err != nil {
		err := err.([]interface{})
		e := err[0].(map[string]interface{})

		return nil, fmt.Errorf("[%d] %s", int(e["status"].(float64)), e["title"])
	}

	// unmarshal the response
	var machine Machine
	err := jsonapi.Unmarshal(*response, &machine)

	return &machine, err
}

// Machine gets a machine by ID
func (C *Client) Machine(ID string) (*Machine, error) {
	m := map[string]string{}

	response := new(map[string]interface{})
	if err := C.get(fmt.Sprintf("/machines/%s", ID), m, response); err != nil {
		return nil, err
	}

	if err := (*response)["errors"]; err != nil {
		err := err.([]interface{})
		e := err[0].(map[string]interface{})

		return nil, fmt.Errorf("[%d] %s", int(e["status"].(float64)), e["title"])
	}

	var machine Machine
	err := jsonapi.Unmarshal(*response, &machine)

	return &machine, err
}

// MachinesAdd will create a Harmony Machine resource
func (C *Client) MachinesAdd(c *Machine) (*Machine, error) {
	// marshal the resource
	payload, err := jsonapi.MarshalToJSON(c)
	if err != nil {
		return nil, err
	}

	// execute the request
	m := map[string]string{}
	response := new(map[string]interface{})
	if err := C.post("/machines", m, payload, response); err != nil {
		return nil, err
	}

	// handle api errors
	if err := (*response)["errors"]; err != nil {
		err := err.([]interface{})
		e := err[0].(map[string]interface{})

		return nil, fmt.Errorf("[%d] %s", int(e["status"].(float64)), e["title"])
	}

	// unmarshal the response
	var machine Machine
	err = jsonapi.Unmarshal(*response, &machine)
	return &machine, err
}

func (C *Client) get(url string, params map[string]string, response interface{}) error {
	// build and execute the request the resource from the api server
	buf, err := C.request("GET", url, params, nil)

	// ensure that we didnt fail making the request
	if err != nil {
		return fmt.Errorf("Failed requesting %s: %s", url, err)
	}

	// decode the json, checking for errors if they exist
	err = json.Unmarshal(buf, response)
	if err != nil {
		return fmt.Errorf("json unmarshal error: %s", err)
	}

	return nil
}

// put will execute a put request
func (C *Client) put(url string, params map[string]string, payload []byte, response interface{}) error {
	body := bytes.NewReader(payload)
	buf, err := C.request("PUT", url, params, body)
	return C.handleResponse(url, buf, err, response)
}

// post will execute a post request
func (C *Client) post(url string, params map[string]string, payload []byte, response interface{}) error {
	body := bytes.NewReader(payload)
	buf, err := C.request("POST", url, params, body)
	return C.handleResponse(url, buf, err, response)
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

func (C *Client) handleResponse(url string, payload []byte, err error, unmarshalInto interface{}) error {
	// ensure that we didnt fail making the request
	if err != nil {
		return fmt.Errorf("Failed requesting %s: %s", url, err)
	}

	// fmt.Printf("%s\n\n", string(payload))

	// decode the json, checking for errors if they exist
	err = json.Unmarshal(payload, unmarshalInto)
	if err != nil {
		return fmt.Errorf("json unmarshal error: %s", err)
	}

	return nil
}
