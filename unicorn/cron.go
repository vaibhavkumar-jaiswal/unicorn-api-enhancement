package unicorn

import (
	"sync"
	"time"
	"unicorn-app/utils"
)

var mutex sync.Mutex

func UnicornProducer() {
	ticker := time.NewTicker(utils.PRODUCER_TIMER_SECONDS * time.Second)

	for range ticker.C {
		unicorn := utils.GenerateUnicorn()

		mutex.Lock()

		if len(requestIdQueue) > 0 {
			requestID := requestIdQueue[0]
			amount := requestAmountMap[requestID]
			requestUnicornMap[requestID] = append(requestUnicornMap[requestID], unicorn)

			if amount > 1 {
				producedUnicornCount := len(producedUnicornsStore)
				if producedUnicornCount >= amount-1 {
					diff := producedUnicornCount - amount + 1
					for i := producedUnicornCount - 1; i >= diff; i-- {
						requestUnicornMap[requestID] = append(requestUnicornMap[requestID], producedUnicornsStore[i])
					}
					producedUnicornsStore = producedUnicornsStore[:diff]
				}
			}

			if amount == 1 || len(requestUnicornMap[requestID]) == amount {
				requestIdQueue = requestIdQueue[1:]
			}
		} else {
			producedUnicornsStore = append(producedUnicornsStore, unicorn)
		}

		mutex.Unlock()
	}
}
