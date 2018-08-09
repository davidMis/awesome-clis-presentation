package lib

import "github.com/pkg/errors"

// Messages is an i18n dictionary. It's keys are string labels, and the values are
// internationalized strings.
type Messages map[string]string

// NewMessages populates an i18n dictionary. The 'language' parameter must be a 2-letter
// ISO language code. In a real application, we would probably load the dictionary from
// a file, but for this toy it's okay to just build the maps on the fly.
func NewMessages(language string) (m Messages, err error) {
	switch language {
	case "EN":
		m = map[string]string{
			"main": "I love cookies",
		}
	case "ES":
		m = map[string]string{
			"main": "Yo quierro cookies",
		}
	default:
		return nil, errors.New("unsupported language " + language)
	}

	return m, nil
}