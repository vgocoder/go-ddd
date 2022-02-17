ROOT := $(realpath ../)

OPENAPI_CONFIG := $(ROOT)/script/restapi/openapi.json
OPENAPI_OUTPUT := $(ROOT)/interfaces/restapi
OPENAPI_PACKAGE := restapi

SQLBOILER_CONFIG := $(ROOT)/script/boiler/sqlboiler.toml
SQLBOILER_OUTPUT := $(ROOT)/infrastructure/persistence
SQLBOILER_PACKAGE := persistence

.PHONY: generate
generate: $(SQLBOILER_OUTPUT) \
					$(OPENAPI_OUTPUT) \
					$(OPENAPI_OUTPUT)/types.go \
					$(OPENAPI_OUTPUT)/server.go
$(SQLBOILER_OUTPUT):
	sqlboiler -o $(SQLBOILER_OUTPUT) -p $(SQLBOILER_PACKAGE) -c $(SQLBOILER_CONFIG) mysql
$(OPENAPI_OUTPUT):
	mkdir -p $(OPENAPI_OUTPUT)
$(OPENAPI_OUTPUT)/types.go:$(OPENAPI_CONFIG) $(OPENAPI_OUTPUT)
	oapi-codegen -generate types -package $(OPENAPI_PACKAGE) -o $@.tmp $<
	gomodifytags -all -add-tags boil -transform snakecase -file $@.tmp > $@
	rm -f $(OPENAPI_OUTPUT)/types.go.tmp
$(OPENAPI_OUTPUT)/server.go: $(OPENAPI_CONFIG)
	oapi-codegen -generate chi-server -package $(OPENAPI_PACKAGE) -o $@ $<

clean:
	rm -rf $(SQLBOILER_OUTPUT)
	rm -rf $(OPENAPI_OUTPUT)