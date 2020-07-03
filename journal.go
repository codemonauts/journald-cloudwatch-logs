package main

import (
	"strconv"
	"strings"

	"github.com/coreos/go-systemd/sdjournal"
)

func addLogFilters(journal *sdjournal.Journal, config *config) {

	// Add Priority Filters
	if config.LogPriority < debugP {
		for p := range priorityJSON {
			if p <= config.LogPriority {
				journal.AddMatch("PRIORITY=" + strconv.Itoa(int(p)))
			}
		}
		journal.AddDisjunction()
	}

	// Add unit filter (multiple values possible, separate by ",")
	if config.LogUnit != "" {
		unitsRaw := strings.Split(config.LogUnit, ",")

		for _, unitRaw := range unitsRaw {
			unit := strings.TrimSpace(unitRaw)
			if unit != "" {
				if !strings.HasSuffix(unit, ".service") {
					unit += ".service"
				}
				journal.AddMatch("_SYSTEMD_UNIT=" + unit)
				journal.AddDisjunction()
			}
		}

	}
}
