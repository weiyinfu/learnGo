package main

import (
	"fmt"
	"regexp"
)

func main() {
	//trigger := `anomaly_detect('$sc_462_system_mem_usage_top10{cluster=default,host=10.112.69.161,pod_name=fed-dp-615be15483-67d9ddd894-cwvc9,psm=dp.nevermore.billboard}')`
	//trigger := `anomaly_detect('$sc_462_system_mem_usage_top10{cluster=default,host=asdfdsf,pod_name=asdfasdf,psm=xcfvvcv}')`
	//indicatorReg := regexp.MustCompile(`.?(\$\w+)({(\w+=\w+,?)+})?`)
	//matches := indicatorReg.FindAllStringIndex(trigger, -1)
	//renameReg := regexp.MustCompile(`[${}=,()]`)
	//for _, idx := range matches {
	//	match := trigger[idx[0]:idx[1]]
	//	index := strings.Index(match, "$")
	//	if index < 0 {
	//		continue
	//	}
	//	match = match[index+1:]
	//	rename := renameReg.ReplaceAllString(match, "_")
	//	fmt.Println("match= ", match, "rename= ", rename)
	//	trigger = trigger[:index+1+idx[0]] + rename + trigger[idx[1]:]
	//	arr := strings.Split(strings.Trim(match, "}"), "{")
	//	if len(arr) == 0 {
	//		fmt.Println("cannot find anything")
	//	}
	//	indicatorName := arr[0]
	//	fmt.Println(indicatorName)
	//	fmt.Println(len(arr))
	//}
	res, err := regexp.Match("sc_\\d+_custom", []byte("sc_234_custom_123432"))
	fmt.Println(err, res)
}
