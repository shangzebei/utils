package lb

type LB interface {
	SelectOne([]string) (string, error)
}
