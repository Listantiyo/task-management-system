swagger:
		swag init --dir cmd/server,internal/delivery/http,internal/utils/dto --parseDependency --parseInternal -g main.go --propertyStrategy pascalcase