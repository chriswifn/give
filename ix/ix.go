package ix

import (
	"fmt"
	"github.com/rwxrob/to"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func post(args ...any) string {
	var in string
	ln := len(args)
	switch {
	case ln >= 2:
		fallthrough
	case ln == 1:
		in = to.String(args[0])
	case ln == 0:
		return ""
	}
	url := "http://ix.io"
	form := fmt.Sprintf("f:1=%s", in)
	payload := strings.NewReader(form)
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Println(err)
		return ""
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}

func get(args ...any) string {
	var in string
	ln := len(args)
	switch {
	case ln >= 2:
		fallthrough
	case ln == 1:
		in = strings.TrimSpace(to.String(args[0]))
	case ln == 0:
		return ""
	}
	url := fmt.Sprintf("http://ix.io/%s", in)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return ""
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return ""
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}
