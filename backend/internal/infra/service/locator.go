package service

const DefaultLocatorServices = 100

type LocatorInterface interface {
	Add(key string, service interface{})
	Get(key string) interface{}
	Remove(key string)
}

type Locator map[string]interface{}

func NewLocator() Locator {
	return make(Locator, DefaultLocatorServices)
}

func (l Locator) Add(key string, service interface{}) {
	if _, exist := l[key]; exist {
		panic("service already exist")
	}
	l[key] = service
}

func (l Locator) Get(key string) interface{} {
	return l[key]
}

func (l Locator) Remove(key string) {
	delete(l, key)
}
