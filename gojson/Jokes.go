package main

import "fmt"
import "github.com/ChimeraCoder/gojson"
import "os"
import "io/ioutil"
import "encoding/json"

type Joke struct {
	Type  string `json:"type"`
	Value struct {
		Categories []interface{} `json:"categories"`
		ID         int           `json:"id"`
		Joke       string        `json:"joke"`
	} `json:"value"`
}

func main() {

	i, err := os.Open("example.json")
	if err != nil {
		fmt.Println(err)
	}

	filecontent, err := ioutil.ReadFile("example.json")
	if err != nil {
		fmt.Println(err)
	}

	chuck := Joke{}
	json.Unmarshal(filecontent, &chuck)

	fmt.Println(chuck.Value.Joke)

	actual, err := json2struct.Generate(i, "Joke", "main")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(actual))

}
