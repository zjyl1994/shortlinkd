package vars

import "fmt"

type ConfigS struct {
	Fallback string
	List     map[string]ListItemS
}

type ListItemS struct {
	URL      string
	Expired  string
	Disabled bool
}

func (m *ListItemS) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type sType ListItemS
	var s sType
	if err := unmarshal(&s); err == nil {
		m.Disabled = s.Disabled
		m.Expired = s.Expired
		m.URL = s.URL
		return nil
	}

	var str string
	if err := unmarshal(&str); err == nil {
		m.URL = str
		return nil
	}

	return fmt.Errorf("failed to unmarshal ListItemS")
}
