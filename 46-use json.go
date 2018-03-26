package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
)

type response1 struct {
	//只有大写的(公有类型)才能转换成json
	page   int
	Fruits []string
}
type response2 struct {
	Page  int      `json:"mypage"`
	Fruit []string `json:"FRUITS"`
}

func main() {
	boolBytes, err := json.Marshal(true)
	fmt.Println(string(boolBytes), err)
	intBytes, err := json.Marshal(123)
	fmt.Println(string(intBytes), err)
	floatBytes, err := json.Marshal(math.Pi)
	fmt.Println(string(floatBytes), err)

	arrayBytes, err := json.Marshal([]string{"one", "two", "three"})
	fmt.Println(string(arrayBytes))

	mapBytes, err := json.Marshal(map[string]int{"one": 1, "two": 2})
	fmt.Println(string(mapBytes))

	//不指定json,输出为空
	response1Bytes := response1{1, []string{"one", "two"}}
	structBytes, err := json.Marshal(response1Bytes)
	fmt.Println(string(structBytes))

	//调用函数的时候传入指针效率更高,避免了Struct复制到栈里面
	structBytes, err = json.Marshal(&response2{1, []string{"one", "two"}})
	fmt.Println(string(structBytes), err)

	/**
	* 下面是反序列化的过程
	*/
	byt := []byte(`{"mypage":2,"FRUITS":["apple","banana"]}`)
	//声明一个变量
	var stringObjMap map[string]interface{}
	//这个地方没有stringObjMap是不对的
	if err := json.Unmarshal(byt, &stringObjMap); err != nil {
		panic(err)
	}
	fmt.Println(stringObjMap)

	//使用强制类型转换
	num := stringObjMap["mypage"].(float64)
	fmt.Println(num)

	//将json字符串反序列化为结构体
	response2body := response2{}
	json.Unmarshal(byt, &response2body)
	fmt.Println(response2body)

	//一种非常奇葩的json打印方式
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
