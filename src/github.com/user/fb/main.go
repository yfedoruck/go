package main

import (
	"fmt"

	fb "github.com/huandu/facebook"
)

func main() {
	res, _ := fb.Get("/100001773162306", fb.Params{
		"fields":       "first_name",
		"access_token": "EAACEdEose0cBAIoqCOBM00MdZAu2xZB3VZAD5qkLz1YjASIIg8O20ZBkjmtELG4k8KTQd0MGy5EFBWdTS69buVC8fZCvsJ6oyCsxpdM32u3JDsDvOTgDHj5L5eZCnndKutjYOeFD13gxzVky2M1vTyWiHoc0T6ix0pUT9UEkRDq5z4dtplcLzHHJNATcqcB78ZD",
	})
	fmt.Println(res["first_name"])
}
