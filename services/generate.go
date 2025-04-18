package services

import (
	"devsoleo/heracles-api/database"
	"encoding/json"
)

const FMT_VERSION = "1.0.0"

type Translation struct {
	Name   string `json:"name"`
	Locale string `json:"locale"`
}

type KillForm struct {
	Category string `json:"category"`
	Amount   string `json:"amount"`
	Creature string `json:"creature"`
}

type GotoForm struct {
	Category string `json:"category"`
	Zone     string `json:"zone"`
	Subzone  string `json:"subzone"`
}

type TargetForm struct {
	Category string `json:"category"`
	Entity   string `json:"entity"`
	Target   string `json:"target"`
}

func Generate(rawForms []map[string]interface{}) (any, error) {
	var headers = map[string]interface{}{
		"format": FMT_VERSION,
	}

	var missions []map[string]interface{}

	// For each mission
	for _, raw := range rawForms {
		if rawType, ok := raw["category"]; ok {
			jsonData, _ := json.Marshal(raw)

			switch rawType {
			case "kill":
				var form KillForm
				json.Unmarshal(jsonData, &form)

				missions = append(missions, map[string]interface{}{
					"category": form.Category,
					"amount":   form.Amount,
					"creature": getSecureLocales("creatures", form.Creature),
				})
			case "goto":
				var form GotoForm
				json.Unmarshal(jsonData, &form)

				missions = append(missions, map[string]interface{}{
					"category": form.Category,
					"zone":     getSecureLocales("zones", form.Zone),
					"subzone":  getSecureLocales("subzones", form.Subzone),
				})
			case "target":
				var form TargetForm
				json.Unmarshal(jsonData, &form)

				var target any

				target = form.Target

				if form.Entity == "creature" {
					target = getSecureLocales("creatures", form.Target)
				}

				missions = append(missions, map[string]interface{}{
					"category": form.Category,
					"entity":   form.Entity,
					"name":     target,
				})
			}
		}
	}

	var result = map[string]interface{}{
		"headers":  headers,
		"missions": missions,
	}

	return result, nil
}

func getLocale(category string, entry string, locale string) string {
	db := database.GetDB()

	var translation Translation

	if err := db.QueryRow("SELECT name FROM "+category+" WHERE entry = ? AND locale = ?", entry, locale).Scan(&translation.Name); err != nil {
		return ""
	}

	return translation.Name
}

func getSecureLocales(category string, entry string) map[string]interface{} {
	return map[string]interface{}{
		"enUS": getLocale(category, entry, "enUS"),
		"frFR": getLocale(category, entry, "frFR"),
		"esES": getLocale(category, entry, "esES"),
		"deDE": getLocale(category, entry, "deDE"),
	}
}
