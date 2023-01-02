package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		prefix := strings.HasPrefix(url, "http")
		if !prefix {
			url = "http://" + url
		}
		fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"fetch:%v\n",
				err,
			)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)//使用io.copy替代
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"fetch:reading %s:%v\n",
				url,
				err,
			)
			os.Exit(1)
		}
		fmt.Printf("%s\n", b)
		fmt.Println("http的状态码是：", resp.Status)

	}
}
