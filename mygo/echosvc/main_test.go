package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)


func TestEcho(t *testing.T) {
	res := httptest.NewRecorder()
	msg := "Hello,World!"
	req, _ := http.NewRequest("POST", "/", strings.NewReader(msg))
	echo(res,req)
	if res.Code != 200 {
		t.Errorf("Unexpected error code. Expected : 200, Actual %d",res.Code)
	}

	contentType := res.HeaderMap.Get("Content-Type")
	const (
		contentTypePlain = "text/plain; charset=utf-8"
	)
	t.Logf("Content type %s",contentType)
	if  contentType != contentTypePlain  {
		t.Errorf("Unexptected content type. Expected %s, Actual %s", contentTypePlain, contentType)}

	if res.Body.String() != msg {
		t.Errorf("Unexpected message body. Expected : %s, Actual %s", msg, res.Body.String())
	}
	}
	
		

