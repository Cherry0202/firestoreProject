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

	raw, err := ioutil.ReadFile("./products.json")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var product []Product

	json.Unmarshal(raw, &product)

	batch := client.Batch()
	//Ref := client.Collection("Products").Doc("Mac")
	start := time.Now()
	fmt.Println("スタート！！")
	// for _, p := range product {
	// 	batch.Set(Ref, p)
	// }
	wg := sync.WaitGroup{}
	for num, p := range product {
		//fmt.Println(product[num].Title)
		//fmt.Println(product[num].Price)
		//fmt.Println(product[num].Url)

		//Ref := client.Collection("Products3").Doc(strconv.Itoa(num))
		//_, err = Ref.Set(ctx, product[num])
		//go SetBatch(wg, batch, client, p, num)
		SetBatch(wg, batch, client, p, num)
	}
	wg.Wait()
	_, err = batch.Commit(ctx)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
	end := time.Now()
	fmt.Printf("かかった時間：%f秒\n", (end.Sub(start)).Seconds())

	defer func() {
		err := recover()
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func SetBatch(wg sync.WaitGroup, batch *firestore.WriteBatch, client *firestore.Client, p Product, num int) {
	//defer wg.Done()
	Ref := client.Collection("Products5").Doc(strconv.Itoa(num))
	batch.Set(Ref, p)
}
