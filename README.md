# example-openapi

## API Docs

- https://swagger.io/specification/
- https://raw.githubusercontent.com/openapitools/openapi-generator/master/modules/openapi-generator/src/test/resources/3_0/petstore.yaml

## Tools

- https://github.com/ogen-go/ogen
- https://github.com/deepmap/oapi-codegen
- https://github.com/OpenAPITools/openapi-generator
- https://github.com/go-swagger/go-swagger

### ogen

- https://ogen.dev/docs/intro/

```
$ go mod init github.com/achiku/example-openapi/ogen
$ go install -v github.com/ogen-go/ogen/cmd/ogen@latest
```

```
$ ogen ../petstore.yml
```

### oapi-codegen

- https://github.com/deepmap/oapi-codegen

```
$ go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest
```
