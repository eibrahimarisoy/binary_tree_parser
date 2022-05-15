package main

import (
	"fmt"
	"log"
	"strconv"
)

type NodeArr []interface{}

// maxPathSum returns the maximum sum of a path in a binary tree
func (a NodeArr) maxPathSum(root map[string]interface{}) (ab int, err error) {
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
				return 0, fmt.Errorf("%v", err)
			}
			tmpMaxSum, err := maxSum(a.Find(leftData))
			if err != nil {
				log.Println(err)
				return 0, fmt.Errorf("%v", err)
			}
			left = max(0, tmpMaxSum)
		case nil:
			left = 0
		default:
			message := fmt.Sprintf("Left Value is not string. value:%v", dataLeft)
			log.Printf(message)
			return 0, fmt.Errorf(message)
		}

		dataRight := root["right"]
		right := 0

		switch dataRight.(type) {
		case string:
			rightData, err := strconv.Atoi(dataRight.(string))
			if err != nil {
				log.Println(err)
				return 0, fmt.Errorf("%v", err)
			}
			tmpMaxSum, err := maxSum(a.Find(rightData))
			if err != nil {
				log.Println(err)
				return 0, fmt.Errorf("%v", err)
			}
			right = max(0, tmpMaxSum)
		case nil:
			right = 0
		default:
			message := fmt.Sprintf("Right Value is not string. value:%v", dataRight)
			log.Printf(message)
			return 0, fmt.Errorf(message)
		}

		sum := int(root["value"].(float64)) + left + right
		if sum > maxValue {
			maxValue = sum
		}

		return max(left, right) + int(root["value"].(float64)), nil
	}

	if err != nil {
		log.Println(err)
		return 0, fmt.Errorf("%v", err)
	}

	_, err = maxSum(root)
	if err != nil {
		log.Println(err)
		return 0, fmt.Errorf("%v", err)
	}

	return maxValue, nil
}

// Find returns the node with the given value
func (a NodeArr) Find(val int) map[string]interface{} {
	for i, v := range a {
		if v.(map[string]interface{})["value"].(float64) == float64(val) {
			a = append(a[:i], a[i+1:]...)
			return v.(map[string]interface{})
		}
	}
	return nil
}

// max returns the max of two ints
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
