package service

func (s *ServiceImpl) Ping() error {
	return s.repo.Ping()
}
