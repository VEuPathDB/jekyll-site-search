package solr

import (
	"fmt"
	"time"
)

const (
	batchType = "jekyll"
	batchName = "all"
	docType   = "batch-meta"
)

type Batch struct {
	BatchType string `json:"batch-type"`
	BatchId   string `json:"batch-id"`
	BatchTime int64  `json:"batch-timestamp"`
	BatchName string `json:"batch-name"`
	Id        string `json:"id"`
	DocType   string `json:"document-type"`
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
		DocType:   docType,
	}
}

func genBatchId(millis int64) string {
	return fmt.Sprintf("%s_%s_%d", batchType, batchName, millis)
}
