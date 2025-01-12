package service

type ServiceImpl struct {
	OriginalURLsMap map[string]string
}

func NewService() *ServiceImpl {
	return &ServiceImpl{OriginalURLsMap: make(map[string]string)}
}
