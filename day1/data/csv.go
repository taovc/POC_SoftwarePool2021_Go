package data

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadFile(path string) ([]string, error) {
	var str_arr []string
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File reading error", err)
		return str_arr, errors.New("error")
	}
	str_arr = strings.Split(string(data), "\n")
	return str_arr, nil
}

func LineToCSV(line string) ([]string, error) {
	var str_arr []string
	str_arr = strings.Split(line, ",")
	fmt.Println(str_arr)
	return str_arr, nil
}
