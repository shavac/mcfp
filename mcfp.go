package mcfp

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	cpu "github.com/shirou/gopsutil/cpu"
	disk "github.com/shirou/gopsutil/disk"
	net "github.com/shirou/gopsutil/net"
)

func GetCpuString() string {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	cinfos, err := cpu.InfoWithContext(ctxTimeout)
	if err != nil || len(cinfos) <= 0 {
		return ""
	}
	str := strings.Join([]string{cinfos[0].VendorID, cinfos[0].Model}, "-")
	return str
}

//hw addr of eth0 or en0, if none, first nic
func GetNicString() string {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	ninfos, err := net.InterfacesWithContext(ctxTimeout)
	if err != nil || len(ninfos) <= 0 {
		return ""
	}
	mac := ninfos[0].HardwareAddr
LOOPNIC:
	for _, ninfo := range ninfos {
		switch ninfo.Name {
		case "en0", "eth0":
			mac = ninfo.HardwareAddr
			break LOOPNIC
		}
	}
	return mac
}

func GetRootDevPath() string {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	dinfos, err := disk.PartitionsWithContext(ctxTimeout, true)
	if err != nil || len(dinfos) <= 0 {
		return ""
	}
	rdev := ""
	for _, dinfo := range dinfos {
		if dinfo.Mountpoint == "/" {
			rdev = dinfo.Device
			break
		}
	}
	return rdev
}

func GetFsUUID(rdev string) string {
	sb := make([]byte, 512)
	f, err := os.Open(rdev)
	if err != nil {
		log.Println(err)
		return ""
	}
	if _, err := f.ReadAt(sb, 1024); err != nil {
		log.Fatalln(err)
	}

	return fmt.Sprint(string(sb[108:124]))
}

func GetFingerPrint(mch Machiner) string {
	fsuuid := GetFsUUID(mch.RootDevPath())
	fp := fmt.Sprintf("%s||%s|%d|%s|%s", mch.Model(), mch.OS(), mch.NCPU(), mch.MAC(), fsuuid)
	return fp
}
