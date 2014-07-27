package gospark

import (
	"testing"
)

var device = NewDevice(DefaultDeviceID)
var aToken *OAuthResponse

func TestInvokeFunction(t *testing.T) {
	aToken, _ = aTokenService.GetAccessToken()
	args := []string{"1", "2"}
	device.NewDeviceFunction("test", args)

	_, err := device.InvokeFunction(device.Functions["test"], aToken)
	if err != nil {
		t.Error(err)
	}
}
