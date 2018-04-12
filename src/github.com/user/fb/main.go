package main

import (
	"fmt"

	fb "github.com/huandu/facebook"
)

func main() {
	tok := "EAACEdEose0cBAIoqCOBM00MdZAu2xZB3VZAD5qkLz1YjASIIg8O20ZBkjmtELG4k8KTQd0MGy5EFBWdTS69buVC8fZCvsJ6oyCsxpdM32u3JDsDvOTgDHj5L5eZCnndKutjYOeFD13gxzVky2M1vTyWiHoc0T6ix0pUT9UEkRDq5z4dtplcLzHHJNATcqcB78ZD"
	res, _ := fb.Get("/831976246906620/feed", fb.Params{
		"access_token": tok,
	})

	var items []fb.Result
	err := res.DecodeField("data", &items)

	if err != nil {
		fmt.Printf("An error has happened %v", err)
		return
	}

	for _, item := range items {
		fmt.Println(item["message"])
	}
}