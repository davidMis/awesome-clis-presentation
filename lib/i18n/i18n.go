package i18n

import "github.com/pkg/errors"

// Messages is an i18n dictionary. It's keys are string labels, and the values are
// internationalized strings.
var Messages map[string]string

// UseLanguage populates an i18n dictionary. The 'language' parameter must be a 2-letter
// ISO language code. In a real application, we would probably load the dictionary from
// a file, but for this toy it's okay to just build the maps on the fly.
func UseLanguage(language string) error {
	switch language {
	case "EN":
		Messages = map[string]string{
			"main": "I love cookies",
		}
	case "ES":
		Messages = map[string]string{
			"main": "Yo quierro cookies",
		}
	default:
		return errors.New("unsupported language " + language)
	}
	
	return nil
}
