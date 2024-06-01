package code

func GetCode(code string) *CodeItem {
	lock.RLock()
	defer lock.RUnlock()
	if data, ok := data[code]; ok {
		return &data
	}
	return nil
}

func InitCode(items []CodeItem) {
	newData := make(map[string]CodeItem)
	for _, item := range items {
		newData[item.Code] = item
	}
	lock.Lock()
	defer lock.Unlock()
	data = newData
}

func ListCodes() []CodeItem {
	lock.RLock()
	defer lock.RUnlock()
	result := make([]CodeItem, 0, len(data))
	for _, item := range data {
		result = append(result, item)
	}
	return result
}
