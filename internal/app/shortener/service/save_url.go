package service

func (s *ServiceImpl) SaveURL(url string) string {
	var shortURL string
	count := 0
	hashURL := HashString(url)
	fifthLength := len(hashURL) / 5

	for {
		// Обрезаем hashURL до нужной длины
		shortURL = hashURL[:fifthLength+count]

		// Проверяем, существует ли уже этот короткий URL
		if _, ok := s.OriginalURLsMap[shortURL]; !ok {
			// Если нет, сохраняем его и выходим из цикла
			s.OriginalURLsMap[shortURL] = url
			return shortURL
		}

		// Увеличиваем count для следующей итерации
		count++

		// Проверяем, не достигли ли мы максимальной длины
		if fifthLength+count > len(hashURL) {
			// Если да, возвращаем пустую строку
			return ""
		}
	}
}
