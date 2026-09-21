package lowpricesampling

import (
    "encoding/json"
    "errors"
)

// Record 表示项目共享的领域资料。
type Record struct {
    Domain string `json:"domain"`
    Version int `json:"version"`
    SampleID string `json:"sample_id"`
    Actors []string `json:"actors"`
    Facts []string `json:"facts"`
    Constraints []string `json:"constraints"`
    RiskSignals []string `json:"risk_signals"`
    ReplanTriggers []string `json:"replan_triggers"`
    ReviewMetrics []string `json:"review_metrics"`
}

// Parse 读取并检查带版本的业务资料。
func Parse(raw []byte) (Record, error) {
    var value Record
    if err := json.Unmarshal(raw, &value); err != nil { return Record{}, err }
    if value.Domain == "" || value.Version < 1 || value.SampleID == "" || len(value.Actors) < 2 || len(value.Facts) < 2 || len(value.Constraints) < 2 || len(value.RiskSignals) < 2 || len(value.ReplanTriggers) < 2 || len(value.ReviewMetrics) < 2 {
        return Record{}, errors.New("共享资料缺少必要字段")
    }
    return value, nil
}
