package unicorn

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"
	"unicorn/utils"
)

var mutex sync.Mutex

func UnicornProducer() {
	LoadData()
	ticker := time.NewTicker(utils.PRODUCER_TIMER_SECONDS * time.Second)
	for range ticker.C {
		unicorn := generateUnicorn()

		mutex.Lock()
		fmt.Println("========================BEFORE============================")
		PrettyPrint("unicorn", unicorn)
		PrettyPrint("requestIdQueue", requestIdQueue)
		PrettyPrint("requestAmountMap", requestAmountMap)
		PrettyPrint("requestUnicornMap", requestUnicornMap)
		PrettyPrint("producedUnicornsStore", producedUnicornsStore)

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
		fmt.Println("========================AFTER============================")
		PrettyPrint("requestIdQueue", requestIdQueue)
		PrettyPrint("requestAmountMap", requestAmountMap)
		PrettyPrint("requestUnicornMap", requestUnicornMap)
		PrettyPrint("producedUnicornsStore", producedUnicornsStore)
		mutex.Unlock()
	}
}

func generateUnicorn() Unicorn {
	fmt.Println("Producing unicorn...!")

	unicornCapabilities := []string{}
	capabilitiesMap := make(map[string]bool)
	name := fmt.Sprintf("%s-%s", adjectives[rand.Intn(len(adjectives))], names[rand.Intn(len(names))])

	for len(unicornCapabilities) < 3 {
		capability := utils.Capabilities[rand.Intn(len(utils.Capabilities))]
		if !capabilitiesMap[capability] {
			capabilitiesMap[capability] = true
			unicornCapabilities = append(unicornCapabilities, capability)
		}
	}
	return Unicorn{Name: name, Capabilities: unicornCapabilities}
}

func LoadData() {
	fmt.Println("Loading data...!")
	var err error
	names, err = utils.ReadFileData("petnames.txt")
	if err != nil {
		fmt.Println("please try later, unicorn factory unavailable")
		return
	}
	adjectives, err = utils.ReadFileData("adj.txt")
	if err != nil {
		fmt.Println("please try later, unicorn factory unavailable")
		return
	}
}

func PrettyPrint(name string, v interface{}) {
	bytes, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(name, " ----> ", string(bytes))
}
