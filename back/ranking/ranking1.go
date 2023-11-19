package ranking

import (
	"github.com/pgossa/wakfu-stuffer/types"
	"github.com/pgossa/wakfu-stuffer/utils"
)

// The idea for this algorithm is to take the top 10 non relic or epic items for each equipement position
// Then to pick the relics and epics
// And finally to try every possible combination
// There might be an issue with low number specs (such as AP, MP, ...)

func GetBetterBuild(request types.RequestRanking) any { //buildTypes.Build {
	itemList := utils.RemoveTooHighLevelItems(request.Level)
	// res := buildTypes.Build{}
	return itemList
}
