// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xWzY4bRRB+lVHDASSzdggcmBsgDrkgBMoJ5TDYvd4OO9OT7h7QKrLkHxSCiFjEGQS8",
	"wcTx4FEcj1/hqzdC1WPseD1eEEl8ysVq11RVf/VV1ae+L7o6TnUiE2dFeF/Y7pmMI3/8MHNnt600n8t7",
	"mbSOTanRqTROSe8g40id8+FUmzhyIlxbWsJdpFKEwjqjkr4YtEQaWfutNr0d741xL2DQEkbey5SRPRF+",
	"uUm7CbizidBf3ZVdx1ds4dpUJ1bu43X6a5n4w7W31W5NV3xsZOTky+LkVBnrPo1i2QDpRRnb5m79N/ae",
	"L+0Qf2oXTJapZiB7yT8xRl+TN5bWRn3Z3JkrydikklPtnZU752/4FRVmNMGCHgXIMaUxKhqixBw5PUCJ",
	"MsAMC7oM6HvkNEYuWuIbaazSiQhF56RzcoNx6lQmUapEKG6edE5ues7cmcfYtqqfvKP8+KS67nxP2q5R",
	"qauz4E9UmGOGHEuUKAIsUWGKCk9QBVihYnyYeyODKDwi/hCgxDMUHBegDP5pGHMUcfZbPRGKz7R1X6h+",
	"cisRda+ldR/p3gUj6erEycSDitL0XHV9WPuu1cl2qfn0ppGnIhRvtLdb316vfHt/uge7Y+VMJr2hbqTn",
	"5d3OjVcCYD0rHsEVnn9vZPJRQBMaYYWCHjLzAY22/eDmvtfpvDSkuwPdBPI3HoApcixQYoYl/ciGejjW",
	"f+bIsaIhKhohrxF+cESEh2nEXygCGtGEfkBBIxpjShMUNGaQ7x+Vxl+wpAmNacjbQZe8LxU9RInHeIqc",
	"e1zQEFP/m7cCVHjMJxqj5IJoxIrwlEVhhmcoecO4oLVPRUOaeJWxWRxH5qJxjQ9t7lse1ZPnMj5ASZdv",
	"syZGfcsyjD+aPMQdvrJWlCy9RlF+5tp95VzPdyi56DrHvwiKlxDWEqyQ84xhQT8dlJTb6SuSlKtPiCML",
	"yt6T4LWcvJaTI8vJ/9rhF5KWweDvAAAA//9RZYNE1QsAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
