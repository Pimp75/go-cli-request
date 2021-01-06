package http

import (
    "net/http"
    "fmt"	
    "log"
    "encoding/json"
)

var (
    client = &http.Client{}
    message map[string]interface{}
)  

func Get(url string, headers map[string]string) {
    req, _ := http.NewRequest("GET", url, nil)

    for key, value := range headers{
    	req.Header.Add(key, value)
    }
    resp, err := client.Do(req)

    if err != nil {
	log.Println("Error on response.\n[ERROR] -", err)
    } 
    defer resp.Body.Close()
    
    json.NewDecoder(resp.Body).Decode(&message)

    fmt.Println(message)
}
