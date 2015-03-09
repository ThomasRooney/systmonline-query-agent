package main

import "fmt"
import "net/http"
import "net/url"
import "crypto/tls"
import "gopkg.in/yaml.v2"
import "io/ioutil"

var configFileName = "config.yml"
var secretFileName = "secret.yml"

type Secret struct {
	username string
	password string
}

func main() {
	secret := Secret{}
	secretData, err := ioutil.ReadFile(secretFileName)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(secretData), &secret)
	if err != nil {
		panic(err)
	}

	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.PostForm("https://systmonline.tpp-uk.com/Login",
		url.Values{"username": {secret.username}, "password": {secret.password}})
	if err != nil {
		fmt.Printf("Unexpected error[%+#v].\n", err)
	}
	if resp.Status != "200 OK" {
		fmt.Printf("Unexpected HTTP Response [%s]\n", resp.Status)
	}
	fmt.Printf("%+#v", resp)
}
