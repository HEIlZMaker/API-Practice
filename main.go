package main

import (
"net/http"
"fmt"
)

func main() {
	resp, err := http.Get ("https://api.stackexchange.com/2.3/questions?order=desc&sort=activity&site=stackoverflow")
	println(resp)
	println(err)

}