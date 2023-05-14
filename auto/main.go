package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

func maxCoroutines(core int, cpuUsage, memUsage float64, t, s int) int {
	// 获取系统 CPU 核心数
	cpuNums, _ := cpu.Counts(true)

	// 计算可用的总 CPU 时间和总内存大小
	totalTime := float64(cpuNums)
	totalMem, _ := mem.VirtualMemory()
	totalMemSize := float64(totalMem.Total) * (1 - memUsage)

	// 计算一个协程所需的 CPU 时间和内存使用量
	cpuPerCoroutine := float64(t) / float64(core)
	memPerCoroutine := float64(s)

	// 根据公式计算最大协程数量
	n1 := int(totalTime / cpuPerCoroutine)          // 限制协程数的CPU时间
	n2 := int(totalMemSize / memPerCoroutine)       // 限制协程数的内存
	n3 := int(float64(core)*(1-memUsage)) / s       // 同时存在的协程数
	n4 := int(cpuUsage / (cpuPerCoroutine * 100.0)) // 限制协程数的CPU利用率
	fmt.Println(n1, n2, n3, n4)
	return min(n1, n2, n3, n4)

}

func min(nums ...int) int {
	res := nums[0]
	for _, num := range nums[1:] {
		if num < res {
			res = num
		}
	}
	return res
}

func main() {
	core := 16       // CPU核心数
	cpuUsage := 80.0 // CPU利用率
	memUsage := 0.8  // 内存利用率
	t := 1           // 单个协程执行的秒数
	s := 1024 * 1024 // 单个协程占用的内存大小

	maxCoroutines := maxCoroutines(core, cpuUsage, memUsage, t, s)
	fmt.Printf("在 CPU 核心数为 %d，CPU 利用率为 %.2f%%，内存利用率为 %.2f 的情况下，最大可以启动 %d 个协程。\n", core, cpuUsage, memUsage, maxCoroutines)
}
