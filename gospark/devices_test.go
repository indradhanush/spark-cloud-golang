package gospark

import (
	"testing"
)

// Test variables common among all tests.
var device = NewDevice(DefaultDeviceID)
var aToken, _ = aTokenService.GetAccessToken()

// TestInvokeFunction tests the InvokeFunction method.
func TestInvokeFunction(t *testing.T) {
	args := []string{"1", "2"}
	device.NewDeviceFunction("test", args)

	_, err := device.InvokeFunction(device.Functions["test"], aToken)
	if err != nil {
		t.Error(err)
	}
}

// TestGetDeviceVariable tests the GetDeviceVariable method.
func TestGetDeviceVariable(t *testing.T) {
	device.NewDeviceVariable("temperature_sensor")
	_, err := device.GetDeviceVariable(
		device.Variables["temperature_sensor"], aToken)
	if err != nil {
		t.Error(err)
	}
}
