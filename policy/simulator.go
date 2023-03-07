package policy

type PolicySimulator struct{}

var ViolateFDInstall = Policy{
	Name:     "A",
	Behavior: "fd_install",
}

var policies []Policy = []Policy{ViolateFDInstall}

func (ps *PolicySimulator) ListPolicies() *[]Policy {
	return &policies
}
