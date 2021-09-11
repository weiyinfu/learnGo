package main

import (
	"fmt"
	"regexp"
)

func main() {
	indicatorReg, err := regexp.Compile("\\$([^\\{'\"\\s]*)\\s*(\\{\\s*(.*?)\\}\\s*)?")
	if err != nil {
		panic("baga yalu")
	}
	for _, s := range []string{
		"anomaly_detect ('$asc_72_upstream_success_header{method=GetItem,to_cluster=online}')",
		"anomaly_detect (  '$sc_72_upstream_success_header{method=GetItem,to_cluster=online}')",
		"anomaly_detect (  '$sc_72_upstream_success_header{method=GetItem,to_cluster=online}'  )",
		"anomaly_detect('$sc_462_system_mem_usage_top10{cluster=default,host=10.112.69.161,pod_name=fed-dp-615be15483-67d9ddd894-cwvc9,psm=dp.nevermore.billboard}')",
		"$sc_462_system_mem_usage_top10{cluster=default,host=10.112.69.161,pod_name=fed-dp-615be15483-67d9ddd894-cwvc9,psm=dp.nevermore.billboard} +  $sc_462_system_mem_usage_top10{cluster=default,host=10.112.69.161,pod_name=fed-dp-615be15483-67d9ddd894-cwvc9,psm=dp.nevermore.billboard}",
	} {
		fmt.Println(indicatorReg.FindAllStringSubmatch(s, -1))
	}
}
