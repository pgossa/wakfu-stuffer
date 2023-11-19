package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/pgossa/wakfu-stuffer/types/wakfuTypes"
)

func CheckForNewVersion() {
	log.Println(wakfuTypes.WVersionData.Version)
	newVersion, err := retrieveVersion()
	if err != nil {
		log.Println("The data may not be up to date, error while fetching wakfu API: ", err)
		return
	}
	if isANewVersion(newVersion) {
		log.Println("New version detected, starting updating data!")
		setNewVersion(newVersion)
		err = retrieveAndSetNewData()
		if err != nil {
			log.Println("The data may not be up to date, error while fetching wakfu API: ", err)
		}
		log.Println("Data successfully updated!")
		if err = ConvertItemListToCustomType(); err != nil {
			log.Println("Error during conversion", err)
		}
		if err = SortData(); err != nil {
			log.Println("Error while sorting the data:", err)
		}
		log.Println("Data successfully sorted!")
	}
}

func isANewVersion(newVersion *wakfuTypes.WVersion) bool {
	return (*newVersion != wakfuTypes.WVersionData)
}

func setNewVersion(newVersion *wakfuTypes.WVersion) {
	wakfuTypes.WVersionData = *newVersion
}

func retrieveVersion() (*wakfuTypes.WVersion, error) {
	var version wakfuTypes.WVersion
	body, err := fetchData("https://wakfu.cdn.ankama.com/gamedata/config.json")
	if err != nil {
		log.Println("Error while getting the version")
		return nil, err
	}
	err = json.Unmarshal(body, &version)
	if err != nil {
		log.Println("error:", err)
		return nil, err
	}
	return &version, err
}

func retrieveAndSetNewData() error {
	items, err := retrieveItems()
	if err != nil {
		return err
	}
	wakfuTypes.WItemData = items
	actions, err := retrieveActions()
	if err != nil {
		return err
	}
	wakfuTypes.WActionData = actions
	itemProperties, err := retrieveItemProperties()
	if err != nil {
		return err
	}
	wakfuTypes.WItemPropertiesData = itemProperties
	itemTypes, err := retrieveItemTypes()
	if err != nil {
		return err
	}
	wakfuTypes.WItemTypesData = itemTypes
	return nil
}

func retrieveItems() ([]wakfuTypes.WItem, error) {
	var items []wakfuTypes.WItem
	body, err := fetchData(createUrl("items"))
	if err != nil {
		log.Println("Error while getting the items")
		return nil, err
	}
	err = json.Unmarshal(body, &items)
	if err != nil {
		log.Println("error:", err)
		return nil, err
	}
	return items, err
}

func retrieveActions() ([]wakfuTypes.WAction, error) {
	var actions []wakfuTypes.WAction
	body, err := fetchData(createUrl("actions"))
	if err != nil {
		log.Println("Error while getting the actions")
		return nil, err
	}
	err = json.Unmarshal(body, &actions)
	if err != nil {
		log.Println("error:", err)
		return nil, err
	}
	return actions, err
}

func retrieveItemProperties() ([]wakfuTypes.WItemProperties, error) {
	var itemProperties []wakfuTypes.WItemProperties
	body, err := fetchData(createUrl("itemProperties"))
	if err != nil {
		log.Println("Error while getting the itemProperties")
		return nil, err
	}
	err = json.Unmarshal(body, &itemProperties)
	if err != nil {
		log.Println("error:", err)
		return nil, err
	}
	return itemProperties, err
}

func retrieveItemTypes() ([]wakfuTypes.WItemTypes, error) {
	var itemTypes []wakfuTypes.WItemTypes
	body, err := fetchData(createUrl("itemTypes"))
	if err != nil {
		log.Println("Error while getting the itemTypes")
		return nil, err
	}
	err = json.Unmarshal(body, &itemTypes)
	if err != nil {
		log.Println("error:", err)
		return nil, err
	}
	return itemTypes, err
}

func fetchData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("error:", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error:", err)
		return nil, err
	}
	return body, err
}

func createUrl(path string) string {
	baseUrl := "https://wakfu.cdn.ankama.com/gamedata/"
	return baseUrl + wakfuTypes.WVersionData.Version + "/" + path + ".json"
}
