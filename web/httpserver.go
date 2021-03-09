package web

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	addr string
}

func (s *Server) SayHelloHandler(w http.ResponseWriter, r *http.Request) {
	upath := fmt.Sprintf("http://%s%s\n", s.addr, r.URL.Path)
	io.WriteString(w, upath)
}

func (s *Server) LoggingHandler(w http.ResponseWriter, r *http.Request) {
	upath := "errorlog path"
	w.WriteHeader(500)
	io.WriteString(w, upath)
}

func (s *Server) Run() {
	log.Println(s.addr, "server run")
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.SayHelloHandler)
	mux.HandleFunc("/base/error", s.LoggingHandler)
	server := &http.Server{
		Addr:         s.addr,
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()
}

func Start() {
	s1 := &Server{addr: "127.0.0.1:2003"}
	s1.Run()
	s2 := &Server{addr: "127.0.0.1:2004"}
	s2.Run()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
