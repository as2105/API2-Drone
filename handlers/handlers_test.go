// Copyright © 2018 Optum
package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestRouting(t *testing.T) {
	srv := httptest.NewServer(newCapabilityStatementHandler(nil))
	defer srv.Close()

	res, err := http.Get(fmt.Sprintf("%s/double?v=2", srv.URL))
	if err != nil {
		t.Fatalf("could not send GET request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	d, err := strconv.Atoi(string(bytes.TrimSpace(b)))
	if err != nil {
		t.Fatalf("expected an integer; got %s", b)
	}
	if d != 4 {
		t.Fatalf("expected double to be 4; got %v", d)
	}
}
