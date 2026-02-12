package kubernetes

import karpenterv1 "sigs.k8s.io/karpenter/pkg/apis/v1"

type ClusterAutoscalingValuesList struct {
	Values []ClusterAutoscalingValues `json:"values"`
}

type ClusterAutoscalingValues struct {
	TargetName string `json:"target_name"`
	TargetHash string `json:"target_hash"`
	Specs      Specs  `json:"specs"`

	// These fields are deprecated and will be removed in a future version
	Name                     string         `json:"name"`
	Labels                   []DomainLabels `json:"labels"`
	Taints                   []Taints       `json:"taints"`
	RecommendedInstanceTypes []string       `json:"recommended_instance_types"`
}

type Specs struct {
	KarpenterV1 karpenterv1.NodePoolSpec `json:"karpenterv1"`
}

type DomainLabels struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Taints struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Effect string `json:"effect"`
}
