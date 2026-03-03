package runner

type Runner interface {
	Run(inp string) map[string][]float64
}
