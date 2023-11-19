package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pgossa/wakfu-stuffer/types/customTypes"
)

// All items
func GetItems(c *gin.Context) (any, error) {
	return customTypes.CustomItemsData, nil
}

// Categories
func GetEquipments(c *gin.Context) (any, error) {
	return customTypes.Equipments, nil
}
func GetConsumables(c *gin.Context) (any, error) {
	return customTypes.Consumables, nil
}
func GetResources(c *gin.Context) (any, error) {
	return customTypes.Resources, nil
}
func GetHavenBags(c *gin.Context) (any, error) {
	return customTypes.HavenBags, nil
}
func GetMiscellanies(c *gin.Context) (any, error) {
	return customTypes.Miscellanies, nil
}
func GetCosmetics(c *gin.Context) (any, error) {
	return customTypes.Cosmetics, nil
}
func GetQuestItems(c *gin.Context) (any, error) {
	return customTypes.QuestItems, nil
}
func GetPets(c *gin.Context) (any, error) {
	return customTypes.Pets, nil
}
func GetMounts(c *gin.Context) (any, error) {
	return customTypes.Mounts, nil
}
func GetNotSorted(c *gin.Context) (any, error) {
	return customTypes.NotSorted, nil
}
