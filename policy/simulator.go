package policy

type PolicySimulator struct{}

var policies []Policy = []Policy{"😉", "👍", "😭", "😄", "🌝"}

func (ps *PolicySimulator) ListPolicies() *[]Policy {
	return &policies
}
