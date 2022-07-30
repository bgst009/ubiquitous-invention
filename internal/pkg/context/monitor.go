package context

import (
	"github.com/bgst009/ubiquitous-invention/internal/pkg/info"
	"time"
)

type Monitor struct {
	processNames []string
	info         map[string]*info.Info
	CpuUsages    map[string]map[string]CpuUsage
	MemUsage     map[string]map[string]MemUsage
}

type CpuUsage struct {
	TempTime   time.Time
	Percentage string
}
type MemUsage struct {
	TempTime   time.Time
	Percentage string
}

func NewMonitor() *Monitor {
	m := new(Monitor)

	m.processNames = make([]string, 10)
	m.info = make(map[string]*info.Info, 10)
	m.CpuUsages = make(map[string]map[string]CpuUsage, 10)
	m.MemUsage = make(map[string]map[string]MemUsage, 10)

	return m
}

func (m *Monitor) SetProcessNames(names []string) {
	m.processNames = names
}

func (m *Monitor) SetInfoByName(name string, i *info.Info) {
	m.info[name] = i
}

func (m *Monitor) GetProcessNames() []string {
	return m.processNames
}

func (m *Monitor) GetInfoByName(name string) *info.Info {
	return m.info[name]
}
