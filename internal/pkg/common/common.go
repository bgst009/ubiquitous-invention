package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bgst009/ubiquitous-invention/internal/pkg/context"
	"github.com/bgst009/ubiquitous-invention/internal/pkg/info"
	"log"
	"os"
	"os/exec"
	"strconv"
)

const Str1 = `top -bn 1 -p `
const StrMem = `| tail -1 | awk '{ssd=NF-6} {print $ssd }'`
const StrCpu = `| tail -1 | awk '{ssd=NF-3} {print $ssd }'`
const StrName = `| tail -1 | awk '{ssd=NF-0} {print $ssd }'`

func WriteInfo(path string, ProcessInfo []info.Info) {
	// 写入文件
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Fatalf("error while opening the file. %v", err)
	}
	defer f.Close()
	bt, _ := json.MarshalIndent(ProcessInfo, "", "\t")
	f.Write(bt)
}

func WriteMonitor(path string, m context.Monitor) {
	// 写入文件
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Fatalf("error while opening the file. %v", err)
	}
	defer f.Close()
	bt, _ := json.MarshalIndent(m, "", "\t")
	f.Write(bt)
}

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
