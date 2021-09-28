package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"net"
	"net/http"
)

type msg struct {
	URL    string
	Header map[string][]string
	Body   []byte
}

var server string
var localPort string

func main() {
	flag.StringVar(&localPort, "port", "8080", "local port")
	flag.StringVar(&server, "server", "awsfelipe.ddns.net:7891", "server host")
	flag.Parse()

	add, _ := net.ResolveTCPAddr("tcp", server)
	conn, err := net.DialTCP("tcp", nil, add)
	if err != nil {
		log.Fatal(err)
	}
	for {
		s := bufio.NewScanner(conn)
		for s.Scan() {
			var m msg
			json.Unmarshal(s.Bytes(), &m)

			req, err := http.NewRequest("POST", "http://localhost:"+localPort+m.URL, bytes.NewBuffer(m.Body))
			if err != nil {
				log.Println(err)
			}
			req.Header = m.Header
			_, err = (&http.Client{}).Do(req)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
