package handlers

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/pgossa/wakfu-stuffer/ranking"
	"github.com/pgossa/wakfu-stuffer/types"
)

func GetBetterBuild(c *gin.Context) (any, error) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return "", err
	}
	requestRanking := types.RequestRanking{}
	err = json.Unmarshal(jsonData, &requestRanking)
	if err != nil {
		return "", err
	}
	// return requestRanking, nil
	build := ranking.WeightRankBuild(requestRanking)
	return build, nil
}
