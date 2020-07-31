package main

import (
    "os"
    "strings"
    "encoding/json"
    "net/http"
)

type Data struct {
   Port string
   Host string
}

func main() {
    ports := strings.Split(os.Getenv("PORTS"), ",")
    for _, v := range ports {
        k := ":" + v
        go func(port string) {
           mux := http.NewServeMux()
           mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
               d := &Data{}
               d.Port = port
               d.Host, _ = os.Hostname()
               data, _ := json.Marshal(d)
               w.Write([]byte(data))
           })
           http.ListenAndServe(port, mux)
        }(k)
    }
    select {}
}
