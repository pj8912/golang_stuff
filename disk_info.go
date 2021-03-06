package main

import (
	"fmt"
	"syscall"
)


type DiskStatus struct{
	All uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}


func DiskUsage(path string) (disk DiskStatus){
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil{
		return
	}

	disk.All = fs.Blocks * uint64(fs.Bsize)
	disk.Free = fs.Bfree * uint64(fs.Bsize)
	disk.Used = disk.All - disk.Free
	return
}

const(
	B = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)



func main(){
	disk := DiskUsage("/")
	
	// %.2f is a placeholder for floating point number

	fmt.Printf("Total Space: %.2f GB\n", float64(disk.All)/float64(GB))
	fmt.Printf("Used Space: %.2f GB\n", float64(disk.Used)/float64(GB))
	fmt.Printf("Free Space %.2f GB\n", float64(disk.Free)/float64(GB))
}

