package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

// RuleType defines a type of rule
type RuleType int

// types of rules
const (
	SHOW RuleType = 0
	HIDE RuleType = 1
)

/*
Rule is a singular rule
*/
type Rule struct {
	Day     string `json:"day"`
	Before  string `json:"before"`
	After   string `json:"after"`
	Channel string `json:"channel"`
	Type    RuleType
}

/*
Mode is many rules
*/
type Mode struct {
	Name      string   `json:"name"`
	Overrides []Rule   `json:"overrides"`
	Order     []string `json:"order"`
}

// Modes is many mode
type Modes struct {
	Available map[string]Mode
	Current   Mode
}

// Switch modes
func (modes *Modes) Switch(newMode string) Mode {
	new := modes.Available[newMode]
	if new.Name == "" {
		log.Println("mode", newMode, "does not exist")
		return modes.Current
	}

	modes.Current = new
	log.Println("switching to", modes.Current.Name, "mode")

	return modes.Current
}

// ShouldAlert whether we should alert the user of the message
func (mode Mode) ShouldAlert(ID string) bool {
	return ShouldAlert(mode, ID, time.Now())
}

// ReadRules will read all rules from json file
func ReadRules(filePath string) Modes {
	// read json from file into interface
	file, _ := ioutil.ReadFile(filePath)
	var data map[string]interface{}
	json.Unmarshal([]byte(file), &data)

	// schema is useless to us, remove it
	delete(data, "$schema")

	structuredData := make(map[string]Mode)
	// for each mode
	for k, v := range data {
		mode := Mode{
			Name: k,
		}

		// safely get values
		mappedData := v.(map[string]interface{})
		iOrder := mappedData["order"].([]interface{})
		for _, u := range iOrder {
			mode.Order = append(mode.Order, u.(string))
		}

		// safely get overrides
		iOverrides := mappedData["overrides"].(map[string]interface{})
		for channelName, override := range iOverrides {
			arrayOfRules, isArrayOfRule := override.([]interface{})
			rule, isNotArrayOfRule := override.(map[string]interface{})

			if isArrayOfRule {
				for _, u := range arrayOfRules {
					mode.Overrides = append(mode.Overrides, readRule(channelName, u))
				}
			} else if isNotArrayOfRule {
				mode.Overrides = append(mode.Overrides, readRule(channelName, rule))
			}
		}

		structuredData[k] = mode
	}

	return Modes{
		Available: structuredData,
	}
}

func readRule(name string, u interface{}) Rule {
	rule, _ := u.(map[string]interface{})
	after, ok := rule["after"].(string)
	if !ok {
		after = "23:59"
	}

	before, ok := rule["before"].(string)
	if !ok {
		before = "00:00"
	}

	day, ok := rule["day"].(string)
	if !ok {
		day = "mon-sun"
	}

	ruleType, ok := rule["type"].(string)
	enumRuleType := HIDE
	if !ok || (ruleType == "show") {
		enumRuleType = SHOW
	}

	return Rule{
		Channel: name,
		After:   after,
		Before:  before,
		Day:     day,
		Type:    enumRuleType,
	}
}

// ShouldAlert will determine if we should alert
func ShouldAlert(mode Mode, ID string, timeNow time.Time) bool {
	log.Println(mode)
	if mode.Overrides == nil {
		log.Println("no overrides")
		for _, v := range mode.Order {
			return v == "show"
		}

		log.Println("no order set")
		// default to false
		return false
	}

	// run through overrides
	for _, override := range mode.Overrides {
		day := strDayToInt(override.Day)
		now := timeNow
		// @todo: this doesn't support multiple day rules
		// rule is concerned with a different day - ignore it
		if day != now.Day() {
			continue
		}

		beforeTime, _ := time.Parse("15:04", override.Before)
		afterTime, _ := time.Parse("15:04", override.After)

		// after "before time" - ignore
		if beforeTime.Hour() < now.Hour() || beforeTime.Hour() == now.Hour() && beforeTime.Minute() < now.Minute() {
			log.Println("before time doesn't match")
			continue
		}

		// before "after time" - ignore
		if afterTime.Hour() > now.Hour() || afterTime.Hour() == now.Hour() && afterTime.Minute() > now.Minute() {
			log.Println("after time doesn't match")
			continue
		}

		// person doesn't match - ignore
		if override.Channel != ID {
			log.Println("channel doesn't match")
			continue
		}

		return override.Type == SHOW
	}

	// if it fails for whatever reason, don't alert
	return mode.Order[0] == "show"
}

func strDayToInt(day string) int {
	dayOfWeek := 1
	switch day {
	case "mon":
		dayOfWeek = 1
	case "tues":
		dayOfWeek = 2
	case "wed":
		dayOfWeek = 3
	case "thurs":
		dayOfWeek = 4
	case "fri":
		dayOfWeek = 5
	case "sat":
		dayOfWeek = 6
	case "sun":
		dayOfWeek = 7
	}

	return dayOfWeek
}
