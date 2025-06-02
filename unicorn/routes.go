package unicorn

import "net/http"

func Routes() {
	http.HandleFunc("/api/unicorn", AddUnicornRequest)
	http.HandleFunc("/api/unicorn/poll", UnicornPoll)
}
