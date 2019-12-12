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

	//_, _, err = client.Collection("Users").Add(ctx, map[string]interface{}{
	//	"name": "First User",
	//	"age": 11,
	//	"email": "first@example.com",
	//})
	raw, err := ioutil.ReadFile("./products.json")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var product []Product

	json.Unmarshal(raw, &product)

	for num, _ := range raw {
		fmt.Println(product[num].Title)
		fmt.Println(product[num].Price)
		fmt.Println(product[num].Url)

		_, _, err = client.Collection("hoge2").Add(ctx, Product{
			Title: product[num].Title,
			Price: product[num].Price,
			Url:   product[num].Url,
		})
	}

	if err != nil {
		fmt.Println(err)
	}

}
