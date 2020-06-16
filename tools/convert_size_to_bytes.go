package tools

import (
	"errors"
	"fmt"
	"strconv"
)

const B = 1
const K = 1024
const M = K * 1024
const G = M * 1024
const SEC = 512

//ConvertSizeToBytes converts a size string to int unit: bytes
//takes a string of number with no unit (sectors), unit B, unit K, unit M, or unit G
func ConvertSizeToBytes(size string) (int, error) {
	var err error
	res := -1
	l := len(size)

	if l <= 0 {
		return -1, errors.New("invalid oemSize: empty string")
	}

	switch size[l-1] {
	case 'B':
		res, err = strconv.Atoi(size[0 : l-1])
		if Check(err, "cannot convert oemSize to int") {
			return -1, err
		}
		res *= B
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
			res *= SEC
		} else {
			return -1, errors.New("Error: wrong format for oemSize")
		}
	}
	return res, nil
}
