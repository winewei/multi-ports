package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"
    "strings"
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
            log.Printf("server start, tcp listen port%s\n", port)
            mux := http.NewServeMux()
            mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                d := Data{}
                d.Port = strings.Split(port, ":")[1]
                d.Host, _ = os.Hostname()
                data, _ := json.Marshal(d)
                go func() {
                    log.Println("listen port:", d.Port, "hostname:", d.Host)
                }()
                w.Header().Set("Content-type", "application/json")
                w.Write([]byte(data))
            })
            http.ListenAndServe(port, mux)
        }(k)
    }
    select {}
}
