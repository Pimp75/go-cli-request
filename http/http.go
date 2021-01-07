package http

import (
    "net/http"
    "fmt"	
    "log"
    "encoding/json"
    "bytes"
)

var (
    client = &http.Client{}
    message map[string]interface{}
)  

func Call(method string, url string, headers map[string]string, body string) {
    body_bytes :=  []byte(body)

    req, _ := http.NewRequest(method, url, bytes.NewBuffer(body_bytes))

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
