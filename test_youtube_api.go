package main

import (
	"os"
	"fmt"
	"log"
	"time"
	"net/http"
	"io/ioutil"
	"github.com/joho/godotenv"
)

func main() {
	developer_key := get_developer_key()
	youtube_search_list := search_youtube_list(developer_key)
	fmt.Println(youtube_search_list)
}

func get_developer_key() string{
	//.envファイルの読み込み
	err := godotenv.Load()
	if(err != nil){
		log.Fatal("Error loading .env file")
	}
	
	developer_key := os.Getenv("DEVELOPER_KEY")
	fmt.Println(developer_key)
	return developer_key
}

func search_youtube_list(developer_key string) string{
	url := "https://www.googleapis.com/youtube/v3/search"
 
	request, err := http.NewRequest("GET", url, nil)
	if err != nil{
		log.Fatal(err)
	}
	
	//クエリパラメータ
	params := request.URL.Query()
	params.Add("key", developer_key)
	params.Add("q", "洋楽")
	params.Add("part", "snippet, id")
	params.Add("maxResults", "1")

    request.URL.RawQuery = params.Encode()
 
	fmt.Println(request.URL.String()) //https://jsonplaceholder.typicode.com/todos?userId=1
	
	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
        Timeout: timeout,
	}
 
	response, err := client.Do(request)
	if err != nil{
		log.Fatal(err)
	}
	
	defer response.Body.Close()
 
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
 
	// fmt.Println(string(body))
	return string(body)
}