/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package operations

import (
	"sync"

	"github.com/hyperledger/udo/common/metrics"
	"github.com/hyperledger/udo/common/metrics/prometheus"
)

var (
	udoVersion = metrics.GaugeOpts{
		Name:         "udo_version",
		Help:         "The active version of UDO.",
		LabelNames:   []string{"version"},
		StatsdFormat: "%{#fqname}.%{version}",
	}

	gaugeLock        sync.Mutex
	promVersionGauge metrics.Gauge
)

func versionGauge(provider metrics.Provider) metrics.Gauge {
	switch provider.(type) {
	case *prometheus.Provider:
		gaugeLock.Lock()
		defer gaugeLock.Unlock()
		if promVersionGauge == nil {
			promVersionGauge = provider.NewGauge(udoVersion)
		}
		return promVersionGauge

	default:
		return provider.NewGauge(udoVersion)
	}
}
