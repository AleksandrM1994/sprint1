package service

type ServiceImpl struct {
	originalURL string
}

func NewService(originalURL string) *ServiceImpl {
	return &ServiceImpl{originalURL: originalURL}
}
