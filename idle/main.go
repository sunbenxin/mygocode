package main

import (
	"flag"
	"log"
	"sync"

	"backend/config"
	"backend/domain/inner"
	"backend/gateway/push/service"
	"backend/slog"
	"backend/version"
)

var (
	configPath    = flag.String("config", "/etc/putong/config.json", "Path to configuration file.")
	fromShard     = flag.Int("fromShard", 0, "Fetch users starting from this logical shard")
	toShard       = flag.Int("toShard", 0, "Fetch users starting from this logical shard")
	pushPerSecond = flag.Int("pushPerSecond", 0, "push num")
	n             = flag.Int("n", 1, "concurrency")

	noop = flag.Bool("noop", false, "Whether push or not, default false")

	// pushClient is used to push reminder event to push service.
	pushClient *service.Client
)

func init() {
	version.Init()
	flag.Parse()
	if err := config.ParseGlobal(*configPath); err != nil {
		log.Fatal(err)
	}
	if err := slog.Init(config.Conf.Log); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// init pushClient
	pushClient = service.NewClient()

	var wg sync.WaitGroup
	for i := 0; i < *n; i++ {
		wg.Add(1)

		go func(pushPerSecond int) {
			defer wg.Done()
			for i := 0; i < pushPerSecond; i++ {
				reminder := inner.Reminder{
					UserId:           "2695593",
					HasUnswipedLikes: true,
					Type:             inner.ReminderType,
				}

				log.Println("PUSH ONE")
				err := pushClient.PushReminder(&reminder)
				if err != nil {
					log.Println(err)
					continue
				}
				log.Println("PushReminder %+v", reminder)
			}
		}(*pushPerSecond)
	}

	wg.Wait()
}
