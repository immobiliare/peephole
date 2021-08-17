package saltapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type LoginResponse struct {
	Return []struct {
		User   string  `json:"user"`
		Eauth  string  `json:"eauth"`
		Token  string  `json:"token"`
		Start  float64 `json:"start"`
		Expire float64 `json:"expire"`
		// Perms  []string `json:"perms"`
	} `json:"return"`
}

func init() {
	if os.Getenv("DEBUG") != "" {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func Login(endpoint, user, pass, client string) (*LoginResponse, error) {
	logrus.WithFields(logrus.Fields{
		"endpoint": endpoint,
		"user":     user,
		"pass":     pass,
		"client":   client,
	}).Debugln("Logging in")
	reqParams := url.Values{}
	reqParams.Add("eauth", client)
	reqParams.Add("username", user)
	reqParams.Add("password", pass)
	reqBody := strings.NewReader(reqParams.Encode())
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/login", endpoint), reqBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	} else if resp.StatusCode != 200 {
		return nil, fmt.Errorf("login endpoint at %s returned %d", endpoint, resp.StatusCode)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	obj := new(LoginResponse)
	if err := json.Unmarshal(body, obj); err != nil {
		return nil, err
	}

	return obj, nil
}
