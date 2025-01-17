package apcupsdplugin

import (
	"os/exec"
	"regexp"
	"strconv"

	mp "github.com/mackerelio/go-mackerel-plugin"
)

// APCUPSPlugin: Name of graph
type APCUPSPlugin struct {
	prefix string
}

// Definition of the graph metric
func (u APCUPSPlugin) GraphDefinition() map[string]mp.Graphs {
	return map[string]mp.Graphs{
		"ups": {
			Label: "UPS Metrics",
			Unit:  "float",
			Metrics: []mp.Metrics{
				{Name: "linev", Label: "Input Voltage", Diff: false},
				{Name: "loadpct", Label: "Load Percentage", Diff: false},
				{Name: "bcharge", Label: "Battery Charge", Diff: false},
				{Name: "timeleft", Label: "Time Left", Diff: false},
				{Name: "battv", Label: "Battery Voltage", Diff: false},
			},
		},
	}
}

// execute and parse
func (u APCUPSPlugin) FetchMetrics() (map[string]float64, error) {
	output, err := exec.Command("apcaccess").Output()
	if err != nil {
		return nil, err
	}

	metrics := make(map[string]float64)
	patterns := map[string]*regexp.Regexp{
		"linev":    regexp.MustCompile(`LINEV\s+:\s+([\d.]+)\s+Volts`),
		"loadpct":  regexp.MustCompile(`LOADPCT\s+:\s+([\d.]+)\s+Percent`),
		"bcharge":  regexp.MustCompile(`BCHARGE\s+:\s+([\d.]+)\s+Percent`),
		"timeleft": regexp.MustCompile(`TIMELEFT\s+:\s+([\d.]+)\s+Minutes`),
		"battv":    regexp.MustCompile(`BATTV\s+:\s+([\d.]+)\s+Volts`),
	}

	for key, pattern := range patterns {
		matches := pattern.FindSubmatch(output)
		if len(matches) > 1 {
			value, err := strconv.ParseFloat(string(matches[1]), 64)
			if err == nil {
				metrics[key] = value
			}
		}
	}

	return metrics, nil
}

// failsafe
func (u APCUPSPlugin) MetricKeyPrefix() string {
	if u.prefix == "" {
		u.prefix = "ups"
	}
	return u.prefix
}

// Main
func Do() {
	u := APCUPSPlugin{}
	plugin := mp.NewMackerelPlugin(u)
	plugin.Run()
}
