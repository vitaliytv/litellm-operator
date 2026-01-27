# API Reference

## Packages
- [litellm.litellm.ai/v1alpha1](#litellmlitellmaiv1alpha1)


## litellm.litellm.ai/v1alpha1

Package v1alpha1 contains API Schema definitions for the litellm v1alpha1 API group.

### Resource Types
- [LiteLLMInstance](#litellminstance)
- [LiteLLMInstanceList](#litellminstancelist)
- [Model](#model)
- [ModelList](#modellist)



#### ConnectionRef







_Appears in:_
- [ModelSpec](#modelspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `secretRef` _[SecretRef](#secretref)_ |  |  |  |
| `instanceRef` _[InstanceRef](#instanceref)_ |  |  |  |


#### DatabaseSecretKeys







_Appears in:_
- [DatabaseSecretRef](#databasesecretref)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `hostSecret` _string_ |  |  |  |
| `passwordSecret` _string_ |  |  |  |
| `usernameSecret` _string_ |  |  |  |
| `dbnameSecret` _string_ |  |  |  |


#### DatabaseSecretRef







_Appears in:_
- [LiteLLMInstanceSpec](#litellminstancespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `nameRef` _string_ |  |  |  |
| `keys` _[DatabaseSecretKeys](#databasesecretkeys)_ |  |  |  |


#### Gateway







_Appears in:_
- [LiteLLMInstanceSpec](#litellminstancespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `enabled` _boolean_ |  |  |  |
| `host` _string_ |  |  |  |


#### Ingress







_Appears in:_
- [LiteLLMInstanceSpec](#litellminstancespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `enabled` _boolean_ |  |  |  |
| `host` _string_ |  |  |  |


#### InitModelInstance



model instance used to create proxy server config map



_Appears in:_
- [LiteLLMInstanceSpec](#litellminstancespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `modelName` _string_ |  |  |  |
| `requiresAuth` _boolean_ |  |  |  |
| `identifier` _string_ |  |  |  |
| `modelCredentials` _[ModelCredentialSecretRef](#modelcredentialsecretref)_ |  |  |  |
| `liteLLMParams` _[LiteLLMParams](#litellmparams)_ |  |  |  |
| `additionalProperties` _[RawExtension](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#rawextension-runtime-pkg)_ |  |  |  |


#### InstanceRef







_Appears in:_
- [ConnectionRef](#connectionref)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `namespace` _string_ |  |  |  |
| `name` _string_ |  |  |  |


#### LiteLLMInstance



LiteLLMInstance is the Schema for the litellminstances API.



_Appears in:_
- [LiteLLMInstanceList](#litellminstancelist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `litellm.litellm.ai/v1alpha1` | | |
| `kind` _string_ | `LiteLLMInstance` | | |
| `kind` _string_ | Kind is a string value representing the REST resource this object represents.<br />Servers may infer this from the endpoint the client submits requests to.<br />Cannot be updated.<br />In CamelCase.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds |  |  |
| `apiVersion` _string_ | APIVersion defines the versioned schema of this representation of an object.<br />Servers should convert recognized schemas to the latest internal value, and<br />may reject unrecognized values.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources |  |  |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[LiteLLMInstanceSpec](#litellminstancespec)_ |  |  |  |
| `status` _[LiteLLMInstanceStatus](#litellminstancestatus)_ |  |  |  |


#### LiteLLMInstanceList



LiteLLMInstanceList contains a list of LiteLLMInstance.





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `litellm.litellm.ai/v1alpha1` | | |
| `kind` _string_ | `LiteLLMInstanceList` | | |
| `kind` _string_ | Kind is a string value representing the REST resource this object represents.<br />Servers may infer this from the endpoint the client submits requests to.<br />Cannot be updated.<br />In CamelCase.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds |  |  |
| `apiVersion` _string_ | APIVersion defines the versioned schema of this representation of an object.<br />Servers should convert recognized schemas to the latest internal value, and<br />may reject unrecognized values.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources |  |  |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[LiteLLMInstance](#litellminstance) array_ |  |  |  |


#### LiteLLMInstanceSpec



LiteLLMInstanceSpec defines the desired state of LiteLLMInstance.



_Appears in:_
- [LiteLLMInstance](#litellminstance)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `image` _string_ |  | ghcr.io/berriai/litellm-database:main-v1.74.9.rc.1 |  |
| `masterKey` _string_ |  |  |  |
| `databaseSecretRef` _[DatabaseSecretRef](#databasesecretref)_ |  |  |  |
| `redisSecretRef` _[RedisSecretRef](#redissecretref)_ |  |  |  |
| `ingress` _[Ingress](#ingress)_ |  |  |  |
| `gateway` _[Gateway](#gateway)_ |  |  |  |
| `replicas` _integer_ |  | 1 |  |
| `models` _[InitModelInstance](#initmodelinstance) array_ |  |  |  |
| `extraEnvVars` _[EnvVar](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#envvar-v1-core) array_ |  |  |  |


#### LiteLLMInstanceStatus



LiteLLMInstanceStatus defines the observed state of LiteLLMInstance.



_Appears in:_
- [LiteLLMInstance](#litellminstance)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `observedGeneration` _integer_ | ObservedGeneration represents the .metadata.generation that the condition was set based upon |  |  |
| `lastUpdated` _[Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#time-v1-meta)_ | LastUpdated represents the last time the status was updated |  |  |
| `configMapCreated` _boolean_ | Resource creation status |  |  |
| `secretCreated` _boolean_ |  |  |  |
| `deploymentCreated` _boolean_ |  |  |  |
| `serviceCreated` _boolean_ |  |  |  |
| `ingressCreated` _boolean_ |  |  |  |
| `conditions` _[Condition](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#condition-v1-meta) array_ | Conditions represent the latest available observations of a LiteLLM instance's state |  |  |


#### LiteLLMParams



LiteLLMParams defines the LiteLLM parameters for a model.



_Appears in:_
- [InitModelInstance](#initmodelinstance)
- [ModelSpec](#modelspec)
- [ModelStatus](#modelstatus)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `inputCostPerToken` _string_ | InputCostPerToken is the cost per input token |  |  |
| `outputCostPerToken` _string_ | OutputCostPerToken is the cost per output token |  |  |
| `inputCostPerSecond` _string_ | InputCostPerSecond is the cost per second for input |  |  |
| `outputCostPerSecond` _string_ | OutputCostPerSecond is the cost per second for output |  |  |
| `inputCostPerPixel` _string_ | InputCostPerPixel is the cost per pixel for input |  |  |
| `outputCostPerPixel` _string_ | OutputCostPerPixel is the cost per pixel for output |  |  |
| `apiKey` _string_ | APIKey is the API key for the model |  |  |
| `apiBase` _string_ | APIBase is the base URL for the API |  |  |
| `apiVersion` _string_ | APIVersion is the version of the API |  |  |
| `vertexProject` _string_ | VertexProject is the Google Cloud project for Vertex AI |  |  |
| `vertexLocation` _string_ | VertexLocation is the location for Vertex AI |  |  |
| `vertexCredentials` _string_ | VertexCredentials is the credentials for Vertex AI |  |  |
| `regionName` _string_ | RegionName is the region name for the service |  |  |
| `awsAccessKeyId` _string_ | AWSAccessKeyID is the AWS access key ID |  |  |
| `awsSecretAccessKey` _string_ | AWSSecretAccessKey is the AWS secret access key |  |  |
| `awsRegionName` _string_ | AWSRegionName is the AWS region name |  |  |
| `watsonxRegionName` _string_ | WatsonXRegionName is the WatsonX region name |  |  |
| `customLLMProvider` _string_ | CustomLLMProvider is the custom LLM provider |  |  |
| `tpm` _integer_ | TPM is tokens per minute |  |  |
| `rpm` _integer_ | RPM is requests per minute |  |  |
| `timeout` _integer_ | Timeout is the timeout in seconds |  |  |
| `streamTimeout` _integer_ | StreamTimeout is the stream timeout in seconds |  |  |
| `maxRetries` _integer_ | MaxRetries is the maximum number of retries |  |  |
| `organization` _string_ | Organization is the organization name |  |  |
| `configurableClientsideAuthParams` _[RawExtension](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#rawextension-runtime-pkg)_ | ConfigurableClientsideAuthParams are configurable client-side auth parameters |  |  |
| `litellmCredentialName` _string_ | LiteLLMCredentialName is the LiteLLM credential name |  |  |
| `litellmTraceId` _string_ | LiteLLMTraceID is the LiteLLM trace ID |  |  |
| `maxFileSizeMb` _integer_ | MaxFileSizeMB is the maximum file size in MB |  |  |
| `maxBudget` _string_ | MaxBudget is the maximum budget |  |  |
| `budgetDuration` _string_ | BudgetDuration is the budget duration |  |  |
| `useInPassThrough` _boolean_ | UseInPassThrough indicates if to use in pass through |  |  |
| `useLitellmProxy` _boolean_ | UseLiteLLMProxy indicates if to use LiteLLM proxy |  |  |
| `mergeReasoningContentInChoices` _boolean_ | MergeReasoningContentInChoices indicates if to merge reasoning content in choices |  |  |
| `modelInfo` _[RawExtension](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#rawextension-runtime-pkg)_ | ModelInfo contains additional model information |  |  |
| `mockResponse` _string_ | MockResponse is the mock response |  |  |
| `autoRouterConfigPath` _string_ | AutoRouterConfigPath is the auto router config path |  |  |
| `autoRouterConfig` _string_ | AutoRouterConfig is the auto router config |  |  |
| `autoRouterDefaultModel` _string_ | AutoRouterDefaultModel is the auto router default model |  |  |
| `autoRouterEmbeddingModel` _string_ | AutoRouterEmbeddingModel is the auto router embedding model |  |  |
| `model` _string_ | Model is the model name |  |  |
| `additionalProp1` _[RawExtension](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#rawextension-runtime-pkg)_ | AdditionalProps contains additional properties |  |  |


#### Model



Model is the Schema for the models API.



_Appears in:_
- [ModelList](#modellist)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `litellm.litellm.ai/v1alpha1` | | |
| `kind` _string_ | `Model` | | |
| `kind` _string_ | Kind is a string value representing the REST resource this object represents.<br />Servers may infer this from the endpoint the client submits requests to.<br />Cannot be updated.<br />In CamelCase.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds |  |  |
| `apiVersion` _string_ | APIVersion defines the versioned schema of this representation of an object.<br />Servers should convert recognized schemas to the latest internal value, and<br />may reject unrecognized values.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources |  |  |
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `spec` _[ModelSpec](#modelspec)_ |  |  |  |
| `status` _[ModelStatus](#modelstatus)_ |  |  |  |


#### ModelCredentialSecretKeys







_Appears in:_
- [ModelCredentialSecretRef](#modelcredentialsecretref)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiKey` _string_ |  |  |  |
| `apiBase` _string_ |  |  |  |
| `awsSecretAccessKey` _string_ |  |  |  |
| `awsAccessKeyId` _string_ |  |  |  |
| `vertexCredentials` _string_ |  |  |  |
| `vertexProject` _string_ |  |  |  |


#### ModelCredentialSecretRef







_Appears in:_
- [InitModelInstance](#initmodelinstance)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `nameRef` _string_ |  |  |  |
| `keys` _[ModelCredentialSecretKeys](#modelcredentialsecretkeys)_ |  |  |  |


#### ModelInfo



ModelInfo defines the model information.



_Appears in:_
- [ModelSpec](#modelspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `id` _string_ | ID is the model ID |  |  |
| `dbModel` _boolean_ | DBModel indicates if this is a database model |  |  |
| `teamId` _string_ | TeamID is the team ID |  |  |
| `teamPublicModelName` _string_ | TeamPublicModelName is the team public model name |  |  |
| `additionalProp1` _[RawExtension](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#rawextension-runtime-pkg)_ | AdditionalProps contains additional properties |  |  |


#### ModelList



ModelList contains a list of Model.





| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `apiVersion` _string_ | `litellm.litellm.ai/v1alpha1` | | |
| `kind` _string_ | `ModelList` | | |
| `kind` _string_ | Kind is a string value representing the REST resource this object represents.<br />Servers may infer this from the endpoint the client submits requests to.<br />Cannot be updated.<br />In CamelCase.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds |  |  |
| `apiVersion` _string_ | APIVersion defines the versioned schema of this representation of an object.<br />Servers should convert recognized schemas to the latest internal value, and<br />may reject unrecognized values.<br />More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources |  |  |
| `metadata` _[ListMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#listmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |  |  |
| `items` _[Model](#model) array_ |  |  |  |


#### ModelSpec



ModelSpec defines the desired state of Model.



_Appears in:_
- [Model](#model)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `connectionRef` _[ConnectionRef](#connectionref)_ | ConnectionRef is the connection reference |  |  |
| `modelName` _string_ | ModelName is the name of the model |  |  |
| `litellmParams` _[LiteLLMParams](#litellmparams)_ | LiteLLMParams contains the LiteLLM parameters |  |  |
| `modelInfo` _[ModelInfo](#modelinfo)_ | ModelInfo contains the model information |  |  |
| `modelSecretRef` _[SecretRef](#secretref)_ | ModelSecretRef is the model secret reference |  |  |


#### ModelStatus



ModelStatus defines the observed state of Model.



_Appears in:_
- [Model](#model)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `observedGeneration` _integer_ | ObservedGeneration represents the .metadata.generation that the condition was set based upon |  |  |
| `lastUpdated` _[Time](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#time-v1-meta)_ | LastUpdated represents the last time the status was updated |  |  |
| `conditions` _[Condition](https://kubernetes.io/docs/reference/generated/kubernetes-api/v/#condition-v1-meta) array_ | Conditions represent the latest available observations of a LiteLLM instance's state |  |  |
| `modelName` _string_ | ModelName is the name of the model |  |  |
| `litellmParams` _[LiteLLMParams](#litellmparams)_ | LiteLLMParams contains the LiteLLM parameters |  |  |
| `modelId` _string_ | ModelId contains the model uuid provided by litellm server |  |  |




#### RedisSecretKeys







_Appears in:_
- [RedisSecretRef](#redissecretref)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `hostSecret` _string_ |  |  |  |
| `portSecret` _string_ |  |  |  |
| `passwordSecret` _string_ |  |  |  |


#### RedisSecretRef







_Appears in:_
- [LiteLLMInstanceSpec](#litellminstancespec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `nameRef` _string_ |  |  |  |
| `keys` _[RedisSecretKeys](#redissecretkeys)_ |  |  |  |


#### SecretRef







_Appears in:_
- [ConnectionRef](#connectionref)
- [ModelSpec](#modelspec)

| Field | Description | Default | Validation |
| --- | --- | --- | --- |
| `namespace` _string_ |  |  |  |
| `secretName` _string_ |  |  |  |


