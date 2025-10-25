// package jsonschema embeds the JSON schema for external usage
package jsonschema

import (
	_ "embed"
)

//go:embed WorkloadValuesList.json
var KubernetesAutoscalingWorkloadValuesList []byte

//go:embed ClusterAutoscalingValuesList.json
var KubernetesAutoscalingClusterAutoscalingValuesList []byte
