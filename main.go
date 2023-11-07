package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func upstream() int {
	log.Print("upstream call")
	randNum := 1
	time.Sleep(time.Duration(randNum) * time.Second)
	return rand.Int()
}

type Cache struct {
	Data         int
	upstreamFunc func() int
	updatePeriod time.Duration
	lastUpdate   time.Time
	updaterChan  chan int
}

type CacheImpl interface {
	Get() (int, error)
}

func (c *Cache) Get() (int, error) {
	if c.Data != 0 && time.Since(c.lastUpdate) < c.updatePeriod {
		return c.Data, nil
	}

	ctx := context.Background()
	ctx, cancelFunc := context.WithTimeout(ctx, time.Second*2)
	defer cancelFunc()

	go func() {
		data := c.upstreamFunc()
		c.updaterChan <- data
	}()

	select {
	case result := <-c.updaterChan:
		c.Data = result
		c.lastUpdate = time.Now()
		return result, nil
	case <-ctx.Done():
		return 0, fmt.Errorf("service not available")
	}
}

func main() {
	c := &Cache{
		upstreamFunc: upstream,
		updatePeriod: time.Duration(time.Second * 2),
		updaterChan:  make(chan int),
	}

	log.Println()
	data, err := c.Get()
	log.Println(data)
	log.Println(err)
	log.Println()
	data, err = c.Get()
	log.Println(data)
	log.Println(err)

}
