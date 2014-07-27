package gospark

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type DeviceService struct {
	Device       *Device
	FunctionList []*DeviceFunction
}

func NewDeviceService(id string) *DeviceService {
	dService := &DeviceService{}
	dService.Device = NewDevice(id)
	return dService
}

func (s *DeviceService) AddFunction(dFunc *DeviceFunction) {
	s.FunctionList = append(s.FunctionList, dFunc)
}

type Device struct {
	ID string
}

func NewDevice(id string) *Device {
	device := &Device{ID: id}
	return device
}

type DeviceFunction struct {
	Name string
	Args []string
}

func NewDeviceFunction(name string, args []string) *DeviceFunction {
	deviceFunction := &DeviceFunction{name, args}
	return deviceFunction
}

type InvokeFunctionResponse struct {
	DeviceID    string `json:"id"`
	Name        string `json:name`
	Connected   bool   `json:connected`
	ReturnValue int32  `json:connected`
}

func (s *DeviceService) InvokeFunction(dFunc *DeviceFunction,
	i interface{}) (*InvokeFunctionResponse, error) {

	endpoint := "/devices/" + s.Device.ID + "/" + dFunc.Name
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
	invokeFumcResp := &InvokeFunctionResponse{}
	err = json.NewDecoder(resp.Body).Decode(invokeFumcResp)
	if err != nil {
		return nil, err
	}
	return invokeFumcResp, nil
}
