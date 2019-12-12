package main

import (
	"context"
	"encoding/json"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type Product struct {
	Title string `json:"title"`
	Price string `json:"price"`
	Url   string `json:"url"`
}

func main() {

	defer func() {
		err := recover()
		if err != nil {
			log.Fatal(err)
		}
	}()

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

	//batch := client.Batch()
	Ref := client.Collection("Products").Doc("Mac")

	for num, _ := range raw {
		fmt.Println(product[num].Title)
		fmt.Println(product[num].Price)
		fmt.Println(product[num].Url)

		Ref = client.Collection("Products").Doc(strconv.Itoa(num))
		_, err = Ref.Set(ctx, Product{
			Title: product[num].Title,
			Price: product[num].Price,
			Url:   product[num].Url,
		})
	}

	if err != nil {
		fmt.Println(err)
	}

}
