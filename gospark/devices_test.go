package gospark

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var dService = NewDeviceService(DefaultDeviceID)
var aToken *OAuthResponse

func TestAddFunction(t *testing.T) {
	args := []string{"1", "2"}
	dFunc := NewDeviceFunction("test", args)
	dService.AddFunction(dFunc)
	assert.Equal(t, dFunc, dService.FunctionList[0])
}

func TestInvokeFunction(t *testing.T) {
	aToken, _ = aTokenService.GetAccessToken()
	_, err := dService.InvokeFunction(dService.FunctionList[0], aToken)
	if err != nil {
		t.Error(err)
	}
}
