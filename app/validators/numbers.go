package validators

type Number struct {
	Value string `query:"number" validate:"required,numeric"`
}
