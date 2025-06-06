package unicorn

import (
	"net/http"
	"strconv"
	"unicorn-app/utils"

	"github.com/google/uuid"
)

func init() {
	requestAmountMap = make(map[string]int)
	requestUnicornMap = make(map[string][]utils.Unicorn)
}

func AddUnicornRequest(writer http.ResponseWriter, request *http.Request) {

	if request.Method != http.MethodGet {
		utils.ResponseJSON(writer, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	amountQuery := request.URL.Query().Get("amount")
	if amountQuery == "" {
		utils.ResponseJSON(writer, http.StatusBadRequest, "'amount' is required.")
		return
	}

	amount, err := strconv.Atoi(amountQuery)
	if err != nil {
		utils.ResponseJSON(writer, http.StatusBadRequest, "'amount' must be a valid integer")
		return
	}

	requestID := uuid.New().String()

	mutex.Lock()

	requestIdQueue = append(requestIdQueue, requestID)
	requestAmountMap[requestID] = amount

	mutex.Unlock()

	utils.ResponseJSON(
		writer,
		http.StatusOK,
		map[string]any{
			"request_id": requestID,
		},
	)
}

func UnicornPoll(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		utils.ResponseJSON(writer, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	requestID := request.URL.Query().Get("request_id")
	if requestID == "" {
		utils.ResponseJSON(writer, http.StatusBadRequest, "'request_id' is required")
		return
	}

	amount, ok := requestAmountMap[requestID]
	if !ok {
		utils.ResponseJSON(writer, http.StatusNotFound, "'request_id' does not exist")
		return
	}

	unicorns := requestUnicornMap[requestID]
	if len(unicorns) < amount {
		utils.ResponseJSON(writer, http.StatusAccepted, map[string]string{"message": "Unicorn is being produced. Please poll again."})
		return
	}

	utils.ResponseJSON(writer, http.StatusOK, unicorns)
}
