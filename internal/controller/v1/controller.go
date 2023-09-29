package v1

type fileHandler interface{}

type controller struct{}

func New() *controller {
	return &controller{}
}
