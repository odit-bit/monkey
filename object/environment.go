package object

func NewEnvironment() *Environment {
	return &Environment{
		store: map[string]Object{},
	}
}

type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnclosed(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func (e *Environment) Get(key string) (Object, bool) {
	obj, ok := e.store[key]
	if !ok && e.outer != nil {
		return e.outer.Get(key)
	}
	return obj, ok
}

func (e *Environment) Set(key string, obj Object) Object {
	e.store[key] = obj
	return obj
}
