package gameMap

var attrMap = map[string]string{"v": "v", "A": "atk", "B": "def",
	"C": "hp",
	"D": "perHit",
	"E": "perDodge",
	"F": "perCrit",
	"G": "critHurt",
	"H": "pvpHurt",
	"I": "pvpHarmless",
	"J": "tAtk",
	"K": "tDef",
	"R": "hpPer",
	"M": "atkPer",
	"N": "defPer",
	"t": "t"}

func FormatAttrItem(attr map[string]interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	for k, v := range attr {
		field := attrMap[k]
		res[field] = v
	}
	return res

}
