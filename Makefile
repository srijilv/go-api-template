export MOD_PRFX := github.com/srijilv/go-api-template.git/pkg/interfaces
export OPENAPI_MOD_PRFX := ${MOD_PRFX}/openapi
export GRPC_MOD_PREFIX="github.com/srijilv/go-api-template.git/pkg/interfaces/grpc"


.PHONY: openapi

openapi:
	oapi-codegen  -generate types,skip-prune -o pkg/interfaces/openapi/common/openapi_types_common.gen.go -package common api/openapi/common.yml
	oapi-codegen  -generate types,skip-prune -o pkg/interfaces/openapi/list_books/openapi_types_list_books.gen.go --import-mapping=./common.yml:${OPENAPI_MOD_PRFX}/common -package listbooks api/openapi/list_books.yml

	oapi-codegen  -generate types,skip-prune -o pkg/interfaces/openapi/openapi_types_books.gen.go --import-mapping="./list_books.yml:${OPENAPI_MOD_PRFX}/list_books,./common.yml:${OPENAPI_MOD_PRFX}/common" -package openapi api/openapi/api.yml
	oapi-codegen  -generate chi-server -o pkg/interfaces/openapi/openapi_server.gen.go --import-mapping="./list_books.yml:${OPENAPI_MOD_PRFX}/list_books,./common.yml:${OPENAPI_MOD_PRFX}/common" -package openapi api/openapi/api.yml
	oapi-codegen  -generate client -o pkg/interfaces/openapi/openapi_client.gen.go --import-mapping="./list_books.yml:${OPENAPI_MOD_PRFX}/list_books,./common.yml:${OPENAPI_MOD_PRFX}/common" -package openapi api/openapi/api.yml



.PHONY: grpc
grpc:
	protoc --go_out=./pkg/interfaces/grpc --go_opt=module=${GRPC_MOD_PREFIX} \
	 --go-grpc_out=./pkg/interfaces/grpc --go-grpc_opt=module=${GRPC_MOD_PREFIX} \
	  api/grpc/books-common.proto
	protoc --go_out=./pkg/interfaces/grpc --go_opt=module=${GRPC_MOD_PREFIX} \
	 --go-grpc_out=./pkg/interfaces/grpc --go-grpc_opt=module=${GRPC_MOD_PREFIX} \
	  api/grpc/books.proto
	protoc --go_out=./pkg/interfaces/grpc --go_opt=module=${GRPC_MOD_PREFIX} \
	 --go-grpc_out=./pkg/interfaces/grpc --go-grpc_opt=module=${GRPC_MOD_PREFIX} \
	  api/grpc/list-books.proto