package gospark

import (
	"testing"
)

var device = NewDevice(DefaultDeviceID)
var aToken, _ = aTokenService.GetAccessToken()

func TestInvokeFunction(t *testing.T) {
	args := []string{"1", "2"}
	device.NewDeviceFunction("test", args)

	_, err := device.InvokeFunction(device.Functions["test"], aToken)
	if err != nil {
		t.Error(err)
	}
}

func TestGetDeviceVariable(t *testing.T) {
	device.NewDeviceVariable("temperature_sensor")
	_, err := device.GetDeviceVariable(
		device.Variables["temperature_sensor"], aToken)
	if err != nil {
		t.Error(err)
	}
}
