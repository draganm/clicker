package histogram

import (
	"time"

	"github.com/draganm/clicker/stats"
)

// HistogramCollector collects last counts for given bucket
// count and duration.
type HistogramCollector struct {
	bucketSize     time.Duration
	buckets        stats.Data
	lastBucketTime time.Time
}

func NewHistogramCollector(bucketSize time.Duration, buckets int) *HistogramCollector {

}
