package types

// Endpoint Structure
type Endpoint struct {
	Protocol string `json:"protocol" bson:"protocol"`
	IP       string `json:"ip" bson:"ip"`
	Port     int    `json:"port" bson:"port"`
}

// K8sEndpoint Structure
type K8sEndpoint struct {
	MicroserviceName string `json:"microservice_name,omitempty" bson:"microservice_name,omitempty"`
	EndpointName     string `json:"endpoint_name,omitempty" bson:"endpoint_name,omitempty"`

	Labels []string `json:"labels,omitempty" bson:"labels,omitempty"`

	Endpoints []Endpoint `json:"mappings" bson:"mappings"`
}

// K8sService Structure
type K8sService struct {
	MicroserviceName string `json:"microservice_name,omitempty" bson:"microservice_name,omitempty"`
	ServiceName      string `json:"service_name,omitempty" bson:"service_name,omitempty"`

	Labels []string `json:"labels,omitempty" bson:"labels,omitempty"`

	Type      string `json:"type,omitempty" bson:"type,omitempty"`
	Protocol  string `json:"protocol,omitempty" bson:"protocol,omitempty"`
	ClusterIP string `json:"cluster_ip,omitempty" bson:"cluster_ip,omitempty"`

	ServicePort   int `json:"service_port" bson:"service_port"`
	NodePort      int `json:"node_port" bson:"node_port"`
	ContainerPort int `json:"container_port" bson:"container_port"`

	Selector map[string]string `json:"selector" bson:"selector"`
}