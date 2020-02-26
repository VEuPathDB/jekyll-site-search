package solr

import (
	"fmt"
	"time"
)

const (
	batchType = "jekyll"
	batchName = "all"
)

type Batch struct {
	BatchType string `solr:"batch-type"`
	BatchId   string `solr:"batch-id"`
	BatchTime int64  `solr:"batch-timestamp"`
	BatchName string `solr:"batch-name"`
	Id        string `solr:"id"`
}

func NewBatch() *Batch {
	now := time.Now().UnixNano() / int64(time.Millisecond)
	id  := genBatchId(now)

	return &Batch {
		BatchType: batchType,
		BatchName: batchName,
		BatchId:   id,
		BatchTime: now,
		Id:        id,
	}
}

func genBatchId(millis int64) string {
	return fmt.Sprintf("%s_%s_%d", batchType, batchName, millis)
}
