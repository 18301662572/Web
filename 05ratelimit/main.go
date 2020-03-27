package main
import (
"io"
"log"
"net/http"
)

//衡量一下这个 web 服务的吞吐量，再具体一些，实际上就是接口的 QPS
//测试：wrk -c 10 -d 10s -t10 http://localhost:9090

func sayhello(wr http.ResponseWriter, r *http.Request) {
	wr.WriteHeader(200)
	io.WriteString(wr, "hello world")
}

func main1() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
