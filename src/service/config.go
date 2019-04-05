package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
)

var (
	NEWLINE = []byte{'\n'}
)

func trimComments(data []byte) (dataWithoutComments []byte) {
	lines := bytes.Split(data, NEWLINE)
	for k, line := range lines {
		lines[k] = removeCommentsLine(line)
	}
	return bytes.Join(lines, NEWLINE)
}
func removeCommentsLine(line []byte) []byte {

	var newLine []byte
	var i, quoteCount int
	lastIdx := len(line) - 1
	for i = 0; i <= lastIdx; i++ {
		if line[i] == '\\' {
			if i != lastIdx && (line[i+1] == '\\' || line[i+1] == '"') {
				newLine = append(newLine, line[i], line[i+1])
				i++
				continue
			}
		}
		if line[i] == '"' {
			quoteCount++
		}
		if line[i] == '#' {
			if quoteCount%2 == 0 {
				break
			}
		}
		newLine = append(newLine, line[i])
	}
	return newLine
}

func parseConfig(configFile string, ret interface{}) (err error) {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return
	}
	data = trimComments(data)

	err = json.Unmarshal(data, ret)
	if err != nil {
		Fatal("unmarshal failed.")
	}
	return
}

func ReadConfig(config interface{}) error {
	configFileName := flag.String("f", "", "./main -f config.conf")
	flag.Parse()
	if *configFileName == "" {
		return errors.New("no config file provided!")
	}
	fmt.Println("load config file:", *configFileName)
	return parseConfig(*configFileName, config)
}
