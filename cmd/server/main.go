package main

import (
	"flag"
	"fmt"
	"github.com/lizongying/my-tv-server/internal/mytv"
	"github.com/lizongying/my-tv-server/statics"
	"io/fs"
	"log"
	"net/http"
	"strconv"
)

func main() {
	portPtr := flag.String("port", "8000", "")
	assetsPtr := flag.String("assets", "assets", "")
	flag.Parse()
	port, _ := strconv.Atoi(*portPtr)

	log.Printf("Local: http://127.0.0.1:%d/mytv/index.html\n", port)
	log.Printf("Lan: http://%s:%d/mytv/index.html\n", mytv.Lan(), port)
	log.Printf("Internet: http://%s:%d/mytv/index.html\n", mytv.Internet(), port)

	m, _ := mytv.NewMyTv(port)
	e, _ := mytv.NewExample(*assetsPtr, fmt.Sprintf("http://%s:%d/assets/", mytv.Lan(), port))
	ip, _ := mytv.NewIp(port)

	files, _ := fs.Sub(statics.Mytv, "mytv")
	http.Handle("/mytv/", http.StripPrefix("/mytv/", http.FileServer(http.FS(files))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(*assetsPtr))))
	http.HandleFunc("/upload", m.Upload)
	http.HandleFunc("/0", m.Channels)
	http.HandleFunc("/example.json", e.Channels)
	http.HandleFunc("/ip/info", ip.Info)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Println(err)
	}
}
