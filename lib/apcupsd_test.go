package apcupsdplugin

import (
	"os/exec"
	"reflect"
	"testing"

	mp "github.com/mackerelio/go-mackerel-plugin"
)

func TestGraphDefinition(t *testing.T) {
	u := APCUPSPlugin{}
	graph := u.GraphDefinition()

	expected := map[string]mp.Graphs{
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

	if !reflect.DeepEqual(graph, expected) {
		t.Errorf("GraphDefinition() = %v, want %v", graph, expected)
	}
}

func TestMetricKeyPrefix(t *testing.T) {
	u := APCUPSPlugin{}
	prefix := u.MetricKeyPrefix()

	if prefix != "ups" {
		t.Errorf("MetricKeyPrefix() = %v, want %v", prefix, "ups")
	}
}

func TestFetchMetrics(t *testing.T) {
	// apcaccessコマンドが利用可能かチェック
	_, err := exec.LookPath("apcaccess")
	if err != nil {
		t.Skip("apcaccess command not available, skipping test")
	}

	u := APCUPSPlugin{}
	metrics, err := u.FetchMetrics()

	if err != nil {
		t.Fatalf("FetchMetrics() error = %v", err)
	}

	expectedKeys := []string{"linev", "loadpct", "bcharge", "timeleft", "battv"}
	for _, key := range expectedKeys {
		if _, ok := metrics[key]; !ok {
			t.Errorf("FetchMetrics() missing key = %v", key)
		}
	}

	for key, value := range metrics {
		if value < 0 {
			t.Errorf("FetchMetrics() got negative value for %v: %v", key, value)
		}
	}
}
