"""
使用python分析pprof文件

看来还是因为格式不同导致python无法分析golang
"""
import pstats

f = pstats.Stats("cpu.prof")
f.print_title()
