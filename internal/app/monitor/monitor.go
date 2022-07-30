package monitor

import (
	"encoding/json"
	"fmt"
	"github.com/bgst009/ubiquitous-invention/internal/pkg/common"
	"github.com/bgst009/ubiquitous-invention/internal/pkg/config"
	"github.com/bgst009/ubiquitous-invention/internal/pkg/context"
	"github.com/bgst009/ubiquitous-invention/internal/pkg/cpu"
	"github.com/bgst009/ubiquitous-invention/internal/pkg/factory"
	"github.com/bgst009/ubiquitous-invention/internal/pkg/info"
	"github.com/bgst009/ubiquitous-invention/internal/pkg/mem"
	"github.com/bgst009/ubiquitous-invention/internal/pkg/process"
	osmem "github.com/shirou/gopsutil/v3/mem"
	"log"
	"os"
	"time"
)

var (
	cfg         *config.Config
	ProcessInfo = make([]info.Info, 10)
)

func NewMonitor() {
	fmt.Println("monitor started")

	// 获取所有进程 PID
	for i := 0; i < len(cfg.Processors); i++ {
		//p, _ := process.GetProcessPIDByName(strconv.Itoa(cfg.Processors[i]))
		//ProcessInfo[i].PID = p
	}

	for _, pinfo := range ProcessInfo {
		// 根据 PID 获取所有进程 memory usage
		m := mem.GetUsageByPID(pinfo.PID)
		pinfo.MemoryUsage = m
		// 根据 PID 获取所有进程 cpu usage
		c := cpu.GetUsageByPID(pinfo.PID)
		pinfo.CpuUsage = c
	}

	f, err := os.OpenFile(cfg.OutPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Fatalf("error while opening the file. %v", err)
	}
	defer f.Close()
	b, _ := json.MarshalIndent(ProcessInfo, "", "\t")
	f.Write(b)

	v, _ := osmem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)

}

func Monitor5gc() {
	// 获取 PID 和 路径
	cmd := ` ps -ef | grep eb5gc/bin/ | grep -v "grep" |  awk '{pid=NF-6}{name=NF-0} {print $pid} {print $name}'`
	ProcessInfo = process.GetProcessPIDByCmd(cmd)

	// 获取命令
	for i := 0; i < len(ProcessInfo); i++ {
		ProcessInfo[i].MemCmd = common.GetCmdByPID(ProcessInfo[i].PID, "mem")
		//fmt.Printf("info.MemCmd: %v\n", ProcessInfo[i].MemCmd)
		ProcessInfo[i].CpuCmd = common.GetCmdByPID(ProcessInfo[i].PID, "cpu")
		//fmt.Printf("info.CpuCmd: %v\n", ProcessInfo[i].CpuCmd)
		ProcessInfo[i].ProcessNameCmd = common.GetCmdByPID(ProcessInfo[i].PID, "name")
		//fmt.Printf("info.ProcessNameCmd: %v\n", ProcessInfo[i].ProcessNameCmd)
	}

	// 获取数据
	for i := 0; i < len(ProcessInfo); i++ {
		ProcessInfo[i].CpuUsage = cpu.GetUsageByPID(ProcessInfo[i].PID)
		ProcessInfo[i].MemoryUsage = mem.GetUsageByPID(ProcessInfo[i].PID)
		ProcessInfo[i].ProcessName = process.GetProcessNameByPID(ProcessInfo[i].PID)
		//fmt.Printf("ps: %s\tcpu: %s\tmem: %s\n,", ProcessInfo[i].ProcessName, ProcessInfo[i].CpuUsage, ProcessInfo[i].MemoryUsage)

	}

	t1 := time.NewTimer(time.Second * 5)
	t2 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			common.WriteInfo("out-5s.json", ProcessInfo)
		case <-t2.C:
			common.WriteInfo("out-10s.json", ProcessInfo)
		}
	}

}

func TickM() {
	monitor := context.NewMonitor()
	monitor.SetProcessNames(cfg.Processors)

	for i := 0; i < len(monitor.GetProcessNames()); i++ {
		name := monitor.GetProcessNames()[i]

		if monitor.CpuUsages[name] == nil {
			monitor.CpuUsages[name] = make(map[string]context.CpuUsage)
		}
		monitor.CpuUsages[name]["10"] = struct {
			TempTime   time.Time
			Percentage string
		}{TempTime: time.Now(), Percentage: cpu.GetUsageByPID(process.GetProcessPIDByName(name))}

		if monitor.MemUsage[name] == nil {
			monitor.MemUsage[name] = make(map[string]context.MemUsage)
		}
		monitor.MemUsage[name]["10"] = struct {
			TempTime   time.Time
			Percentage string
		}{TempTime: time.Now(), Percentage: mem.GetUsageByPID(process.GetProcessPIDByName(name))}

	}

	common.WriteMonitor("monitor-10.json", *monitor)
}

func init() {
	factory.InitConfigFactory()
	cfg = factory.MonitorCfg
}
