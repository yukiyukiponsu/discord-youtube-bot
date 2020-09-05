package main

import (
	"os"
	"fmt"
	"log"
	// "net/http"
	"github.com/joho/godotenv"
)

func main() {
	//.envファイルの読み込み
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading .env file")
	}
	
	env := os.Getenv("DEVELOPER_KEY")
	fmt.Println(env)

	//指定した単語のyoutubeのリストを返してくれる
	

}
