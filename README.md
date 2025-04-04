# heracles
## 1. Creatures

### 1.1 Extraire dans un `data.json` la liste des creatures
```sql
SELECT
  t.entry,
  t.name,
  t.locale,
  m.CreatureDisplayID AS displayId
FROM
  (
    SELECT
      entry,
      name,
      'enUS' AS locale
    FROM
      creature_template
    UNION
    SELECT
      entry,
      Name AS name,
      locale
    FROM
      creature_template_locale
    WHERE
      locale IN ('frFR', 'esES', 'deDE')
  ) AS t
  LEFT JOIN creature_template_model AS m ON t.entry = m.CreatureId WHERE name != "";
```
### 1.2 Supprimer les combos name / locale en double
```python
import json

# Lecture du fichier JSON
try:
    with open('data.json', 'r', encoding='utf8') as file:
        data = file.read()
except Exception as err:
    print(f"Erreur de lecture du fichier JSON: {err}")
    exit()

# Conversion des données JSON en objet Python
json_data = json.loads(data)

# Filtrage des données pour supprimer les doublons basés sur la combinaison de locale et name
filtered_data = []
seen = set()

for item in json_data:
    # Crée une clé unique pour chaque combinaison de locale et name
    unique_key = f"{item['t.locale']}-{item['t.name']}"

    # Si cette combinaison n'a pas été rencontrée, on l'ajoute à la liste filtrée
    if unique_key not in seen:
        filtered_data.append(item)
        seen.add(unique_key)

# Sauvegarde des données filtrées dans un nouveau fichier JSON
try:
    with open('filteredData_py.json', 'w', encoding='utf8') as file:
        json.dump(filtered_data, file, ensure_ascii=False, indent=2)
    print("Fichier JSON filtré enregistré avec succès.")
except Exception as err:
    print(f"Erreur de sauvegarde du fichier JSON: {err}")
```

## 2. Zone / Subzones

### 2.1 Récupérer les données depuis un .dbc
Avec https://github.com/wowgaming/node-dbc-reader
```node
npm run start -- --out-type=sql --file=output.sql AreaTable
```
Attention ! En français / allemand / espagnol il sera nécessaire de convertir le fichier du format utf-8 au format latin.

Pour cela, il faut utiliser le script `./script/latin.js`.

### 2.2 Importer les données vers une base de données AzerothCore

Antares -> import

### 2.3 Récupérer les données

```sql
SELECT 
    MIN(entry) AS entry, 
    name, 
    parentId, 
    locale
FROM (
    SELECT 
        ID AS entry, 
        AreaName_Lang_enUS AS name, 
        AreaTableID AS parentId, 
        "enUS" AS locale
    FROM 
        wmoareatable_dbc_enUS
    WHERE 
        AreaName_Lang_enUS != ""
    
    UNION
    
    SELECT 
        ID AS entry, 
        AreaName_Lang_enUS AS name, 
        ParentAreaID AS parentId, 
        "enUS" AS locale
    FROM 
        areatable_dbc_enUS 
    WHERE 
        ParentAreaID != 0
) AS combined
WHERE entry NOT IN (
    55, 187, 407, 470, 474, 475, 476, 696, 697, 698, 699, 700, 701, 
    1118, 1580, 1597, 1598, 1600, 1601, 1602, 1603, 2137, 2259, 2777, 
    3237, 3609, 3635, 4621
)
GROUP BY 
    parentId, name, locale
UNION
SELECT 
    MIN(entry) AS entry, 
    name, 
    parentId, 
    locale
FROM (
    SELECT 
        ID AS entry, 
        AreaName_Lang_frFR AS name, 
        AreaTableID AS parentId, 
        "frFR" AS locale
    FROM 
        wmoareatable_dbc_frFR
    WHERE 
        AreaName_Lang_frFR != ""
    
    UNION
    
    SELECT 
        ID AS entry, 
        AreaName_Lang_frFR AS name, 
        ParentAreaID AS parentId, 
        "frFR" AS locale
    FROM 
        areatable_dbc_frFR 
    WHERE 
        ParentAreaID != 0
) AS combined
WHERE entry NOT IN (
    55, 187, 407, 470, 474, 475, 476, 696, 697, 698, 699, 700, 701, 
    1118, 1580, 1597, 1598, 1600, 1601, 1602, 1603, 2137, 2259, 2777, 
    3237, 3609, 3635, 4621
)
GROUP BY 
    parentId, name, locale
UNION
SELECT 
    MIN(entry) AS entry, 
    name, 
    parentId, 
    locale
FROM (
    SELECT 
        ID AS entry, 
        AreaName_Lang_esES AS name, 
        AreaTableID AS parentId, 
        "esES" AS locale
    FROM 
        wmoareatable_dbc_esES
    WHERE 
        AreaName_Lang_esES != ""
    
    UNION
    
    SELECT 
        ID AS entry, 
        AreaName_Lang_esES AS name, 
        ParentAreaID AS parentId, 
        "esES" AS locale
    FROM 
        areatable_dbc_esES 
    WHERE 
        ParentAreaID != 0
) AS combined
WHERE entry NOT IN (
    55, 187, 407, 470, 474, 475, 476, 696, 697, 698, 699, 700, 701, 
    1118, 1580, 1597, 1598, 1600, 1601, 1602, 1603, 2137, 2259, 2777, 
    3237, 3609, 3635, 4621
)
GROUP BY 
    parentId, name, locale
UNION
SELECT 
    MIN(entry) AS entry, 
    name, 
    parentId, 
    locale
FROM (
    SELECT 
        ID AS entry, 
        AreaName_Lang_deDE AS name, 
        AreaTableID AS parentId, 
        "deDE" AS locale
    FROM 
        wmoareatable_dbc_deDE
    WHERE 
        AreaName_Lang_deDE != ""
    
    UNION
    
    SELECT 
        ID AS entry, 
        AreaName_Lang_deDE AS name, 
        ParentAreaID AS parentId, 
        "deDE" AS locale
    FROM 
        areatable_dbc_deDE 
    WHERE 
        ParentAreaID != 0
) AS combined
WHERE entry NOT IN (
    55, 187, 407, 470, 474, 475, 476, 696, 697, 698, 699, 700, 701, 
    1118, 1580, 1597, 1598, 1600, 1601, 1602, 1603, 2137, 2259, 2777, 
    3237, 3609, 3635, 4621
)
GROUP BY 
    parentId, name, locale
```
