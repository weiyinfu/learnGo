package main

import (
	s "strings"
	"fmt"
)
var p=fmt.Println
func main()  {
	p(s.Contains("weidiao","idi"))
	p(s.Compare("weidiao","haha"))
	p(s.Count("weidiao","i"))
	p(s.HasPrefix("weidiao","wei"))
}
