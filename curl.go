package main

import (
    "fmt"
    "github.com/andelf/go-curl"
    "net/url"
)

func main() {
    easy := curl.EasyInit()
    defer easy.Cleanup()


    request := `{"subresource":"acquiredTransfers","asoworker":"asodciangot1", "grouping":0}`

    //urls := url.Parse(request)

    fmt.Printf(request)

    easy.Setopt(curl.OPT_URL, "https://cmsweb-testbed.cern.ch/crabserver/preprod/filetransfers?subresource=acquiredTransfers&asoworker=asodciangot1&grouping=0")

    // make a callback function
    fooTest := func (buf []byte, userdata interface{}) bool {
        println("DEBUG: size=>", len(buf))
        println("DEBUG: content=>", string(buf))
        return true
    }

    easy.Setopt(curl.OPT_WRITEFUNCTION, fooTest)
    easy.Setopt(curl.OPT_SSL_VERIFYPEER, false)
    easy.Setopt(curl.OPT_SSLCERT, "usercert.pem")
    easy.Setopt(curl.OPT_SSLKEY, "unencrypted.pem")
    easy.Setopt(curl.OPT_VERBOSE, 1)

    if err := easy.Perform(); err != nil {
        fmt.Printf("ERROR: %v\n", err)
    }
}
