package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	s := "Ad deserunt sunt in sint aliquip. Proident non sint eu enim incididunt in dolore aute occaecat consectetur consequat. Qui est cillum do elit occaecat laboris adipisicing consectetur eu duis dolor. Do consectetur proident eu esse cupidatat culpa culpa non laboris magna. Culpa nulla consectetur consequat exercitation ipsum dolor veniam velit ipsum nostrud do cillum pariatur. Et nostrud labore et ex proident consectetur anim ex nostrud velit irure. Laborum et pariatur veniam eu magna nulla sit reprehenderit mollit do."

	encodingStd := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	s64 := base64.NewEncoding(encodingStd).EncodeToString([]byte(s))

	fmt.Println(len(s))
	fmt.Println(len(s64))
	fmt.Println(s)
	fmt.Println("\n", s64)

	newS64 := base64.StdEncoding.EncodeToString([]byte(s))
	// newS64 := base64.URLEncoding.EncodeToString([]byte(s))

	fmt.Println("\n", len(s))
	fmt.Println(len(newS64))
	fmt.Println(s)
	fmt.Println("\n", newS64)

	bs, err := base64.StdEncoding.DecodeString(newS64)
	// bs, err := base64.URLEncoding.DecodeString(newS64)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("\n", "Decoded:", string(bs))
}
