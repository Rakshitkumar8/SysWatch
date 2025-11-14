package main

import (
    "github.com/shirou/gopsutil/v3/cpu"
    "github.com/shirou/gopsutil/v3/mem"
    "github.com/shirou/gopsutil/v3/disk"
    "github.com/shirou/gopsutil/v3/net"
    "time"
)

// Get CPU usage %
func GetCPUUsage() (float64, error) {
    percent, err := cpu.Percent(time.Second, false)
    if err != nil {
        return 0, err
    }
    return percent[0], nil
}

// Get Memory usage %
func GetMemoryUsage() (float64, error) {
    vm, err := mem.VirtualMemory()
    if err != nil {
        return 0, err
    }
    return vm.UsedPercent, nil
}

// Get Disk usage %
func GetDiskUsage() (float64, error) {
    usage, err := disk.Usage("C:/")
    if err != nil {
        return 0, err
    }
    return usage.UsedPercent, nil
}

// Get Network bytes sent + received
func GetNetworkUsage() (uint64, uint64, error) {
    ioStats, err := net.IOCounters(false)
    if err != nil {
        return 0, 0, err
    }
    return ioStats[0].BytesSent, ioStats[0].BytesRecv, nil
}
