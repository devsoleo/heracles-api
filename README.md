[![GoLang](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white&style=for-the-badge)](https://golang.org/)
[![SQLite](https://img.shields.io/badge/SQLite-003B57?logo=sqlite&logoColor=white&style=for-the-badge)](https://sqlite.org/)
[![Gin](https://img.shields.io/badge/Gin-00A97F?style=for-the-badge)](https://gin-gonic.com/)
[![Discord](https://img.shields.io/discord/1070782232219877477?color=7289da&label=Discord&logo=discord&style=for-the-badge)](https://discord.gg/vhs87h6fvD)

# Heracles API

> Backend service for the Heracles WebUI

## 🌐 Introduction

**Heracles API** is a lightweight backend service written in **Go**, using the **Gin** framework and **SQLite** as its database engine.

It powers the [Heracles WebUI](https://gitlab.com/devsoleo/heracles-webui) by providing access to in-game data from **World of Warcraft 3.3.5**, including:

- Creatures and NPCs
- Zones and subzones

This data comes from the AzerothCore emulator's world database for Creatures, NPCs, and Zones, while the Subzones are sourced from the game's client .dbc files.

## 🧩 Custom Content Support

Want to add your own creatures or areas?
You can modify the bundled `database.db` file to:

- Add custom NPCs, mobs 
- Insert zones or subzones from other WoW versions

The API is flexible enough to support **any flavor of WoW**, as long as the structure of `database.db` is respected.

## 🐳 Docker

```bash
docker pull registry.gitlab.com/devsoleo/heracles-api:latest
```

## 🙏 Acknowledgments

- 💖 [AzerothCore](https://github.com/azerothcore/azerothcore-wotlk) — for maintaining the emulator and DB structures we rely on