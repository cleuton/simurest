package tests

import (
	"errors"
	"strings"
	"testing"

	"network.golang/simurest/cmd"
)

var (
	deleteInvoked bool
	deletedKey    string
)

/*
[--quiet | -q <do not print this message (default true)>]
[--method | -m <http method (default=GET): GET | POST | HEAD | DELETE | PUT>]
[--port | -p <tcp port (default=8080)> ]
[--uri | -u <http URI (default="/")]
[--status | -s <http Status of the response (default=200)]
[--body | -b <http response body (default: '{"status": "ok"}')]
*/
func TestSetflagsOk(t *testing.T) {
	/*
	   // Given
	   os.Args = []string{"cmd"}

	   // When
	   SetFlags()

	   // Then

	   	if FlagQuiet {
	   		t.Error("Flagquiet should be false")
	   	}

	   	if FlagMethod != "GET" {
	   		t.Error("FlatMethod should be GET")
	   	}

	   	if FlatUri != "/" {
	   		t.Error("FlagUri should be '/'")
	   	}

	   	if FlagStatus != 200 {
	   		t.Error("FlagStatus should be 200")
	   	}

	   	if FlagBody != `{"status": "ok"}` {
	   		t.Error("FlagBody should be: {\"status\": \"ok\"}")
	   	}
	*/
}
