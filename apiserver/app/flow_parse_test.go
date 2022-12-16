package app

import (
	"encoding/json"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestFlowJsonParse(t *testing.T) {

	rule := flow.Rule{
		Resource:               "GET:/link/preview-token",
		MetricType:             0,
		TokenCalculateStrategy: flow.Direct,
		ControlBehavior:        flow.Reject,
		Count:                  3,
	}

	ss := rule.String()
	println(ss)

	src := `[{
    "resource":"GET:/link/preview-token",
    "metricType":0,
    "tokenCalculateStrategy":0,
    "controlBehavior":0,
    "count": 3
}]`
	rules := make([]*flow.Rule, 0)

	err := json.Unmarshal([]byte(src), &rules)
	if err != nil {
		println(err.Error())
	}
	println("再返回来--------Marshal")
	bytes, _ := json.Marshal(&rules)
	println(string(bytes))

	assert.Equal(t, rule, *rules[0])
}
