package main

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

type Product struct {
	Title string `json:"title"`
	Price string `json:"price"`
	Url   string `json:"url"`
}

func main() {

	//初期化
	ctx := context.Background()
	opt := option.WithCredentialsFile("secret.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		fmt.Println(err)
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		fmt.Println(err)
	}

	wg := sync.WaitGroup{}
	raw, err := ioutil.ReadFile("./products.json")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var product []Product

	json.Unmarshal(raw, &product)

	batch := client.Batch()
	start := time.Now()
	fmt.Println("スタート")
	for num, p := range product {
		//go SetBatch(wg, batch, client, p, num)
		SetBatch(wg, batch, client, p, num)
	}
	wg.Wait()
	_, err = batch.Commit(ctx)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
	end := time.Now()
	fmt.Printf("かかった時間：%f秒\n", (end.Sub(start)).Seconds())
}

func SetBatch(wg sync.WaitGroup, batch *firestore.WriteBatch, client *firestore.Client, p Product, num int) {
	//defer wg.Done()
	Ref := client.Collection("Product6").Doc(strconv.Itoa(num))
	batch.Set(Ref, p)
}
