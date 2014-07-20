package gospark

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"net/url"
// 	"strings"
// )

// type SparkCore struct {
// 	client *http.Client
// }

// func NewSparkCore() *SparkCore {
// 	sc := &SparkCore{}
// 	sc.client = &http.Client{}
// 	sc.Oauth = &OAuthService{sparkCore: sc}

// 	return sc
// }

// func (sc *SparkCore) NewRequest(method string, urlStr string,
// 	body string) (*http.Request, error) {
// 	u, err := url.Parse(urlStr)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// buff := &bytes.Buffer{}
// 	// if body != nil {
// 	// 	err := json.NewEncoder(buff).Encode(body)
// 	// 	if err != nil {
// 	// 		return nil, err
// 	// 	}
// 	// }

// 	// fmt.Println(buff)

// 	req, err := http.NewRequest(method, u.String(), strings.NewReader(body))
// 	if err != nil {
// 		return nil, err
// 	}

// 	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
// 	return req, nil
// }

// func (sc *SparkCore) Do(req *http.Request,
// 	i interface{}) (interface{}, error) {
// 	fmt.Println(req)
// 	resp, err := sc.client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	fmt.Println(resp)
// 	defer resp.Body.Close()

// 	if i != nil {
// 		// If a io.Writer interface is supplied. Say for
// 		// writing to a log file.
// 		if w, ok := i.(io.Writer); ok {
// 			io.Copy(w, resp.Body)
// 		} else {
// 			err := json.NewDecoder(resp.Body).Decode(i)
// 			if err != nil {
// 				return nil, err
// 			}
// 		}
// 	}

// 	fmt.Println(i)
// 	return i, nil
// }
