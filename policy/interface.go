package policy

type Policy struct {
	Name     string
	Behavior string
}

type PolicyCache interface {
	// GetPolicy(string) Policy
	ListPolicies() *[]Policy
}
