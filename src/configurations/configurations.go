package configurations

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/IngAngelMartinez/go-api-multitenant/src/models"
	"github.com/gofiber/fiber/v2"
)

func GetTenantConfigurations() []models.TenantConfiguration {

	var tenantConfigurations []models.TenantConfiguration

	fileContent, err := os.Open("configurations.json")

	if err != nil {

		log.Fatal(err)

		return tenantConfigurations
	}

	defer fileContent.Close()

	byteResult, err := ioutil.ReadAll(fileContent)

	if err != nil {

		log.Fatal(tenantConfigurations)

		return tenantConfigurations
	}

	json.Unmarshal(byteResult, &tenantConfigurations)

	return tenantConfigurations
}

func GetServerConfigurations() fiber.Config {

	return fiber.Config{
		Immutable: true,
	}
}
