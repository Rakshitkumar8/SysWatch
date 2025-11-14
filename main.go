package main

import (
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "time"
)

var (
    cpuGauge = prometheus.NewGauge(prometheus.GaugeOpts{
        Name: "system_cpu_usage_percent",
        Help: "CPU usage percentage",
    })

    memoryGauge = prometheus.NewGauge(prometheus.GaugeOpts{
        Name: "system_memory_usage_percent",
        Help: "Memory usage percentage",
    })

    diskGauge = prometheus.NewGauge(prometheus.GaugeOpts{
        Name: "system_disk_usage_percent",
        Help: "Disk usage percentage",
    })

    networkSentGauge = prometheus.NewGauge(prometheus.GaugeOpts{
        Name: "system_network_sent_bytes",
        Help: "Total network bytes sent",
    })

    networkRecvGauge = prometheus.NewGauge(prometheus.GaugeOpts{
        Name: "system_network_received_bytes",
        Help: "Total network bytes received",
    })
)

func recordMetrics(logger *zap.Logger) {
    go func() {
        for {
            cpu, _ := GetCPUUsage()
            mem, _ := GetMemoryUsage()
            disk, _ := GetDiskUsage()
            sent, recv, _ := GetNetworkUsage()

            cpuGauge.Set(cpu)
            memoryGauge.Set(mem)
            diskGauge.Set(disk)
            networkSentGauge.Set(float64(sent))
            networkRecvGauge.Set(float64(recv))

            logger.Info("Metrics updated",
                zap.Float64("cpu", cpu),
                zap.Float64("memory", mem),
                zap.Float64("disk", disk),
            )

            time.Sleep(5 * time.Second)
        }
    }()
}

func main() {
    logger, _ := zap.NewProduction()
    defer logger.Sync()

    // Register Prometheus metrics
    prometheus.MustRegister(cpuGauge, memoryGauge, diskGauge, networkSentGauge, networkRecvGauge)

    recordMetrics(logger)

    router := gin.Default()

    router.GET("/metrics", gin.WrapH(promhttp.Handler()))
    router.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    logger.Info("Starting SysWatch monitoring service")
    router.Run(":8080")
}
