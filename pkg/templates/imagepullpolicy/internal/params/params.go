package params

// Params represents the params accepted by this template.
type Params struct {
	// list of bad image pull policy
	ForbiddenPolicies []string
}
