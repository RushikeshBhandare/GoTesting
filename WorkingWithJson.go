package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

//unmarshal in json
func main() {
	S := `[{"First":"Sunny","Last":"Bhandare","Age":21},{"First":"Miss","Last":"Mannypenny","Age":34}]`

	BS := []byte(S)

	var People []Person

	err := json.Unmarshal(BS, &People)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range People {
		fmt.Println(v)
	}

}

//arrshling in json
/*


type Persone struct {
	First string
	Last  string
	Age   int
}

func main() {
	PersonOne := Persone{
		First: "Sunny",
		Last:  "Bhandare",
		Age:   21,
	}

	personeTwo := Persone{
		First: "Miss",
		Last:  "Mannypenny",
		Age:   34,
	}

	people := []Persone{PersonOne, personeTwo}
	fmt.Println(people)

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}

*/
