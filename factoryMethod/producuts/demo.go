package producuts

type Demo struct {
	*Gun
}

func NewDemo() IGun {
	return &Demo{
		Gun: &Gun{
			name:  "demo",
			power: 1,
		},
	}
}
