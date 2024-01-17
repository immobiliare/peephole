package salt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"io"

	_util "github.com/immobiliare/peephole/util"
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
	if _util.Debugging() {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func Login(endpoint, user, pass, client string) (*LoginResponse, error) {
	logrus.WithFields(logrus.Fields{
		"endpoint": endpoint,
		"user":     user,
		"pass":     strings.Repeat("*", len(pass)),
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

	body, err := io.ReadAll(io.Reader(resp.Body))
	if err != nil {
		return nil, err
	}

	obj := new(LoginResponse)
	if err := json.Unmarshal(body, obj); err != nil {
		return nil, err
	}

	return obj, nil
}
