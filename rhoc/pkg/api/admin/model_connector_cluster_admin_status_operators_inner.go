/*
Connector Service Fleet Manager Admin APIs

Connector Service Fleet Manager Admin is a Rest API to manage connector clusters.

API version: 0.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package admin

// ConnectorClusterAdminStatusOperatorsInner struct for ConnectorClusterAdminStatusOperatorsInner
type ConnectorClusterAdminStatusOperatorsInner struct {
	Operator ConnectorOperator `json:"operator,omitempty"`
	// the namespace to which the operator has been installed
	Namespace string `json:"namespace,omitempty"`
	// the status of the operator
	Status string `json:"status,omitempty"`
}