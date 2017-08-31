package main

import (
	"strings"

	"k8s.io/client-go/pkg/api/v1"
)

// Endpoint is a summary of kubernetes endpoint
type Endpoint struct {
	Name    string
	Address string
	Port    int32
	Labels  []string
	RefName string
}

// NewEndpoint allows to create Endpoint
func NewEndpoint(name, address string, port int32, labels []string, refName string) Endpoint {
	return Endpoint{name, address, port, labels, refName}
}

func generateEntries(endpoint *v1.Endpoints) []Endpoint {
	var (
		eps     []Endpoint
		refName string
		labels  []string
	)

	for k, v := range endpoint.GetLabels() {
		switch k {
		case "k2c-singletags":
			for _, singleLabel := range strings.Split(v, "_") {
				labels = append(labels, singleLabel)
			}
		default:
			labels = append(labels, k+"="+v)
		}
	}

	for _, subset := range endpoint.Subsets {
		for _, addr := range subset.Addresses {
			if addr.TargetRef != nil {
				refName = addr.TargetRef.Name
			}
			for _, port := range subset.Ports {
				eps = append(eps, NewEndpoint(endpoint.Name, addr.IP, port.Port, labels, refName))
			}
		}
	}

	return eps
}

func (k2c *kube2consul) updateEndpoints(ep *v1.Endpoints) {
	endpoints := generateEntries(ep)
	for _, e := range endpoints {
		k2c.registerEndpoint(e)
	}

	k2c.removeDeletedEndpoints(ep.Name, endpoints)
}
