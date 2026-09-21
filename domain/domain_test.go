package lowpricesampling

import (
    "os"
    "testing"
)

func TestFixtureMatchesDomain(t *testing.T) {
    raw, err := os.ReadFile("../fixtures/domain.json")
    if err != nil { t.Fatal(err) }
    value, err := Parse(raw)
    if err != nil { t.Fatal(err) }
    if value.Domain != "low-price-sampling" { t.Fatalf("领域标识不一致: %s", value.Domain) }
}

func TestFixtureKeepsRedLines(t *testing.T) {
    raw, err := os.ReadFile("../fixtures/domain.json")
    if err != nil { t.Fatal(err) }
    value, err := Parse(raw)
    if err != nil { t.Fatal(err) }
    if !contains(value.Constraints, "低价本身绝不能直接成为违法结论") { t.Fatal("缺少低价不定罪红线") }
    if !contains(value.Constraints, "现场人员只接收执行所需线索") { t.Fatal("缺少最小知情约束") }
    if !contains(value.RiskSignals, "价格分布") || !contains(value.RiskSignals, "渠道曝光") { t.Fatal("风险信号不完整") }
    if !contains(value.ReplanTriggers, "主体改名") || !contains(value.ReplanTriggers, "风险突然升级") { t.Fatal("重排触发不完整") }
    if !contains(value.ReviewMetrics, "命中率") || !contains(value.ReviewMetrics, "商家规避行为") { t.Fatal("复盘指标不完整") }
}

func contains(items []string, want string) bool {
    for _, item := range items {
        if item == want { return true }
    }
    return false
}
