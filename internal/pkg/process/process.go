package process

import (
	"bytes"
	"fmt"
	"github.com/bgst009/ubiquitous-invention/internal/pkg/common"
	"github.com/bgst009/ubiquitous-invention/internal/pkg/info"
	"os/exec"
	"strconv"
	"strings"
)

// GetProcessPIDByCmd 根据输入的命令 cmd1 批量获取进程 pid 和进程 执行路径
func GetProcessPIDByCmd(cmd1 string) []info.Info {
	// 获取 PID 和 路径
	ProcessInfo := make([]info.Info, 10)
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

	return ProcessInfo
}

func GetProcessNameByPID(PID int) string {
	psCmd := common.GetCmdByPID(PID, "name")
	psBytes := common.ExexCmd(psCmd)
	psName := strings.ReplaceAll(bytes.NewBuffer(psBytes).String(), "\n", "")
	return psName
}

// GetProcessPIDByName 根据进程名，通过执行 shell 命令获取进程 pid
func GetProcessPIDByName(name string) int {
	cmd := ` ps -ef | grep ` + name + ` | grep -v "grep" |  awk '{pid=NF-6} {print $pid}'`
	cmdBytes := common.ExexCmd(cmd)

	pid, err := strconv.Atoi(strings.ReplaceAll(bytes.NewBuffer(cmdBytes).String(), "\n", ""))
	if err != nil {
		return 0
	}
	return pid
}

// GetProcessPathByName 根据进程名，通过执行 shell 命令获取进程 执行路径
func GetProcessPathByName(name string) string {
	cmd := ` ps -ef | grep ` + name + ` | grep -v "grep" |  awk '{path=NF-0} {print $path}'`
	cmdBytes := common.ExexCmd(cmd)
	return strings.ReplaceAll(bytes.NewBuffer(cmdBytes).String(), "\n", "")
}
