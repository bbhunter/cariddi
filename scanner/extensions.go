package scanner

func GetExtensions() map[string]string {
	//extensions contains a list of known extensions
	//and the TYPICAL (also say `in general`) associated severity.
	var extensions = map[string]string{
		"key":        "1",
		"env":        "1",
		"pem":        "1",
		"git":        "1",
		"ovpn":       "1",
		"log":        "1",
		"secret":     "1",
		"secrets":    "1",
		"sh":         "2",
		"py":         "2",
		"json":       "2",
		"xml":        "2",
		"yml":        "2",
		"yaml":       "2",
		"properties": "2",
		"toml":       "2",
		"zip":        "3",
		"doc":        "3",
		"docx":       "3",
		"odt":        "3",
		"xls":        "3",
		"xlsx":       "3",
		"txt":        "3",
		"ts":         "4",
		"js":         "4",
		"php":        "5",
		"phtml":      "5",
		"php5":       "5",
		"html":       "6",
		"htm":        "6",
		"pdf":        "7",
	}

	return extensions
}
