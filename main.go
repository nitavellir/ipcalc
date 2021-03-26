package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	opt := os.Args[1]
	input := os.Args[2]
	if opt == "1" {
		res, err := ipToInt(input)
		if err != nil {
			panic(err)
		}
		fmt.Println("IP:", input, ", INT:", res)
	}
}

func ipToInt(input string) (int, error) {
	elems := strings.Split(input, ".")
	if len(elems) != 4 {
		return 0, errors.New("Invalid IP:" + input)
	}

	block1, err := strconv.Atoi(elems[0])
	if err != nil {
		return 0, err
	}
	block2, err := strconv.Atoi(elems[1])
	if err != nil {
		return 0, err
	}
	block3, err := strconv.Atoi(elems[2])
	if err != nil {
		return 0, err
	}
	block4, err := strconv.Atoi(elems[3])
	if err != nil {
		return 0, err
	}

	res := block1*256*256*256 + block2*256*256 + block3*256 + block4
	return res, nil
}
