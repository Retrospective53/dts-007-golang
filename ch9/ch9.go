package ch9

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

type WaternWind struct {
	Water int `json:"water"`
	Wind int `json:"wind"`
}

func PostJsonPlaceHolder() {
	waterValue, windValue := generateWnW()

	// data := map[string]interface{}{
	// 	"water":  waterValue,
	// 	"wind":   windValue,
	// }

	data := WaternWind{Water: waterValue, Wind: windValue}

	requestJson, err := json.Marshal(data)
	client := &http.Client{}

	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts",
	bytes.NewBuffer(requestJson))
	req.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sbody := string(body)
	// extract the JSON object from the response string
	jsonStr := sbody[strings.Index(sbody, "{"):]

	var dataBody WaternWind
	if err := json.Unmarshal([]byte(jsonStr), &dataBody); err != nil {
		panic(err)
	}

	// convert the updated Response struct back to a JSON string
	updatedJson, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	// print the updated response string
	fmt.Println(string(updatedJson))

	waterStatus := waterStatus(waterValue)
	fmt.Printf("status water: %s\n", waterStatus)

	windStatus := windStatus(windValue)
	fmt.Printf("status wind: %s\n", windStatus)

	// fmt.Println(string(body))
}

func generateWnW() (waterValue int, windValue int) {
	waterValue = rand.Intn(100) + 1
	windValue = rand.Intn(100) + 1

	// // fmt.Printf("water %dm\n", waterNum)
	// // fmt.Printf("wind %dm/s\n", windNum)

	return
}

func waterStatus(waterNum int) (status string) {
	if waterNum < 5 {
		status = "aman"
	} else if waterNum <= 8 {
		status = "siaga"
	} else {
		status = "bahaya"
	}
	return
}

func windStatus(windNum int) (status string) {
	if windNum < 6 {
		status = "aman"
	} else if windNum <= 15 {
		status = "siaga"
	} else {
		status = "bahaya"
	}
	return
}

