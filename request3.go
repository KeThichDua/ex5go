package main

import (
	"fmt"
	"log"
	"net/http"
)

// Run3 Sử dụng kiến thức đã đọc tao 1 server net/http hoặc mux sử dụng port 3001 mở 1 url.
// Viết 1 route / trả về pong
func Run3() {
	// mapping url ứng với hàm routing echo
	http.HandleFunc("/", echo)
	// địa chỉ http://127.0.0.1:3001/
	err := http.ListenAndServe(":3001", nil)
	// log ra lỗi nếu bị trùng port
	fmt.Println("Server is running on port 3001..")
	if err != nil {
		log.Fatal(err)
	}

}

// hàm routing echo, gồm hai params
// r *http.Request : dùng để đọc yêu cầu từ client
// wr http.ResponseWriter : dùng để ghi phản hồi về client
func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
}
