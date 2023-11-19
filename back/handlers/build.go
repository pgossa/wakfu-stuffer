package handlers

import (
	"encoding/json"
	"io"
	"strconv"

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
	build := ranking.GetBetterBuild(requestRanking)
	return build, nil
}

func GetLevelItems(c *gin.Context) (any, error) {
	level, err := strconv.Atoi(c.Param("level"))
	if err != nil {
		return nil, err
	}
	requestRanking := types.RequestRanking{Level: level}
	return ranking.GetBetterBuild(requestRanking), nil
}
