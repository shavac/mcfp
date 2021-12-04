package mcfp

import (
	"context"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

var (
	_ = cpu.Times
)

func init() {

}

func String() string {
	ctxTimeout, _ := context.WithTimeout(nil, 1*time.Second)
	infos, err := cpu.InfoWithContext(ctxTimeout)
	if err != nil {
		return ""
	}
	return fmt.Sprint(infos)
}
