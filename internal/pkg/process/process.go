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

func GetProcessPIDByCmd(cmd1 string) []info.Info {
	// 获取 PID 和 路径
	ProcessInfo := make([]info.Info, 10)
	//cmd1 := ` ps -ef | grep eb5gc/bin/ | grep -v "grep" |  awk '{pid=NF-6}{name=NF-0} {print $pid} {print $name}'`
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

	return ProcessInfo
}

func GetProcessNameByPID(PID int) string {
	psCmd := common.GetCmdByPID(PID, "name")
	psBytes := common.ExexCmd(psCmd)
	psName := strings.ReplaceAll(bytes.NewBuffer(psBytes).String(), "\n", "")
	return psName
}
