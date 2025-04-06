package services

func Generate(rawForms []map[string]interface{}) {

}

// type KillForm struct {
// 	Type     string `json:"type"`
// 	Amount   string `json:"amount"`
// 	Creature string `json:"creature"`
// }

// type GoForm struct {
// 	Type    string `json:"type"`
// 	Zone    string `json:"zone"`
// 	Subzone string `json:"subzone"`
// }

// type TargetForm struct {
// 	Type     string `json:"type"`
// 	Player   string `json:"player"`
// 	Creature string `json:"creature"`
// }

// func GenerateKey(c *gin.Context) {
// 	var rawForms []map[string]interface{}

// 	if err := c.ShouldBindJSON(&rawForms); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	for _, raw := range rawForms {
// 		if rawType, ok := raw["type"]; ok {

// 			jsonData, _ := json.Marshal(raw)

// 			switch rawType {
// 			case "K":
// 				var kForm KillForm
// 				json.Unmarshal(jsonData, &kForm)
// 				println(kForm.Type, kForm.Amount, kForm.Creature)
// 			case "G":
// 				var gForm GoForm
// 				json.Unmarshal(jsonData, &gForm)
// 				println(gForm.Type, gForm.Zone, gForm.Subzone)
// 			case "T":
// 				var tForm TargetForm
// 				json.Unmarshal(jsonData, &tForm)
// 				println(tForm.Type, tForm.Player, tForm.Creature)
// 			}

// 		}
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Formulaire reçu",
// 	})
// }
