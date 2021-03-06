package responses

import (
	"log"

	"github.com/fatih/structs"
	"gopkg.in/yaml.v1"
)

type Response struct {
	Text     map[string]map[string]string "text"
	Triggers map[string][]string          "triggers"
}

type Content struct {
	Stop             Response "stop"
	PollingLocation  Response "pollingLocation"
	DropOffLocation  Response "dropOffLocation"
	ElectionOfficial Response "electionOfficial"
	ChangeLanguage   Response "changeLanguage"
	Registration     Response "registration"
	Elo              Response "elo"
	Help             Response "help"
	About            Response "about"
	Intro            Response "intro"
	LastContact      Response "lastContact"
	Errors           Response "errors"
}

func Load(raw []byte) (*Content, map[string]map[string]string) {
	c := &Content{}

	err := yaml.Unmarshal(raw, c)
	if err != nil {
		log.Panic("[ERROR] Failed to parse responses : ", err)
	}

	return c, buildTriggerLookup(c)
}

func buildTriggerLookup(c *Content) map[string]map[string]string {
	lookup := make(map[string]map[string]string)

	for _, field := range structs.New(c).Fields() {
		name := field.Name()

		triggers := field.Value().(Response).Triggers

		for language, words := range triggers {
			if lookup[language] == nil {
				lookup[language] = make(map[string]string)
			}
			for _, word := range words {
				lookup[language][word] = name
			}
		}
	}

	return lookup
}
