package main

import (
	"reflect"
	"time"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func test1()  {
	log.Println("test1")
}

func test2()  {
	log.Println("test2")
}

func test3(module string)  {
	log.Printf("test3 %v\n",module)
}

func main() {
	scheduler(test1,"00:00:00","5s")
	scheduler(test2,"00:00:00","10s")
	scheduler(test3,"00:00:00","15s","test scheduler params")
	//var c = make(chan bool)
	//c <- true
	//<-c
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done
}

// sched to start scheduler job at start time by interval duration.

//jobFunc : func(){}
//start :   2006-01-02 15:04:05	ParseInLocation is like Parse but differs in two important ways. First, in the absence of time zone information, Parse interprets a time as UTC; ParseInLocation interprets the time as in the given location. Second, when given a zone offset or abbreviation, Parse tries to match it against the Local location; ParseInLocation uses the given location.
//interval : ParseDuration parses a duration string. A duration string is a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as "300ms", "-1.5h" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
//jobArgs : func params
func scheduler(jobFunc interface{}, start, interval string, jobArgs ...interface{}) {
	jobValue := reflect.ValueOf(jobFunc)
	if jobValue.Kind() != reflect.Func {
		log.Panic("only function can be schedule.")
	}
	if len(jobArgs) != jobValue.Type().NumIn() {
		log.Panic("The number of args valid.")
	}
	// Get job function args.
	in := make([]reflect.Value, len(jobArgs))
	for i, arg := range jobArgs {
		in[i] = reflect.ValueOf(arg)
	}

	// Get interval d.
	d, err := time.ParseDuration(interval)
	if err != nil {
		log.Panic(err)
	}
	//location, err := time.LoadLocation("Asia/Shanghai")
	location, err := time.LoadLocation("Local")
	if err != nil {
		log.Panic(err)
	}
	t, err := time.ParseInLocation("15:04:05", start, location)
	if err != nil {
		log.Panic(err)
	}
	now := time.Now()

	// Start time.
	t = time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), t.Second(), 0, location)

	if now.After(t) {
		t = t.Add((now.Sub(t)/d + 1) * d)
	}

	time.Sleep(t.Sub(now))
	go jobValue.Call(in)
	ticker := time.NewTicker(d)
	go func() {
		for _ = range ticker.C {
			go jobValue.Call(in)
		}
	}()
}
