package common

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
)

const Str1 = `top -bn 1 -p `
const StrMem = `| tail -1 | awk '{ssd=NF-6} {print $ssd }'`
const StrCpu = `| tail -1 | awk '{ssd=NF-3} {print $ssd }'`
const StrName = `| tail -1 | awk '{ssd=NF-0} {print $ssd }'`

func GetCmdByPID(pid int, w string) string {
	strPid := strconv.Itoa(pid)
	var cmdBuf bytes.Buffer
	cmdBuf.WriteString(Str1)
	cmdBuf.WriteString(strPid)

	switch w {
	case "cpu":
		cmdBuf.WriteString(StrCpu)
	case "mem":
		cmdBuf.WriteString(StrMem)
	case "name":
		cmdBuf.WriteString(StrName)
	default:
		cmdBuf.WriteString("")
	}

	cpuCmd := cmdBuf.String()
	return cpuCmd
}
func ExexCmd(cmd string) []byte {
	cmdByte, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to execute command: %s", cmd)
	}
	return cmdByte
}
