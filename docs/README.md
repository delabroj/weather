# Development Setup

The documentation is compiled using [Swagger Docs](https://swagger.io/docs/).

## Generating Swagger 2.0 (aka OpenAPI 2.0) docs

Running this will update and validate the `swagger.yaml` file:
```
$ sh gendocs.sh
```

Requires:
```
go install github.com/go-swagger/go-swagger/cmd/swagger@latest
go install github.com/mikefarah/yq/v4@latest
```

Helpful documentation:
- https://goswagger.io/use/spec.html
- https://ldej.nl/post/generating-swagger-docs-from-go/
- https://github.com/ldej/swagger-go-example
