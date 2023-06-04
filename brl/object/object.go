package object

type env struct {
	store map[int]Object
}

func NewEnv() *env {
	s := make(map[int]Object)
	return &env{store: s}
}
