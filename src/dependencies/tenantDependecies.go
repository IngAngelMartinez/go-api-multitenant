package dependencies

import (
	"errors"
	"sync"

	"github.com/IngAngelMartinez/go-api-multitenant/src/configurations"
	"github.com/IngAngelMartinez/go-api-multitenant/src/database"
	"github.com/IngAngelMartinez/go-api-multitenant/src/models"
	"github.com/IngAngelMartinez/go-api-multitenant/src/repository"
	"github.com/IngAngelMartinez/go-api-multitenant/src/services"
)

var (
	allTenantDependencies map[string]*dependencies
	once                  sync.Once
)

type dependencies struct {
	TenantId         string
	DocumentsService *services.DocumentsService
}

func GetTenantDependencies(tenantId string) (*dependencies, error) {

	once.Do(func() {

		allTenantDependencies = make(map[string]*dependencies)
	})

	tenantDependecies, exists := allTenantDependencies[tenantId]

	if exists {

		return tenantDependecies, nil
	}

	return createDependencies(tenantId)
}

func createDependencies(tenantId string) (*dependencies, error) {

	tenantConfigurations := getConfigurations(tenantId)

	if tenantConfigurations == (models.TenantConfiguration{}) {

		return nil, errors.New("Not found configuration")
	}

	mongoTenant, err := getMongoTenant(tenantConfigurations.MongoDBConfiguration)

	if err != nil {

		return nil, err
	}

	dependencies := dependencies{
		TenantId:         tenantId,
		DocumentsService: documentsDependencies(mongoTenant),
	}

	allTenantDependencies[tenantId] = &dependencies

	return &dependencies, nil
}

func documentsDependencies(mongoTenant *database.MongoDB) *services.DocumentsService {

	collection := mongoTenant.Database.Collection("a")

	repository := repository.DocumentsRepository{Collection: collection}

	service := services.DocumentsService{DocumentsRepository: &repository}

	return &service
}

func getMongoTenant(mongoDBConfiguration *models.MongoDBConfiguration) (*database.MongoDB, error) {

	mongoTenant := database.MongoDB{
		Uri:          mongoDBConfiguration.ConnectionString,
		DatabaseName: mongoDBConfiguration.DatabaseName,
	}

	if err := mongoTenant.Connect(); err != nil {

		return nil, err
	}

	return &mongoTenant, nil
}

func getConfigurations(tenantId string) models.TenantConfiguration {

	tenantConfigurations := configurations.GetTenantConfigurations()

	for i := 0; i < len(tenantConfigurations); i++ {

		if tenantId == tenantConfigurations[i].TenantId {

			return tenantConfigurations[i]
		}
	}

	return models.TenantConfiguration{}
}
