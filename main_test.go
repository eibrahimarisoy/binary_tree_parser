package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_handleBinaryTree(t *testing.T) {
	t.Run("handleBinaryTree_CaseOne_Success", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handleBinaryTree)
		req.Header.Set("Content-Type", "application/json")
		req.Body = ioutil.NopCloser(bytes.NewBuffer(CaseOnePayload()))
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		resp := Response{}
		err = json.NewDecoder(rr.Body).Decode(&resp)
		if err != nil {
			t.Fatal(err)
		}
		if resp.Sum != 18 {
			t.Errorf("handler returned unexpected result: got %v want %v",
				resp.Sum, 18)
		}
	})

	t.Run("handleBinaryTree_CaseTwo_Success", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handleBinaryTree)
		req.Header.Set("Content-Type", "application/json")
		req.Body = ioutil.NopCloser(bytes.NewBuffer(CaseTwoPayload()))
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		resp := Response{}
		err = json.NewDecoder(rr.Body).Decode(&resp)
		if err != nil {
			t.Fatal(err)
		}
		if resp.Sum != 6 {
			t.Errorf("handler returned unexpected result: got %v want %v",
				resp.Sum, 6)
		}

	})

	t.Run("handleBinaryTree_CaseThree_Fault", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handleBinaryTree)
		req.Header.Set("Content-Type", "application/json")
		req.Body = ioutil.NopCloser(bytes.NewBuffer(CaseThreePayload()))
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

	})
}

func CaseOnePayload() []byte {
	return []byte(
		`{
			"tree": {
			"nodes": [
				{"id": "1", "left": "2", "right": "3", "value": 1},
				{"id": "3",	"left":	"6", "right": "7", "value": 3},
				{"id": "7",	"left":	null, "right": null, "value": 7},
				{"id": "6",	"left":	null, "right": null, "value": 6},
				{"id": "2",	"left":	"4", "right": "5", "value": 2},
				{"id": "5",	"left":	null, "right": null, "value": 5},
				{"id": "4",	"left":	null, "right": null, "value": 4}
			],
			"root": "1"
			}	
		}`)
}

func CaseTwoPayload() []byte {
	return []byte(
		`{
			"tree": {
				"nodes": [
					{"id": "1", "left": "2", "right": "3", "value": 1},
					{"id": "3", "left": null, "right": null, "value": 3},
					{"id": "2", "left": null, "right": null, "value": 2}
					],
				"root": "1"
				}
		}`)
}

func CaseThreePayload() []byte {
	return []byte(
		`{
			"tree": {
				"nodes": [
					{"id": "1", "left": "-10", "right": "-5", "value": 1},
					{"id": "-5", "left": "-20", "right": "-21", "value": -5},
					{"id": "-21", "left": "100-2", "right": "1-3", "value": -21},
					{"id": "1-3", "left": null, "right": null, "value": 1},
					{"id": "100-2", "left": null, "right": null, "value": 100},
					{"id": "-20", "left": "100", "right": "2", "value": -20},
					{"id": "2", "left": null, "right": null, "value": 2},
					{"id": "100", "left": null, "right": null, "value": 100},
					{"id": "-10", "left": "30", "right": "45", "value": -10},
					{"id": "45", "left": "3", "right": "-3", "value": 45},
					{"id": "-3", "left": null, "right": null, "value": -3},
					{"id": "3", "left": null, "right": null, "value": 3},
					{"id": "30", "left": "5", "right": "1-2", "value": 30},
					{"id": "1-2", "left": null, "right": null, "value": 1},
					{"id": "5", "left": null, "right": null, "value": 5}
				],
				"root": "1"
				}
			}
		}`)
}
