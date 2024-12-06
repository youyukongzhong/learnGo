package main

import (
	"encoding/json"
	"fmt"
)

type OrderItem struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Order struct { //tag
	ID         string `json:"id"`
	Items      []OrderItem
	TotalPrice float64 `json:"total_price"`
}

func main() {
	parseNLP()
}

func marshal() {
	o := Order{
		ID:         "1234",
		TotalPrice: 20,
		Items: []OrderItem{
			{
				ID:    "item_1",
				Name:  "learn go",
				Price: 15,
			},
			{
				ID:    "item_2",
				Name:  "interview",
				Price: 10,
			},
		},
	}

	b, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", b)
	//fmt.Printf("%+v\n", o)
}

func unmarshal() {
	s := `{"id":"1234","Items":[{"id":"item_1","name":"learn go","price":15},{"id":"item_2","name":"interview","price":10}],"total_price":20}`
	var o Order
	err := json.Unmarshal([]byte(s), &o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", o)
}

func parseNLP() {
	res := `{
  "result": [
    {
      "synonym": "",
      "weight": "0.100000",
      "tag": "普通词",
      "word": "请"
    },
    {
      "synonym": "",
      "weight": "0.100000",
      "tag": "普通词",
      "word": "输入"
    },
    {
      "synonym": "双手合十",
      "weight": "1.000000",
      "tag": "品类",
      "word": "文本"
    }
  ]
}`
	m := struct {
		Result []struct {
			Synonym string `json:"synonym"`
			Tag     string `json:"tag"`
		} `json:"result"`
	}{}
	err := json.Unmarshal([]byte(res), &m)
	if err != err {
		panic(err)
	}

	fmt.Printf("%+v,\n %+v,", m.Result[2].Synonym, m.Result[2].Tag)
}
