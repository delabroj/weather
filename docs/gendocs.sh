set -e

OUT_FILE=swagger.yaml
echo "Generating docs/$OUT_FILE file"
SWAGGER_GENERATE_EXTENSION=false swagger generate spec -o "$OUT_FILE" --scan-models ../cmd/weather/main.go
yq -i -P 'sort_keys(..)' $OUT_FILE
echo "Finished generating docs/$OUT_FILE file"
echo

echo "Validating docs/$OUT_FILE file"
# Swagger 2.0 does not allow examples for parameters but postman supports it and it is useful to have.
# This step removes all examples so that we can still validate against the Swagger 2.0 spec.
echo "Creating duplicate swagger file with no examples"
yq 'del(.. | select(has("example")).example)' $OUT_FILE > swagger.no-examples.yaml
swagger validate --skip-warnings -q swagger.no-examples.yaml
rm swagger.no-examples.yaml
echo "Finished validating docs/$OUT_FILE file"
echo
