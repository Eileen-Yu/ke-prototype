package policy

type Policy string

type PolicyCache interface {
	// GetPolicy(string) Policy
	ListPolicies() *[]Policy
}
