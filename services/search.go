package services

import (
	"devsoleo/heracles-api/database"
)

type Creature struct {
	Entry  int    `json:"entry"`
	Name   string `json:"name"`
	Locale string `json:"locale"`
}

type Zone struct {
	Entry  int    `json:"entry"`
	Name   string `json:"name"`
	Locale string `json:"locale"`
}

type Subzone struct {
	Entry    int    `json:"entry"`
	Name     string `json:"name"`
	Locale   string `json:"locale"`
	ParentId int    `json:"parentId"`
}

func Search(category string, query string, locale string, extra ...any) (any, error) {
	switch category {
	case "creature":
		return SearchCreatures(query, locale)
	case "zone":
		return SearchZones(query, locale)
	case "subzone":
		parentId := extra[0].(string)
		return SearchSubzones(query, locale, parentId)
	case "subzone_preload":
		parentId := extra[0].(string)
		return SearchSubzonesPreload(locale, parentId)
	}

	return nil, nil
}

func SearchCreatures(query string, locale string) ([]Creature, error) {
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM creatures WHERE name LIKE ? AND locale = ?", "%"+query+"%", locale)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var creatures []Creature
	for rows.Next() {
		var creature Creature
		if err := rows.Scan(&creature.Entry, &creature.Name, &creature.Locale); err != nil {
			return nil, err
		}
		creatures = append(creatures, creature)
	}

	return creatures, nil
}

func SearchZones(query string, locale string) ([]Zone, error) {
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM zones WHERE name LIKE ? AND locale = ?", "%"+query+"%", locale)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var zones []Zone
	for rows.Next() {
		var zone Zone
		if err := rows.Scan(&zone.Entry, &zone.Name, &zone.Locale); err != nil {
			return nil, err
		}
		zones = append(zones, zone)
	}
	return zones, nil
}

func SearchSubzones(query string, locale string, parentId string) ([]Subzone, error) {
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM subzones WHERE name LIKE ? AND locale = ? AND parentId = ?", "%"+query+"%", locale, parentId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subzones []Subzone
	for rows.Next() {
		var subzone Subzone
		if err := rows.Scan(&subzone.Entry, &subzone.Name, &subzone.ParentId, &subzone.Locale); err != nil {
			return nil, err
		}
		subzones = append(subzones, subzone)
	}

	return subzones, nil
}

func SearchSubzonesPreload(locale string, parentId string) ([]Subzone, error) {
	db := database.GetDB()

	rows, err := db.Query("SELECT * FROM subzones WHERE locale = ? AND parentId = ?", locale, parentId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subzones []Subzone
	for rows.Next() {
		var subzone Subzone
		if err := rows.Scan(&subzone.Entry, &subzone.Name, &subzone.ParentId, &subzone.Locale); err != nil {
			return nil, err
		}
		subzones = append(subzones, subzone)
	}

	return subzones, nil
}
