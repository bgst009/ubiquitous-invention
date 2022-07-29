package monitor

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"

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
	cmd1 := ` ps -ef | grep eb5gc/bin/ | grep -v "grep" |  awk '{pid=NF-6}{name=NF-0} {print $pid} {print $name}'`
	c := exec.Command("bash","-c",cmd1)
	fmt.Println(c.Output())
}

func init() {
	factory.InitConfigFactory("/home/yin/workspace/golang-code/ubiquitous-invention/config/cfg.yaml")
	cfg = factory.MonitorCfg
}
