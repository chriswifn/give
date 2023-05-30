package server

import (
	"fmt"
	"github.com/skip2/go-qrcode"
	"html/template"
	"io"
	"log"
	"net"
)

func RenderString(s string) {
	q, err := qrcode.New(s, qrcode.Medium)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(q.ToSmallString(false))
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func serveTemplate(name string, tmpl string, w io.Writer, data interface{}) {
	t, err := template.New(name).Parse(tmpl)
	if err != nil {
		panic(err)
	}
	if err := t.Execute(w, data); err != nil {
		panic(err)
	}
}
