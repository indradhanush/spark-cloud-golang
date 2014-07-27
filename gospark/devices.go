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
	Variables map[string]*DeviceVariable
}

func NewDevice(id string) *Device {
	device := &Device{}
	device.ID = id
	device.Functions = make(map[string]*DeviceFunction)
	device.Variables = make(map[string]*DeviceVariable)
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

	token, err := ParseToken(i)
	if err != nil {
		return nil, err
	}
	form.Set("access_token", token)

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

func (s *Device) NewDeviceVariable(name string) {
	dVar := &DeviceVariable{}

	// Truncating the string upto 12 chars.
	truncName := name
	if len(name) > MaxVariableLen {
		truncName = name[:MaxVariableLen]
	}
	dVar.Name = truncName
	s.Variables[name] = dVar
}

type CoreInfo struct {
	LastApp   string `json:"last_app, omitempty"`
	LastHeard string `json:"last_heard, omitempty"`
	Connected bool   `json:"connected, omitempty"`
	DeviceID  string `json:"deviceID, omitempty"`
}

type GetDeviceVariableResponse struct {
	Cmd  string `json:"cmd, omitempty"`
	Name string `json:"name, omitempty"`
	// Result can be a string, int, float or bool. Not the best
	// way to do this here, but I'm sorry about this. Ugly I will agree.
	Result   interface{} `json:"result, omitempty"`
	CoreInfo CoreInfo    `json:"coreInfo, omitempty"`
}

func (d *Device) GetDeviceVariable(dVar *DeviceVariable, i interface{}) (
	*GetDeviceVariableResponse, error) {

	endpoint := "/devices/" + d.ID + "/" + dVar.Name

	urlStr := GetCompleteEndpointUrl(&APIUrl{BaseUrl, APIVersion, endpoint})

	form := url.Values{}

	token, err := ParseToken(i)
	if err != nil {
		return nil, err
	}
	form.Set("access_token", token)
	req, err := http.NewRequest("GET", urlStr, nil)
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
		GetDeviceVariableResponse
		ErrorResponse
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	if response.Err != "" {
		return nil, response.ErrorResponse
	}
	return &response.GetDeviceVariableResponse, nil
}
