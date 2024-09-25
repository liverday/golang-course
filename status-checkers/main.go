package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/robfig/cron"
)

type LinkResponse struct {
	link     string
	status   string
	duration int64
}

func main() {
	f, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("An error ocurred: ", err)
		os.Exit(1)
	}

	defer f.Close()
	sc := bufio.NewScanner(f)

	links := []string{}
	for sc.Scan() {
		links = append(links, sc.Text())
	}

	cr := cron.New()
	ch := make(chan *LinkResponse)

	fmt.Println("Starting program: ")
	fmt.Println("---------------")
	cr.AddFunc("0 */1 * * *", func() {
		fmt.Println("Running cron job at: ", time.Now().Format(time.Kitchen))

		for _, link := range links {
			go checkLink(link, ch)
		}

		for i := 0; i < len(links); i++ {
			r := <-ch
			fmt.Printf("Link: %s, Status: %s, Duration: %dms\n", r.link, r.status, r.duration)
		}

		fmt.Println("-------------------")
	})

	cr.Start()

	q := make(chan os.Signal, 1)
	signal.Notify(q, syscall.SIGINT, syscall.SIGTERM)
	<-q

	fmt.Println("Stopping program")
}

func checkLink(link string, ch chan *LinkResponse) {
	start := time.Now()
	_, err := http.Get(link)

	if err != nil {
		end := time.Now()
		d := end.UnixMilli() - start.UnixMilli()
		ch <- &LinkResponse{link: link, status: "down", duration: d}
	}

	end := time.Now()
	d := end.UnixMilli() - start.UnixMilli()
	ch <- &LinkResponse{link: link, status: "up", duration: d}
}
