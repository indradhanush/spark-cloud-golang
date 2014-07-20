package gospark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type SparkCore struct {
	client *http.Client

	Oauth *OAuthService
}

func NewSparkCore() *SparkCore {
	sc := &SparkCore{}
	sc.client = &http.Client{}
	sc.Oauth = &OAuthService{sparkCore: sc}

	return sc
}

func (sc *SparkCore) NewRequest(method string, urlStr string,
	body interface{}) (*http.Request, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	buff := &bytes.Buffer{}
	if body != nil {
		err := json.NewEncoder(buff).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buff)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (sc *SparkCore) Do(req *http.Request,
	i interface{}) (interface{}, error) {
	resp, err := sc.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if i != nil {
		// If a io.Writer interface is supplied. Say for
		// writing to a log file.
		if w, ok := i.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err := json.NewDecoder(resp.Body).Decode(i)
			if err != nil {
				return nil, err
			}
		}
	}

	fmt.Println(i)
	return i, nil
}
