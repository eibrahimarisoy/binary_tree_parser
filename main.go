package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Arr []interface{}

type Result struct {
	Sum   int    `json:"sum"`
	Error string `json:"error"`
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var p = map[string]interface{}{}

		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			log.Println(err)
			return
		}
		var arr = Arr{}
		arr = p["tree"].(map[string]interface{})["nodes"].([]interface{})
		rootData := p["tree"].(map[string]interface{})["root"].(string)

		rootNum, err := strconv.Atoi(rootData)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}

		sum, err := arr.maxPathSum(arr.Find(rootNum))
		result := Result{Sum: 0, Error: ""}

		if err != nil {
			log.Println(err)
			result.Error = err.Error()
			json.NewEncoder(w).Encode(result)
			return
		}
		result.Sum = sum
		json.NewEncoder(w).Encode(result)

	})

	// Start server
	srv := &http.Server{
		Addr:         "0.0.0.0:8090",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	srv.ListenAndServe()

}

func (a Arr) maxPathSum(root map[string]interface{}) (ab int, err error) {
	if root == nil {
		return 0, nil
	}
	maxValue := int(root["value"].(float64))

	// maxSum return the max sum path that starts from root
	var maxSum func(root map[string]interface{}) (int, error)

	maxSum = func(root map[string]interface{}) (int, error) {

		if root == nil {
			return 0, nil
		}
		dataLeft := root["left"]
		left := 0

		switch dataLeft.(type) {
		case string:
			leftData, err := strconv.Atoi(dataLeft.(string))
			if err != nil {
				log.Println(err)
				return 0, errors.New("error 1")
			}
			tmpMaxSum, err := maxSum(a.Find(leftData))
			if err != nil {
				log.Println(err)
				return 0, errors.New("error 2")
			}
			left = max(0, tmpMaxSum)
		case nil:
			left = 0
		default:
			log.Println("error 3")
			return 0, errors.New("error 3")
		}

		dataRight := root["right"]
		right := 0

		switch dataRight.(type) {
		case string:
			rightData, err := strconv.Atoi(dataRight.(string))
			if err != nil {
				log.Println(err)
				return 0, errors.New("error 4")
			}
			tmpMaxSum, err := maxSum(a.Find(rightData))
			if err != nil {
				log.Println(err)
				return 0, errors.New("error 5")
			}
			right = max(0, tmpMaxSum)
		case nil:
			right = 0
		default:
			log.Println("error 6")
			return 0, errors.New("error 6")
		}

		sum := int(root["value"].(float64)) + left + right
		if sum > maxValue {
			maxValue = sum
		}

		return max(left, right) + int(root["value"].(float64)), nil
	}

	if err != nil {
		log.Println(err)
		return 0, errors.New("error 7")
	}

	_, err = maxSum(root)
	if err != nil {
		log.Println(err)
		return 0, errors.New("error 8")
	}

	return maxValue, nil
}

func (a Arr) Find(val int) map[string]interface{} {
	for i, v := range a {
		if v.(map[string]interface{})["value"].(float64) == float64(val) {
			a = append(a[:i], a[i+1:]...)
			return v.(map[string]interface{})
		}
	}
	return nil
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
