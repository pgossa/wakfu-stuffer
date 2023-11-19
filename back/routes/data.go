package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pgossa/wakfu-stuffer/handlers"
)

// TODO: Add finer routes
func AddDataRoutes(rg *gin.RouterGroup) {
	data := rg.Group("/data")
	data.GET("/items", handlers.StructWrapper(handlers.GetItems))
}

func AddDataCategoriesRoutes(rg *gin.RouterGroup) {
	dataCategories := rg.Group("/data/categories")
	dataCategories.GET("/equipments", handlers.StructWrapper(handlers.GetEquipments))
	dataCategories.GET("/consumables", handlers.StructWrapper(handlers.GetConsumables))
	dataCategories.GET("/resources", handlers.StructWrapper(handlers.GetResources))
	dataCategories.GET("/havenbags", handlers.StructWrapper(handlers.GetHavenBags))
	dataCategories.GET("/miscellanies", handlers.StructWrapper(handlers.GetMiscellanies))
	dataCategories.GET("/cosmetics", handlers.StructWrapper(handlers.GetCosmetics))
	dataCategories.GET("/questItems", handlers.StructWrapper(handlers.GetQuestItems))
	dataCategories.GET("/pets", handlers.StructWrapper(handlers.GetPets))
	dataCategories.GET("/mounts", handlers.StructWrapper(handlers.GetMounts))
	dataCategories.GET("/notSorted", handlers.StructWrapper(handlers.GetNotSorted))
}
