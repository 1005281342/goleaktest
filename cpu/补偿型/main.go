package main

import (
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"sync"
	"time"
)

func main() {
	// 开启pprof
	go func() {
		ip := "0.0.0.0:6060"
		if err := http.ListenAndServe(ip, nil); err != nil {
			log.Printf("start pprof failed on %s\n", ip)
			os.Exit(1)
		}
	}()

	go syncDone()
	go genTask()
	Compensator()
}

const Max = 1024 * 64

var set sync.Map

type tTask struct {
	FID int32
	FX  int32
	FA1 string
	FA2 string
	FA3 string
	FA4 string
	FA5 string
	FA6 string
	FA7 string
	FA8 string
	FT1 int64
	FT2 int64
}

func init() {
	var (
		k         = rand.Intn(32)
		tTaskTemp = make([]tTask, k)
	)
	for i := 0; i < k; i++ {
		tTaskTemp[i].FID = int32(i)
		tTaskTemp[i].FX = int32(i)
		set.Store(i, tTaskTemp[i])
	}
}

// Compensator
// 在线上遇到一个类似补偿不当导致CPU飙高的场景
func Compensator() {
	for {
		time.Sleep(time.Second)
		var tasks = getTasks()
		log.Printf("task count: %d", len(tasks))
		if len(tasks) <= 0 {
			continue
		}
		for _, task := range tasks {
			go handlerTask(task)
		}
	}
}

func handlerTask(t tTask) {
	t.FT1 = time.Now().Unix()
	var d = 16 + rand.Intn(64)
	time.Sleep(time.Duration(d) * time.Millisecond)
	for i := 0; i < d*1024; i++ {
	}
	log.Printf("handle task: %+v", t)
	var wg sync.WaitGroup
	wg.Add(1)
	go doTask(t, &wg)
	wg.Wait()
}

func doSomething() {
	for i := 0; i < 1024*(128+rand.Intn(16)); i++ {
	}
}

func doTask(t tTask, wg *sync.WaitGroup) {
	defer wg.Done()
	t.FT2 = time.Now().Unix()

	doSomething()

	var _, has = doneSet.Load(int(t.FID))
	if has {
		return
	}

	if rand.Intn(61) == 1 {
		for i := 0; i < 1024*rand.Intn(1024); i++ {
		}
	}
	log.Printf("task: %+v", t)
	doneSet.Store(int(t.FID), struct{}{})
}

// 模拟获取任务集合
func getTasks() []tTask {
	var tasks = make([]tTask, 0)
	set.Range(func(key, value interface{}) bool {
		var task, ok = value.(tTask)
		if !ok {
			log.Printf("conv failed %+v", value)
			return false
		}
		tasks = append(tasks, task)
		for i := 0; i < 1024*rand.Intn(16); i++ {
		}
		return true
	})
	return tasks
}

var doneSet sync.Map

func syncDone() {
	var tick = time.NewTicker(10 * time.Millisecond)
	for range tick.C {
		doneSet.Range(func(key, _ interface{}) bool {
			set.Delete(key)
			doneSet.Delete(key)
			return true
		})
	}
}

func genTask() {
	var tick = time.NewTicker(11 * time.Millisecond)
	for range tick.C {
		var (
			cnt       = rand.Intn(128)
			tTaskTemp = make([]tTask, cnt)
		)
		for i := 0; i < cnt; i++ {
			var x = rand.Int31() & (Max - 1)
			tTaskTemp[i].FID = x
			tTaskTemp[i].FX = x
			set.Store(i, tTaskTemp[i])
		}
	}
}
