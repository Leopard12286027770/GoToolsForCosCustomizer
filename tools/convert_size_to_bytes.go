package tools

import (
	"errors"
	"fmt"
	"strconv"
)

const K = 1024
const M = K * 1024
const G = M * 1024

//ConvertSizeToBytes converts a size string to int unit: bytes
//takes a string of number with no unit, unit K, unit M, or unit G
func ConvertSizeToBytes(size string) (int, error) {
	var err error
	res := -1
	l := len(size)
	switch size[l-1] {
	case 'K':
		res, err = strconv.Atoi(size[0 : l-1])
		if Check(err, "cannot convert oemSize to int") {
			return -1, err
		}
		res *= K
	case 'M':
		res, err = strconv.Atoi(size[0 : l-1])
		if Check(err, "cannot convert oemSize to int") {
			return -1, err
		}
		res *= M
	case 'G':
		res, err = strconv.Atoi(size[0 : l-1])
		if Check(err, "cannot convert oemSize to int") {
			return -1, err
		}
		res *= G
	default:
		if size[l-1] >= '0' && size[l-1] <= '9' {
			res, err = strconv.Atoi(size)
			if err != nil {
				fmt.Printf("\nError: cannot convert oemSize to int\n\n")
				return -1, err
			}
		} else {
			return -1, errors.New("Error: wrong format for oemSize")
		}
	}
	return res, nil
}
