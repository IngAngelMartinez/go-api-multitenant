package models

type TenantConfiguration struct {
	TenantId             string
	MongoDBConfiguration *MongoDBConfiguration
}

type MongoDBConfiguration struct {
	ConnectionString string
	DatabaseName     string
}
