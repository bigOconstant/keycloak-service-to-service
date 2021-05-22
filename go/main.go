package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"net/http"
	"net/url"
	"encoding/json"
)

func main() {
	
	resp, _ := http.PostForm(os.Getenv("TOKENURL"),
		url.Values{"grant_type": {"client_credentials"}, "client_id": {os.Getenv("CLIENTID")},"client_secret": {os.Getenv("CLIENTSECRET")}})

	defer resp.Body.Close()
	
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		c := make(map[string]json.RawMessage)
		err = json.Unmarshal(bodyBytes, &c)
		if err != nil {
			panic(err)
		}
		fmt.Println("Token expires in:",string(c["expires_in"]), " seconds \n")
		fmt.Println("Token: ",string(c["access_token"]))
	}
	
}
