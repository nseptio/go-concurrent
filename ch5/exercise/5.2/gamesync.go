package main

import (
	"fmt"
	"sync"
	"time"
)

const expiredSecond = 5

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	playersInGame := 10
	isCancel := false
	go timeout(cond, &isCancel)
	for playerID := 0; playerID < playersInGame; playerID++ {
		go playerHandler(cond, &playersInGame, playerID, &isCancel)
		time.Sleep(2 * time.Second)
	}
	time.Sleep(5 * time.Second)
}

func timeout(cond *sync.Cond, isCancel *bool) {
	time.Sleep(expiredSecond * time.Second)
	cond.L.Lock()
	*isCancel = true
	cond.Broadcast()
	cond.L.Unlock()
}

func playerHandler(cond *sync.Cond, playersRemaining *int, playerID int, isCancel *bool) {
	cond.L.Lock()
	fmt.Println(playerID, ": Connected")
	*playersRemaining--

	if *playersRemaining == 0 {
		cond.Broadcast()
	}

	for *playersRemaining > 0 && !*isCancel {
		fmt.Println(playerID, ": Waiting for more players")
		cond.Wait()
	}

	cond.L.Unlock()

	if *isCancel {
		fmt.Println(playerID, ": Game cancelled")
	} else {
		fmt.Println("All players connected. Ready player", playerID)
	}
}
