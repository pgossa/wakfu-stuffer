package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/pgossa/wakfu-stuffer/type/wakfuType"
)

func CheckForNewVersion() {
	log.Println(wakfuType.WVersionData.Version)
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
	}
}

func isANewVersion(newVersion *wakfuType.WVersion) bool {
	return (*newVersion != wakfuType.WVersionData)
}

func setNewVersion(newVersion *wakfuType.WVersion) {
	wakfuType.WVersionData = *newVersion
}

func retrieveVersion() (*wakfuType.WVersion, error) {
	var version wakfuType.WVersion
	body, err := fetchData("https://wakfu.cdn.ankama.com/gamedata/config.json")
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
	wakfuType.WItemData = items
	actions, err := retrieveActions()
	if err != nil {
		return err
	}
	wakfuType.WActionData = actions
	itemProperties, err := retrieveItemProperties()
	if err != nil {
		return err
	}
	wakfuType.WItemPropertiesData = itemProperties
	itemTypes, err := retrieveItemTypes()
	if err != nil {
		return err
	}
	wakfuType.WItemTypesData = itemTypes
	return nil
}

func retrieveItems() ([]wakfuType.WItem, error) {
	var items []wakfuType.WItem
	body, err := fetchData(createUrl("items"))
	err = json.Unmarshal(body, &items)
	if err != nil {
		log.Println("error:", err)
		return nil, err
	}
	return items, err
}

func retrieveActions() ([]wakfuType.WAction, error) {
	var actions []wakfuType.WAction
	body, err := fetchData(createUrl("actions"))
	err = json.Unmarshal(body, &actions)
	if err != nil {
		log.Println("error:", err)
		return nil, err
	}
	return actions, err
}

func retrieveItemProperties() ([]wakfuType.WItemProperties, error) {
	var itemProperties []wakfuType.WItemProperties
	body, err := fetchData(createUrl("itemProperties"))
	err = json.Unmarshal(body, &itemProperties)
	if err != nil {
		log.Println("error:", err)
		return nil, err
	}
	return itemProperties, err
}

func retrieveItemTypes() ([]wakfuType.WItemTypes, error) {
	var itemTypes []wakfuType.WItemTypes
	body, err := fetchData(createUrl("itemTypes"))
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
	return baseUrl + wakfuType.WVersionData.Version + "/" + path + ".json"
}
