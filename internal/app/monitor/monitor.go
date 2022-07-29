package monitor

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/bgst009/ubiquitous-invention/internal/pkg/config"
	"github.com/bgst009/ubiquitous-invention/internal/pkg/cpu"
	"github.com/bgst009/ubiquitous-invention/internal/pkg/factory"
	"github.com/bgst009/ubiquitous-invention/internal/pkg/info"
	"github.com/bgst009/ubiquitous-invention/internal/pkg/mem"
	"github.com/bgst009/ubiquitous-invention/internal/pkg/process"
	os_mem "github.com/shirou/gopsutil/v3/mem"
)

var (
	cfg         *config.Config
	ProcessInfo = make([]info.Info, 10)
)

func NewMonitor() {
	fmt.Println("monitor started")

	// 获取所有进程 PID
	for i := 0; i < len(cfg.Processors); i++ {
		p, _ := process.GetProcessPIDByName(cfg.Processors[i])
		ProcessInfo[i].PID = p
	}

	for _, pinfo := range ProcessInfo {
		// 根据 PID 获取所有进程 memory usage
		m, _ := mem.GetUsageByPID(pinfo.PID)
		pinfo.MemoryUsage = m
		// 根据 PID 获取所有进程 cpu usage
		c, _ := cpu.GetUsageByPID(pinfo.PID)
		pinfo.CpuUsage = c
	}

	f, err := os.OpenFile(cfg.OutPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Fatalf("error while opening the file. %v", err)
	}
	defer f.Close()
	b, _ := json.MarshalIndent(ProcessInfo, "", "\t")
	f.Write(b)

	v, _ := os_mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)

}

func Monitor5gc() {
	// 获取 PID 和 路径
	cmd1 := ` ps -ef | grep eb5gc/bin/ | grep -v "grep" |  awk '{pid=NF-6}{name=NF-0} {print $pid} {print $name}'`
	fmt.Printf("cmd1: %v\n", cmd1)
	b, err := exec.Command("bash", "-c", cmd1).Output()
	if err != nil {
		fmt.Printf("Failed to execute command: %s", cmd1)
	}
	// fmt.Println(string(b))

	sa := strings.Split(fmt.Sprint(string(b)), "\n")

	j := 0
	for i := 0; i < len(sa)-1; i += 2 {
		pid, err := strconv.Atoi(sa[i])
		if err != nil {
			fmt.Printf("err: %v\n", err)
			break
		}
		ProcessInfo[j].PID = pid
		ProcessInfo[j].ProcessPath = sa[i+1]
		j++
	}

	fmt.Printf("len(sa): %v\n", len(sa))

	fmt.Printf("ProcessInfo: %v\n", ProcessInfo)

	f, err := os.OpenFile("out.json", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Fatalf("error while opening the file. %v", err)
	}
	defer f.Close()
	bt, _ := json.MarshalIndent(ProcessInfo, "", "\t")
	f.Write(bt)

}

func init() {
	factory.InitConfigFactory("/home/yin/workspace/golang-code/ubiquitous-invention/config/cfg.yaml")
	cfg = factory.MonitorCfg
}
