package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Catfact struct {
	Fact   string
	Length int
}

func main() {
	// https://catfact.ninja/fact
	resp, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		fmt.Println("Error Get: ", err)
		return
	}
	// fmt.Println(resp)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error Read: ", err)
		return
	}

	var fact Catfact
	json.Unmarshal(body, &fact)
	//fmt.Printf("%#v", string(body))
	fmt.Println(fact.Fact)
}
