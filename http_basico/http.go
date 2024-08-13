package main 

import (
	"net/http"
	"fmt"
	"io"
)

func main() {
	req, err := http.Get("https://www.x.com")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close() // defer deixa pra executar a linha por último
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
	
}