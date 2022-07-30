package monitor

import (
	"bytes"
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

	str1 := `top -bn 1 -p `
	str2 := `| tail -1 | awk '{ssd=NF-6} {print $ssd }'`
	str3 := `| tail -1 | awk '{ssd=NF-3} {print $ssd }'`

	for i := 0; i < len(ProcessInfo); i++ {
		stringPid := strconv.Itoa(ProcessInfo[i].PID)
		var memCmdbuf bytes.Buffer
		memCmdbuf.WriteString(str1)
		memCmdbuf.WriteString(stringPid)
		memCmdbuf.WriteString(str2)
		ProcessInfo[i].MemCmd = memCmdbuf.String()
		fmt.Printf("info.MemCmd: %v\n", ProcessInfo[i].MemCmd)
		var cpuCmdbuf bytes.Buffer
		cpuCmdbuf.WriteString(str1)
		cpuCmdbuf.WriteString(stringPid)
		cpuCmdbuf.WriteString(str3)
		ProcessInfo[i].CpuCmd = cpuCmdbuf.String()
		fmt.Printf("info.CpuCmd: %v\n", ProcessInfo[i].CpuCmd)
	}

	for i := 0; i < len(ProcessInfo); i++ {
		cpuByte, err := exec.Command("bash", "-c", ProcessInfo[i].CpuCmd).CombinedOutput()
		if err != nil {
			fmt.Printf("Failed to execute command: %s", ProcessInfo[i].CpuCmd)
		}
		ProcessInfo[i].CpuUsage = bytes.NewBuffer(cpuByte).String()[:1]

		memBytes, err := exec.Command("bash", "-c", ProcessInfo[i].MemCmd).CombinedOutput()
		if err != nil {
			fmt.Printf("Failed to execute command: %s", ProcessInfo[i].MemCmd)
		}
		ProcessInfo[i].MemoryUsage = bytes.NewBuffer(memBytes).String()[:1]
		fmt.Printf("cpu: %s\tmem: %s \n,", ProcessInfo[i].CpuUsage, ProcessInfo[i].MemoryUsage)

	}

	// 打印信息
	indent, err := json.MarshalIndent(ProcessInfo, "", "\t")
	if err != nil {
		return
	}
	fmt.Printf("%s\n", bytes.NewBuffer(indent).String())
	// 写入文件
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
