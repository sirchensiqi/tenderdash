// Code generated by metricsgen. DO NOT EDIT.

package statesync

import (
	"github.com/go-kit/kit/metrics/discard"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func PrometheusMetrics(namespace string, labelsAndValues ...string) *Metrics {
	labels := []string{}
	for i := 0; i < len(labelsAndValues); i += 2 {
		labels = append(labels, labelsAndValues[i])
	}
	return &Metrics{
		TotalSnapshots: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "total_snapshots",
			Help:      "The total number of snapshots discovered.",
		}, labels).With(labelsAndValues...),
		ChunkProcessAvgTime: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "chunk_process_avg_time",
			Help:      "The average processing time per chunk.",
		}, labels).With(labelsAndValues...),
		SnapshotHeight: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "snapshot_height",
			Help:      "The height of the current snapshot the has been processed.",
		}, labels).With(labelsAndValues...),
		SnapshotChunk: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "snapshot_chunk",
			Help:      "The current number of chunks that have been processed.",
		}, labels).With(labelsAndValues...),
		BackFilledBlocks: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "back_filled_blocks",
			Help:      "The current number of blocks that have been back-filled.",
		}, labels).With(labelsAndValues...),
		BackFillBlocksTotal: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "back_fill_blocks_total",
			Help:      "The total number of blocks that need to be back-filled.",
		}, labels).With(labelsAndValues...),
	}
}

func NopMetrics() *Metrics {
	return &Metrics{
		TotalSnapshots:      discard.NewCounter(),
		ChunkProcessAvgTime: discard.NewGauge(),
		SnapshotHeight:      discard.NewGauge(),
		SnapshotChunk:       discard.NewCounter(),
		BackFilledBlocks:    discard.NewCounter(),
		BackFillBlocksTotal: discard.NewGauge(),
	}
}