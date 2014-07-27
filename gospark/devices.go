package gospark

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

// type DeviceService struct {
// 	DeviceList []*Device
// }

// func NewDeviceService() *DeviceService {
// 	dService := &DeviceService{}
// 	return dService
// }

type Device struct {
	ID        string
	Functions map[string]*DeviceFunction
}

func NewDevice(id string) *Device {
	device := &Device{}
	device.ID = id
	device.Functions = make(map[string]*DeviceFunction)
	return device
}

type DeviceFunction struct {
	Name string
	Args []string
}

func (s *Device) NewDeviceFunction(name string, args []string) {
	dFunc := &DeviceFunction{name, args}
	s.Functions[name] = dFunc
}

type InvokeFunctionResponse struct {
	DeviceID    string `json:"id"`
	Name        string `json:name`
	Connected   bool   `json:connected`
	ReturnValue int32  `json:connected`
}

func (s *Device) InvokeFunction(dFunc *DeviceFunction,
	i interface{}) (*InvokeFunctionResponse, error) {

	endpoint := "/devices/" + s.ID + "/" + dFunc.Name
	urlStr := GetCompleteEndpointUrl(&APIUrl{BaseUrl, APIVersion,
		endpoint})

	form := url.Values{}

	if t, ok := i.(*OAuthResponse); ok {
		form.Set("access_token", t.AccessToken)
	} else if t, ok := i.(*AccessToken); ok {
		form.Set("access_token", t.Token)
	} else {
		return nil, &ApiError{"Pass either OAuthResponse, or AccessToken"}
	}

	argsList := ""
	for i := 0; i < len(dFunc.Args)-1; i++ {
		argsList += dFunc.Args[i] + ","
	}
	argsList += dFunc.Args[len(dFunc.Args)-1]
	form.Set("args", argsList)

	req, err := http.NewRequest("POST", urlStr,
		strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var response struct {
		InvokeFunctionResponse
		ErrorResponse
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	if response.Err != "" {
		return nil, response.ErrorResponse
	}
	return &response.InvokeFunctionResponse, nil
}

type DeviceVariable struct {
	Name string
}

func NewDeviceVariable(name string) *DeviceVariable {
	deviceVariable := &DeviceVariable{}

	// Truncating the string upto 12 chars.
	if len(name) > 12 {
		name = name[:12]
	}
	deviceVariable.Name = name
	return deviceVariable
}
