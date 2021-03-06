package shared

type PluginNotFoundError struct {
	Name string
}

func (e PluginNotFoundError) Error() string {
	return "Plugin {{.Name}} does not exist."
}

func (e PluginNotFoundError) Translate(translate func(string, ...interface{}) string) string {
	return translate(e.Error(), map[string]interface{}{
		"Name": e.Name,
	})
}
