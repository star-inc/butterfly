// Butterfly - The web crawler base on Apache Solr for StarStart!
// Copyright(c) 2020 Star Inc. All Rights Reserved.
// The software licensed under Mozilla Public License Version 2.0

package http

import "testing"

const testEndPoint = "https://reqbin.com"

func Test_HttpGet(t *testing.T) {
	client := NewHttpClient(testEndPoint)
	if status, response := client.GET("/echo/get/json"); status == 200 {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_HttpPost(t *testing.T) {
	client := NewHttpClient(testEndPoint)
	if status, response := client.POST("/echo/post/json", nil); status == 200 {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_HttpPut(t *testing.T) {
	client := NewHttpClient(testEndPoint)
	if status, response := client.PUT("/echo/put/json", nil); status == 200 {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_HttpDelete(t *testing.T) {
	client := NewHttpClient(testEndPoint)
	if status, response := client.DELETE("/echo/delete/json", nil); status == 200 {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}

func Test_HttpPatch(t *testing.T) {
	client := NewHttpClient(testEndPoint)
	if status, response := client.PATCH("/echo/patch/json", nil); status == 200 {
		t.Logf("Success: %s", response)
	} else {
		t.Errorf("Failed: [%d] %s", status, response)
	}
}
