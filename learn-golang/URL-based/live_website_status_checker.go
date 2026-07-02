package main

import(
	"fmt"
	"net/http"
	"time"
)

func checkStatus(url string, result chan string) {
	resp, err := http.Get(url)
	if err != nil{
		result <- fmt.Sprintf("%s is DOWN: %v", url, err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK{
		result <- fmt.Sprintf("%s is ACTIVE: %s", url, resp.Status)
	}else{
		result <- fmt.Sprintf("%s returned Error - STATUS: %s", url, resp.Status)
	}
}

func main() {
	start := time.Now()

	result := make(chan string)
	links := []string{
		"https://www.google.com",               
		"https://httpbin.org/status/404",       
		"https://httpbin.org/status/500",       
		"https://httpbin.org/status/403",       
		"https://this-domain-does-not-exist.biz",
	}

	for _, link := range links{
		go checkStatus(link, result)
	}

	for i := 0; i < len(links); i += 1 {
		msg := <- result
		fmt.Println(msg)
	}

	fmt.Println("Process over. Time taken: ", time.Since(start))

}