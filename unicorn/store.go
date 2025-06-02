package unicorn

import "unicorn-app/utils"

var (
	requestIdQueue        []string
	producedUnicornsStore []utils.Unicorn
	requestAmountMap      map[string]int
	requestUnicornMap     map[string][]utils.Unicorn
)
