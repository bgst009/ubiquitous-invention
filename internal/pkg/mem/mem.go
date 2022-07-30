package mem

import (
	"bytes"
	"fmt"
	"github.com/bgst009/ubiquitous-invention/internal/pkg/common"
	"strconv"
	"strings"
)

func GetUsageByPID(pid int) string {
	memCmd := common.GetCmdByPID(pid, "mem")
	memBytes := common.ExexCmd(memCmd)

	memUsage := strings.ReplaceAll(bytes.NewBuffer(memBytes).String(), "\n", "")
	memUsagef, _ := strconv.ParseFloat(memUsage, 64)
	memUsage = fmt.Sprintf("%f%s", memUsagef/1024.00, ` M`)

	return memUsage
}
