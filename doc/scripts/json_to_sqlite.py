import sqlite3
import json

# 1. Charger les données JSON
def charger_json(fichier_json):
    with open(fichier_json, 'r', encoding='utf-8') as f:
        return json.load(f)

# 2. Se connecter à la base de données SQLite (elle sera créée si elle n'existe pas)
def se_connecter_db(nom_db):
    conn = sqlite3.connect(nom_db)
    return conn

# 3. Créer une table dans SQLite
def creer_table(conn):
    cursor = conn.cursor()
    cursor.execute('''
    CREATE TABLE IF NOT EXISTS subzones (
        entry INTEGER,
        name TEXT,
        parentId INTEGER,
        locale TEXT
    )
    ''')
    conn.commit()

# 4. Insérer des données dans la table SQLite
def inserer_donnees(conn, donnees):
    cursor = conn.cursor()
    for item in donnees:
        cursor.execute('''
        INSERT INTO subzones (entry, name, parentId, locale)
        VALUES (?, ?, ?, ?)
        ''', (item['.entry'], item['.name'], item['.parentId'], item['.locale']))  # Adapte ces clés en fonction de ton fichier JSON
    conn.commit()

# 5. Exemple d'utilisation
def alimenter_db(fichier_json, nom_db):
    # Charger les données JSON
    donnees = charger_json(fichier_json)
    
    # Se connecter à la base de données
    conn = se_connecter_db(nom_db)
    
    # Créer la table si elle n'existe pas
    creer_table(conn)
    
    # Insérer les données dans la table
    inserer_donnees(conn, donnees)
    
    # Fermer la connexion
    conn.close()

# Appel de la fonction pour alimenter la base de données SQLite
alimenter_db('subzones.json', 'database.db')
