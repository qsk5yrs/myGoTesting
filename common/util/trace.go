package util

import (
	"context"
	"encoding/binary"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

// GenerateSpanID 通过结合时间戳、IP 地址和随机数生成一个唯一的SpanID，用于在分布式系统中唯一标识一个请求或操作。
// 这种方法可以确保生成的 SpanID 具有较高的唯一性和可区分性
func GenerateSpanID(addr string) string {
	strAddr := strings.Split(addr, ":")
	ip := strAddr[0]
	ipLong, _ := Ip2Long(ip)
	times := uint64(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	// go1.20后生成随机数方法
	//rand.New(rand.NewSource(time.Now().UnixNano()))
	spanId := ((times ^ uint64(ipLong)) << 32) | uint64(rand.Int31())
	return strconv.FormatUint(spanId, 16)
}

// Ip2Long 将一个 IPv4 地址字符串转换成一个无符号的 32 位整数
func Ip2Long(ip string) (uint32, error) {
	ipAddr, err := net.ResolveIPAddr("ip", ip)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(ipAddr.IP.To4()), nil
}

func GetTraceInfoFromCtx(ctx context.Context) (traceId, spanId, pSpanId string) {
	if ctx.Value("traceid") != nil {
		traceId = ctx.Value("traceid").(string)
	}
	if ctx.Value("spanid") != nil {
		spanId = ctx.Value("spanid").(string)
	}
	if ctx.Value("pspanid") != nil {
		pSpanId = ctx.Value("pspanid").(string)
	}
	return
}
