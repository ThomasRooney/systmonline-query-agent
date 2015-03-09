package main

import "fmt"
import "net/http"
import "crypto/tls"

func main() {
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("https://example.com")
	if err != nil {
		fmt.Printf("Unexpected error[%+#v].\n", err)
	}
	if resp.Status != "200 OK" {
		fmt.Printf("Unexpected HTTP Response [%s]\n", resp.Status)
	}
	fmt.Printf("%+#v", resp)
}
