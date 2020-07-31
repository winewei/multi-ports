package main

import (
    "fmt"
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
            fmt.Printf("tcp listen port%s\n", port)
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
