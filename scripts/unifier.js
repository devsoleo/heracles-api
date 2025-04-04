const fs = require("fs")

const data = require("data.json")

const creatures = data.filter((item, index, self) => {
  // Crée une clé unique pour chaque combinaison de locale et name
  const uniqueKey = `${item["t.local"]}-${item["t.name"]}`;

  // Vérifie si une combinaison locale-name a déjà été rencontrée
  return self.findIndex(i => `${i["t.local"]}-${i["t.name"]}` === uniqueKey) === index;
});

fs.writeFileSync('./output.json', JSON.stringify(creatures, null, 2));