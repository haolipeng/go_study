package main

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

//这里还仅仅是针对文件夹的
func main() {
	//调用os.Stat获取FileInfo结构体
	var err error
	path := "/var/spool/cron"
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.Stat failed!")
		return
	}

	stat, ok := info.Sys().(*syscall.Stat_t)
	if ok {
		uid := int(stat.Uid)
		gid := int(stat.Gid)
		u := strconv.FormatUint(uint64(uid), 10)
		g := strconv.FormatUint(uint64(gid), 10)
		usr, err := user.LookupId(u)
		if err != nil {
			fmt.Println("user.LookupId failed")
			return
		}
		group, err := user.LookupGroupId(g)
		if err != nil {
			fmt.Println("user.LookupGroupId failed")
			return
		}
		fmt.Printf("user:%s group:%v\n", usr.Name, group.Name)
	} else {
		fmt.Println("info.Sys().(*syscall.Stat_t) failed")
	}
}
