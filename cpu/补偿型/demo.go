package main

//
//import (
//	"log"
//	"math/rand"
//	"net/http"
//	_ "net/http/pprof"
//	"os"
//	"sync"
//	"time"
//)
//
//const (
//	base = 128 * 1024
//	max  = 8 * 1024 * 1024
//	min  = 4 * 1024 * 1024
//
//	workers = 8
//)
//
//func main() {
//	// 开启pprof
//	go func() {
//		ip := "0.0.0.0:6060"
//		if err := http.ListenAndServe(ip, nil); err != nil {
//			log.Printf("start pprof failed on %s\n", ip)
//			os.Exit(1)
//		}
//	}()
//
//	for i := 0; i < workers; i++ {
//		go func(a int) {
//			var ticker = time.NewTicker(10 * time.Millisecond)
//			defer ticker.Stop()
//			for range ticker.C {
//				if (int32(a)+rand.Int31())&1 == 1 {
//					send()
//				}
//				if (int32(a)+rand.Int31())&1 == 0 {
//					receive()
//				}
//			}
//		}(i)
//	}
//
//	select {}
//}
//
//// 重试集
//var retrySet sync.Map
//
//func receive() {
//	var i int
//	retrySet.Range(func(key, value interface{}) bool {
//		log.Println("key: ", key)
//		retrySet.Delete(key)
//		i++
//		if i == base {
//			return false
//		}
//		return true
//	})
//}
//
//func send() {
//	for i := 0; i < base; i++ {
//		var v = rand.Intn(max) - rand.Intn(min)
//		retrySet.Store(i*v, struct{}{})
//	}
//}
