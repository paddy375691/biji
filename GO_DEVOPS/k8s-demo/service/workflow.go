package service

var Workflow workflow

type workflow struct {}

type WorkflowCreate struct {
	Name           string  `json:"name"`
	Namespace      string  `json:"namespace"`
	Replicas       int32   `json:"replicas"`
	Image          string  `json:"image"`
	Label          map[string]string  `json:"label"`
	Cpu            string  `json:"cpu"`
	Memory         string  `json:"memory"`
	ContainerPort  int32   `json:"container_port"`
	HealthCheck    bool    `json:"health_check"`
	HealthPath     string  `json:"health_path"`
	Type           string  `json:"type"`
	Port           int32   `json:"port"`
	NodePort       int32   `json:"node_port"`
	Hosts          map[string][]*HttpPath `json:"hosts"`
}

func(w *workflow) CreateWorkflow(data *WorkflowCreate) (err error) {
	var serviceType string

	dc := &DeployCreate{
		Name:          data.Name,
		Namespace:     data.Namespace,
		Replicas:      data.Replicas,
		Image:         data.Image,
		Label:         data.Label,
		Cpu:           data.Cpu,
		Memory:        data.Memory,
		ContainerPort: data.ContainerPort,
		HealthCheck:   data.HealthCheck,
		HealthPath:    data.HealthPath,
	}
	err = Deployment.CreateDeployment(dc)
	if err != nil {
		return err
	}

	if data.Type != "Ingress" {
		serviceType = data.Type
	} else {
		serviceType = "ClusterIP"
	}
	sc := &ServiceCreate{
		Name:          data.Name,
		Namespace:     data.Namespace,
		Type:          serviceType,
		ContainerPort: data.ContainerPort,
		Port:          data.Port,
		NodePort:      data.NodePort,
		Label:         data.Label,
	}
	err = Servicev1.CreateService(sc)
	if err != nil {
		return err
	}

	if data.Type == "Ingress" {
		ic := &IngressCreate{
			Name:      data.Name,
			Namespace: data.Namespace,
			Label:     data.Label,
			Hosts:     data.Hosts,
		}
		err = Ingress.CreateIngress(ic)
		if err != nil {
			return err
		}
	}

	return nil
}