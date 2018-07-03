package main

import (
    "os"
    "fmt"
    "github.com/ChimeraCoder/anaconda"
)

func main() {
	fmt.Println(os.Getenv("CONSUMER_KEY"))
	fmt.Println(os.Getenv("CONSUMER_SECRET"))
	fmt.Println(os.Getenv("ACCESS_TOKEN"))
	fmt.Println(os.Getenv("ACCESS_TOKEN_SECRET "))
    anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
    anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
    api := anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
    _,err := api.PostTweet("Goでボット書きました。",nil)
    if err != nil{
    	fmt.Println(err)
    }
}
