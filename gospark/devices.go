package gospark

import (
	"encoding/json"
	"fmt"
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

// Device is a representation of a Spark Core.
type Device struct {
	ID        string
	Functions map[string]*DeviceFunction
	Variables map[string]*DeviceVariable
}

// NewDevice is a constructor for Device.
func NewDevice(id string) *Device {
	device := &Device{}
	device.ID = id
	device.Functions = make(map[string]*DeviceFunction)
	device.Variables = make(map[string]*DeviceVariable)
	return device
}

// DeviceFunction is a representation of a function.
type DeviceFunction struct {
	Name string
	Args []string
}

// NewDeviceFunction is a constuctor for DeviceFunction that also
// links the function to a Device object.
func (s *Device) NewDeviceFunction(name string, args []string) {
	dFunc := &DeviceFunction{name, args}
	s.Functions[name] = dFunc
}

// InvokeFunctionResponse is the representation of a response when a
// particular function is invoked via the REST API.
type InvokeFunctionResponse struct {
	DeviceID    string `json:"id"`
	Name        string `json:name`
	Connected   bool   `json:connected`
	ReturnValue int32  `json:connected`
}

// InvokeFunction is used to Invoke a DeviceFunction instance.
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

// DeviceVariable is a representation of a Variable that can be
// associated with a device.
type DeviceVariable struct {
	Name string
}

// NewDeviceVariable is a constructor for DeviceVariable and links it
// to a Device object.
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

// CoreInfo is a representation of the json struct coreInfo.
type CoreInfo struct {
	LastApp   string `json:"last_app, omitempty"`
	LastHeard string `json:"last_heard, omitempty"`
	Connected bool   `json:"connected, omitempty"`
	DeviceID  string `json:"deviceID, omitempty"`
}

// GetDeviceVariableResponse is a representation of the response
// received on doing a GET on a Device Variable.
type GetDeviceVariableResponse struct {
	Cmd  string `json:"cmd, omitempty"`
	Name string `json:"name, omitempty"`
	// Result can be a string, int, float or bool. Not the best
	// way to do this here, but I'm sorry about this. Ugly I will agree.
	Result   interface{} `json:"result, omitempty"`
	CoreInfo CoreInfo    `json:"coreInfo, omitempty"`
}

// GetDeviceVariable is the method to GET the value of a particular
// Device Variable.
func (d *Device) GetDeviceVariable(dVar *DeviceVariable, i interface{}) (
	*GetDeviceVariableResponse, error) {

	endpoint := "/devices/" + d.ID + "/" + dVar.Name

	urlStr := GetCompleteEndpointUrl(&APIUrl{BaseUrl, APIVersion, endpoint})

	token, err := ParseToken(i)
	if err != nil {
		return nil, err
	}

	urlStr += fmt.Sprintf("?access_token=%v", token)

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
