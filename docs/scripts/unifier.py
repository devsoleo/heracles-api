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
    with open('output.json', 'w', encoding='utf8') as file:
        json.dump(filtered_data, file, ensure_ascii=False, indent=2)
    print("Fichier JSON filtré enregistré avec succès.")
except Exception as err:
    print(f"Erreur de sauvegarde du fichier JSON: {err}")
