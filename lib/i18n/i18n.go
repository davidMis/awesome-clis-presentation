package i18n

import (
	"github.com/pkg/errors"
	"log"
)

// Messages is an i18n dictionary. It's keys are string labels, and the values are
// internationalized strings.
var Messages map[string]string


// UseLanguage populates an i18n dictionary. The 'language' parameter must be a 2-letter
// ISO language code. In a real application, we would probably load the dictionary from
// a file, but for this toy it's okay to just build the maps on the fly.
func UseLanguage(language string) error {
	log.SetPrefix("[DEBUG] ")
	log.Printf("Using language \"%s\"", language)

	switch language {
	case "EN":
		Messages = map[string]string{
			"main": "I love cookies\n",
			"eatCookie": "YUM! I ate a cookie!\n",
			"eatOther": "YUCK! I ate a %s\n",
			"prepareCookie": "Time to bake some cookies\n",
			"prepareOther": "I don't know how to prepare %s\n",
		}
	case "ES":
		Messages = map[string]string{
			"main": "Me encanta las galletas\n",
			"eatCookie": "MMM! Comí la galleta\n",
			"eatOther": "YUCK! Comí %s\n",
			"prepareCookie": "Vamanos a hornear\n",
			"prepareOther": "Yo no se como cocinar %s\n",
		}
	default:
		return errors.New("unsupported language " + language)
	}

	return nil
}
