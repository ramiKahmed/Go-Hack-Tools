package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func removeDuplicateValues(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {

	var domain string
	var hostnames []string

	flag.StringVar(&domain, "domain", "", "find hostnames related to the main company domain")
	flag.Parse()

	host := domain + ":" + "443"

	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stdout, "Usage:\t%v -domain domainName\n", os.Args[0])
		fmt.Fprintf(os.Stdout, "Ex:\t%v -domain google.com\n", os.Args[0])
		os.Exit(1)
	}

	conn, err := tls.Dial("tcp", host, nil)
	if err != nil {
		log.Println("Error in Dial", err)
		return
	}
	defer conn.Close()
	certs := conn.ConnectionState().PeerCertificates
	for _, cert := range certs {
		for i := 0; i < len(cert.DNSNames); i++ {
			name := strings.Contains(cert.DNSNames[i], "*")
			if name {
				name := strings.Trim(cert.DNSNames[i], "*.")
				//fmt.Println(name)
				hostnames = append(hostnames, name)
			} else {

				hostnames = append(hostnames, cert.DNSNames[i])
			}

		}

	}
	removeDuplicateValuesSlice := removeDuplicateValues(hostnames)
	for i := 0; i < len(removeDuplicateValuesSlice); i++ {
		fmt.Printf("%v\n", removeDuplicateValuesSlice[i])
	}
	fmt.Printf("\n\nTotal unique hosts found: [ %v ]\n", len(removeDuplicateValuesSlice))
}