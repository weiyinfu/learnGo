golang语言自带性能分析工具，有两种使用方式：
* HTTP服务形式：参考web.go
* 以test的形式产生性能文件：参考prof_file_test.go
* 以code的形式产生性能文件：参考code_pprof.go文件
* 分析pprof文件，使用python分析go的prof文件是走不通的，使用golang可以分析go产生的prof文件。golang分析pprof文件在runtime/pprof/internal/profile/profile.go文件中。    