package http

import (
    "net/http"
    "fmt"	
    "log"
    "encoding/json"
)

var (
    client = &http.Client{}
    token = "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJ3SkJ5SWxzTWU4UGFJS0MySi1oYWVrMWYxZW5JYVQtQ2VLdTdKazhPX2EwIn0.eyJleHAiOjE2MDk5NjIzOTEsImlhdCI6MTYwOTkxOTE5MSwianRpIjoiOTVmMDE5NmMtNjA0OC00MzM4LWI2OTMtNWY1NWYwMzY3N2IxIiwiaXNzIjoiaHR0cHM6Ly9pYW0uaW50Lm1hbm9tYW5vLmNvbS9hdXRoL3JlYWxtcy9NYXJrZXRwbGFjZSIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiJmOmFjMGRmMjA5LWJlMTktNGZmMi04OThkLTZhZGIyMzcxMTRhMzpiYXN0aWVuLmxhY29tYmUiLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJmcm9udC1ibyIsInNlc3Npb25fc3RhdGUiOiJlZmNjZDQ5MC1mYzI4LTQ1ZDYtYjMxNi03ZTg1MWRlZGZjMzUiLCJhY3IiOiIxIiwiYWxsb3dlZC1vcmlnaW5zIjpbIioiXSwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwiUk9MRV9BRE1JTklTVFJBVE9SIiwidW1hX2F1dGhvcml6YXRpb24iLCJST0xFX01NX1RCX0FETUlOIl19LCJyZXNvdXJjZV9hY2Nlc3MiOnsiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJvcGVuaWQgdXNlcm1hbm8gdXNlcm1hbm9fYWQgZW1haWwgcHJvZmlsZSIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwibmFtZSI6IkJhc3RpZW4gTEFDT01CRSIsInByZWZlcnJlZF91c2VybmFtZSI6ImJhc3RpZW4ubGFjb21iZSIsImdpdmVuX25hbWUiOiJCYXN0aWVuIiwidXVpZCI6IjRiNjVmMzU0LTI1MTUtNDU1Mi1iZGE0LWY0ODBiNTJkNzA1ZSIsImZhbWlseV9uYW1lIjoiTEFDT01CRSIsImVtYWlsIjoiYmFzdGllbi5sYWNvbWJlQG1hbm9tYW5vLmNvbSJ9.lDriypm8-3k0wLhCMJOgCdlJgS0GqKTzbOqUYAQn60IpPWYWFQKwfaqsgu5KfSYc5PEzE6CFxabwLZVB8Cf20RK6I9lL5JQeg1mVGum0l6eQJ3QcuPVJLdcCcRMNdLdJvk4H1UTICYPHjcg-JdZNduijFF5L-ZPFoK2aWjlP50mjtpJggwLz7WZZtJpJ6uAddv65xKGPK98gV1hWU-etIY5VtW6ljpvI5YX9IPeXh4IuouoiZ21xfTXNKnA5IHBoPuspQPGoXK-SC-oB8Hp74-Xk3IYo1BRN3prZBXVo0EWqbwXqd6yWW1-i38MxF3cabdRuu6e6bN4DwVauIcGSag"
    headers map[string]string = map[string]string{"Authorization": "Bearer " + token}
    message map[string]interface{}
)  

func Get(url string) {
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
