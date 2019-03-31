package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
)

type MongoDbConfig struct {
	Host           string `json:"host"`
	Database       string `json:"db"`
	CollectionName string `json:"coll"`
	Mode           string `json:"mode"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	SyncTimeoutSec int64  `json:"timeout"` // 以秒为单位
	Pool           int    `json:"pool"`
}

type Config struct {
	MaxProcs int           `json:"max_procs"`
	DbConfig MongoDbConfig `json:"mongo"`
}

var (
	NEWLINE = []byte{'\n'}
)

func Fatal(v ...interface{}) {
	os.Stderr.WriteString(fmt.Sprint(v, "\n"))
	os.Exit(1)
}
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
	configFileName := flag.String("f", "", "simple_server -f config_file_name")
	flag.Parse()
	if *configFileName == "" {
		return errors.New("no config file provided!")
	}
	fmt.Println("load config file:", *configFileName)
	return parseConfig(*configFileName, config)
}

func main() {
	fmt.Println("simple server with mongodb")
	var config Config
	if err := ReadConfig(&config); err != nil {
		Fatal("ReadConfig Failed. ", err)
	}

	runtime.GOMAXPROCS(config.MaxProcs)

	dbSesstion := ConnectMongo(&config.DbConfig)
	defer dbSesstion.Close()

	fmt.Println(config.DbConfig.Database)

	collection := dbSesstion.DB(config.DbConfig.Database).C(config.DbConfig.CollectionName)
	collection.Insert(DocItem{"bucketone", 1234})

	fmt.Println("Fin.")
}
