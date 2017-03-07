package main

import (
	"crypto/tls"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	certFile = flag.String("cert", "usercert.pem", "A PEM eoncoded certificate file.")
	keyFile  = flag.String("key", "unencrypted.pem", "A PEM encoded private key file.")
)

func requestHandler(url string, uri string, verb string, cert string, key string) error{

	// Load client cert
	certificate, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		log.Fatal(err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{certificate},
        InsecureSkipVerify: true,
	}

    tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

    req, _ := http.NewRequest(verb, url + uri, nil)

	log.Println("start")
    req.Header.Set("Accept", "application/json")
    resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
        return err
	}
	defer resp.Body.Close()

	// Dump response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
        return err
	}
	log.Println(string(data))
    return nil
}

func main() {
	flag.Parse()
    requestHandler("https://cmsweb-testbed.cern.ch/crabserver/preprod/filetransfers",
                   "?subresource=acquiredTransfers&asoworker=asodciangot1&grouping=0", "GET", *certFile, *keyFile)
}

