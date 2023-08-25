package designModule

import "context"

// 对规则链 Interface 进行定义
type RuleChain interface {
	Apply(ctx context.Context, params map[string]interface{}) error
	Next() RuleChain
}

type baseRuleChain struct {
	next RuleChain
}

// baseRuleChain 的 Apply 方法未做具体实现，需要由继承类进行实现
func (b *baseRuleChain) Apply(ctx context.Context, params map[string]interface{}) error {
	panic("not implement")
}

func (b *baseRuleChain) Next() RuleChain {
	return b.next
}
