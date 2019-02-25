package main

import (
	"./lib"
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

var usage = `MySQL HTTP Agent - HTTP interface

This is an HTTP interface for the MySQL HTTP Agent.

Run:

	agent [-host=<host>] [-port=<port>]

Example:

	$ curl -s --request POST \
	--header 'Content-Type: application/json' \
	--url 'http://localhost:5000/' \
	--data '
	{
		"connection": {
			"host": "localhost",
			"port": 3306,
			"user": "root",
			"password": "123456",
			"database": "test"
		},
		"sql": "SELECT name, email, status FROM t_user WHERE id = :id",
		"params": {
			"id": 1
		}
	}
`

var (
	host string // 监听主机名
	port string // 监听端口号
)

type Payload struct {
	//Driver     string                 `yaml:"driver"`
	Connection map[string]interface{} `yaml:"connection"`
	SQL        string                 `yaml:"sql"`
	Params     map[string]interface{} `yaml:"params"`
}

func init() {

	flag.StringVar(&host, "host", "localhost", "Host of the agent.")
	flag.StringVar(&port, "port", "5000", "Port of the agent.")

	flag.Usage = func() {
		fmt.Println(usage)
		flag.PrintDefaults()
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)

}

func main() {

	flag.Parse()

	addr := net.JoinHostPort(host, port)
	log.Printf("* Listening on %s...\n", addr)

	http.HandleFunc("/", requestHandler)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}

}

func requestHandler(writer http.ResponseWriter, request *http.Request) {

	if request.Method == "GET" || request.Method == "HEAD" {
		return
	}

	if request.Method != "POST" {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// 读取请求消息体
	data, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		writer.Write([]byte(fmt.Sprintf("could not read body: %s", err)))
		return
	}

	// 解析消息体，使用Payload结构化
	var payload Payload
	if err := yaml.Unmarshal(data, &payload); err != nil {
		writer.WriteHeader(http.StatusUnprocessableEntity)
		writer.Write([]byte(fmt.Sprintf("could not decode body: %s", err)))
		return
	}

	// 检查数据库连接
	db, err := lib.Connect(&payload.Connection)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(fmt.Sprintf("error encoding data: %s", err)))
		return
	}

	// 执行查询请求
	rows, err := lib.Query(db, payload.SQL, payload.Params)
	if err != nil {
		writer.WriteHeader(http.StatusServiceUnavailable)
		writer.Write([]byte(fmt.Sprintf("error executing query: %s", err)))
		return
	}

	// 遍历结果迭代器，转换为JSON数据类型
	var lst []map[string]interface{}
	for rows.Next() {
		u := make(map[string]interface{})
		rows.MapScan(u)
		lst = append(lst, u)
	}

	// 转换为JSON字符串
	j, _ := lib.EncodeJson(&lst)
	/*
		[
		  {
			"cnd": 120
		  }
		]
	*/
	fmt.Println(string(j))

	// 输出到响应消息体中
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(j)

	return
}
