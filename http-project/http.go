package httpproject

import (
	"fmt"
	"io"
	"net/http"
)

func HttpEntry() {
	method := 1
	switch method {
	case 1:
		MakeGetRequest()
	}
}

type LogWriter struct{}

func MakeGetRequest() error {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// io.Copy(os.Stdout, resp.Body) // writer

	lw := LogWriter{}
	io.Copy(lw, resp.Body)
	return nil
}

// Custom writer
func (LogWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
