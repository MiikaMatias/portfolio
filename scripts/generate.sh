cd ./go
go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=api/config.yaml ../api/spec.yaml
echo "Generated Golang types"

cd ../web
npx openapi-typescript ./../api/spec.yaml -o ./src/api/api.d.ts
echo "Generated TS types"