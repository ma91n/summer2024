// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * openapi 3.0 auth code generator sample
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 */

package openapi

type Hello struct {
	Message string `json:"message,omitempty"`
}

// AssertHelloRequired checks if the required fields are not zero-ed
func AssertHelloRequired(obj Hello) error {
	return nil
}

// AssertHelloConstraints checks if the values respects the defined constraints
func AssertHelloConstraints(obj Hello) error {
	return nil
}
