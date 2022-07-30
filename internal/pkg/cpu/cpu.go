package cpu

import (
	"bytes"
	"github.com/bgst009/ubiquitous-invention/internal/pkg/common"
	"strings"
)

func GetUsageByPID(pid int) string {
	cpuCmd := common.GetCmdByPID(pid, "cpu")
	cpuByte := common.ExexCmd(cpuCmd)
	cpuUsage := strings.ReplaceAll(bytes.NewBuffer(cpuByte).String(), "\n", " %")
	return cpuUsage
}
