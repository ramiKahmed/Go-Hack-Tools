package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"os"
)

func extractCN(host string) {

	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", host, conf)

	if err != nil {
		return
	}

	defer conn.Close()

	certs := conn.ConnectionState().PeerCertificates

	if len(certs) > 1 || len(certs) == 1 {
		for _, cert := range certs {
			fmt.Printf("IP: %v \t CN: %v\n", host, cert.Subject.CommonName[0:])
			break
		}
	}

}

func readFile(filename string) {
	var Hosts []string
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		Hosts = append(Hosts, line)
	}
	for _, host := range Hosts {
		go extractCN(host)

	}
}

func usage(name string) {
	fmt.Fprintf(os.Stdout, "Usage:\t%v -file <filename> \n", name)
	os.Exit(1)
}

func main() {
	if len(os.Args) == 1 {
		usage(os.Args[0])
	}
	var file string
	flag.StringVar(&file, "file", "", "file containing ip:port entries to check for their ssl CN")
	flag.Parse()
	readFile(file)
	var input string
	fmt.Scanln(&input)
}
