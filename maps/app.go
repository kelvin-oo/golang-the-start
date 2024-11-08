package main

import "fmt"

func main() {
	websites := map[string]string{
		"amazon":  "https://aws.com",
		"del":     "del.com",
		"adobe":   "adobe.com",
		"samsung": "samsung.com",
		"toshiba": "toshiba.com",
		"hisense": "hisense.com",
	}
	// fmt.Println(websites)
	websites["google"] = "https://google.com"
	// fmt.Println(websites)
	delete(websites, "amazon")
	// fmt.Println(websites)

	for _, value := range websites {
		// fmt.Println(index)
		fmt.Println(value)
	}
}
