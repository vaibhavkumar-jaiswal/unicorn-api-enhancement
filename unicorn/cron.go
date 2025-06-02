package unicorn

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"unicorn/utils"
)

var mutex sync.Mutex

func UnicornProducer() {
	if err := LoadData(); err != nil {
		fmt.Println(err.Error())
		return
	}

	ticker := time.NewTicker(utils.PRODUCER_TIMER_SECONDS * time.Second)
	for range ticker.C {
		unicorn := generateUnicorn()

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

func generateUnicorn() Unicorn {
	fmt.Println("Producing unicorn...!")

	unicornCapabilities := []string{}
	capabilitiesMap := make(map[string]bool)
	name := fmt.Sprintf("%s-%s", adjectives[rand.Intn(len(adjectives))], names[rand.Intn(len(names))])

	for len(unicornCapabilities) < utils.MAX_CAPABILITY {
		capability := utils.Capabilities[rand.Intn(len(utils.Capabilities))]
		if !capabilitiesMap[capability] {
			capabilitiesMap[capability] = true
			unicornCapabilities = append(unicornCapabilities, capability)
		}
	}
	return Unicorn{Name: name, Capabilities: unicornCapabilities}
}

func LoadData() error {
	nameData, err := utils.ReadFileData("petnames.txt")
	if err != nil {
		return fmt.Errorf("please try later, unicorn factory unavailable")
	}
	names = nameData

	adjData, err := utils.ReadFileData("adj.txt")
	if err != nil {
		return fmt.Errorf("please try later, unicorn factory unavailable")
	}
	adjectives = adjData

	return nil
}
