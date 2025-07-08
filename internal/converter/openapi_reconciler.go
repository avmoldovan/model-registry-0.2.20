package converter

import "github.com/kubeflow/model-registry/pkg/openapi"

// NOTE: methods must follow these patterns, otherwise tests could not find possible issues:
// Converters patch fields entity: UpdateExisting<ENTITY>

// goverter:converter
// goverter:output:file ./generated/openapi_reconciler.gen.go
// goverter:wrapErrors
// goverter:enum:unknown @error
// goverter:matchIgnoreCase
// goverter:useZeroValueOnPointerInconsistency
type OpenAPIReconciler interface {
	// Ignore all fields that can't be updated
	// goverter:default InitWithExisting
	// goverter:autoMap Update
	// goverter:ignore Id CreateTimeSinceEpoch LastUpdateTimeSinceEpoch Name Owner UserId
	UpdateExistingRegisteredModel(source OpenapiUpdateWrapper[openapi.RegisteredModel]) (openapi.RegisteredModel, error)

	// Ignore all fields that can't be updated
	// goverter:default InitWithExisting
	// goverter:autoMap Update
	// goverter:ignore Id CreateTimeSinceEpoch LastUpdateTimeSinceEpoch Name RegisteredModelId Owner UserId
	UpdateExistingModelVersion(source OpenapiUpdateWrapper[openapi.ModelVersion]) (openapi.ModelVersion, error)

	// Ignore all fields that can't be updated
	// goverter:default InitWithExisting
	// goverter:autoMap Update
	// goverter:ignore Id CreateTimeSinceEpoch LastUpdateTimeSinceEpoch Name ArtifactType Owner UserId
	UpdateExistingDocArtifact(source OpenapiUpdateWrapper[openapi.DocArtifact]) (openapi.DocArtifact, error)

	// Ignore all fields that can't be updated
	// goverter:default InitWithExisting
	// goverter:autoMap Update
	// goverter:ignore Id CreateTimeSinceEpoch LastUpdateTimeSinceEpoch Name ArtifactType Owner UserId
	UpdateExistingModelArtifact(source OpenapiUpdateWrapper[openapi.ModelArtifact]) (openapi.ModelArtifact, error)

	// Ignore all fields that can't be updated
	// goverter:default InitWithExisting
	// goverter:autoMap Update
	// goverter:ignore Id CreateTimeSinceEpoch LastUpdateTimeSinceEpoch Name Owner UserId
	UpdateExistingServingEnvironment(source OpenapiUpdateWrapper[openapi.ServingEnvironment]) (openapi.ServingEnvironment, error)

	// Ignore all fields that can't be updated
	// goverter:default InitWithExisting
	// goverter:autoMap Update
	// goverter:ignore Id CreateTimeSinceEpoch LastUpdateTimeSinceEpoch Name RegisteredModelId ServingEnvironmentId Owner UserId
	UpdateExistingInferenceService(source OpenapiUpdateWrapper[openapi.InferenceService]) (openapi.InferenceService, error)

	// Ignore all fields that can't be updated
	// goverter:default InitWithExisting
	// goverter:autoMap Update
	// goverter:ignore Id CreateTimeSinceEpoch LastUpdateTimeSinceEpoch Name ModelVersionId Owner UserId
	UpdateExistingServeModel(source OpenapiUpdateWrapper[openapi.ServeModel]) (openapi.ServeModel, error)
}
