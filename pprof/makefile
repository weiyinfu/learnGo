gen:
	# error:cannot use go run
	#go run --cpuprofile=cpu.prof prof_file.go
	# -memprofile 也可以了解一下
	go test -bench=. -cpuprofile=cpu.prof prof_file_test.go
see:
	echo http://localhost:8082
	go tool pprof -http=:8082 cpu.prof
prepare:
	go get -u github.com/google/pprof
pprof:
	# pprof是谷歌提供的基于语言自带pprof的工具，它的功能与自带的pprof大差不差
	pprof -http=:8082 cpu.prof
gen-file:
	go run code_pprof.go