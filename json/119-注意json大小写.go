package main

import (
	"encoding/json"
	"fmt"
)

type One119 struct {
	NameI18n string
}
type Two119 struct {
	NameI18n string `json:"name_i18n"`
}

func main() {
	s := `{"name_i18n":"haha"}`
	x := One119{}
	json.Unmarshal([]byte(s), &x)
	fmt.Println(x.NameI18n)
	y := Two119{}
	json.Unmarshal([]byte(s), &y)
	fmt.Println(y.NameI18n)
}
