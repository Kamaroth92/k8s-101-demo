package main

import (
	demo "github.com/saschagrunert/demo"
)

func main() {
	d := demo.New()
	d.Name = "k8s-101-demo"
	d.Description = "A kubernetes 101 demo"
	d.Add(info(), "info", "Introduction")
	d.Add(docker(), "docker", "Docker")
	d.Add(microservices(), "microservices", "Microservices")
	d.Add(kubernetesOverview(), "kubernetes-overview", "Kubernetes Overview")
	d.Add(kubernetesResources(), "kubernetes-resources", "Kubernetes Resources")
	d.Add(kubernetesHelmOperators(), "kubernetes-helm-operators", "Helm and Operators")
	d.Add(kubernetesDemo(), "kubernetes-demo", "Kubernetes Microservices Demo")
	d.Add(kubernetesO11Y(), "kubernetes-011y", "Kubernetes O11Y (Observability)")
	d.Add(kubernetesSecurity(), "kubernetes-security", "Kubernetes Security")
	d.Add(EKSOverview(), "eks-overview", "EKS Overview")
	d.Run()
}

func info() *demo.Run {
	r := demo.NewRun(
		"Introduction",
		"- What are containers?",
		"",
		"- What are microservices?",
		"",
		"- An overview of Kubernetes",
		"",
		"- Types of resources",
		"",
		"- Helm, Kustomize, and Operators",
		"",
		"- Deployment demo",
		"",
		"- Logging and o11y",
		"",
		"- EKS: Overview",
	)

	r.Step(nil, nil)
	return r
}

func docker() *demo.Run {
	r := demo.NewRun(
		"Docker and containers",
		"- Containers are self contained, runnable, packages that typically contain everything that a service or application requires to run.",
		"",
		"- Ideally because containers include everything they need to run they should be portable between systems however this is not always the case due to potential host-level requirements.",
		"",
		"- Images are comprised of layers, and layers that are shared between images can be cached meaning smartly stacked images can be changed a re-generated very quickly.",
	)

	r.Step(
		demo.S(
			"List images",
		),
		demo.S(
			"docker image ls",
		),
	)

	r.Step(
		demo.S(
			"Pull images",
		),
		demo.S(
			"docker pull busybox",
		),
	)

	r.Step(
		demo.S(
			"Create our own images",
		),
		demo.S(
			"cat notes/dockerfile-instructions",
		),
	)

	r.Step(
		nil,
		demo.S(
			"cat python-demo/*",
		),
	)

	r.Step(
		nil,
		demo.S(
			"docker build python-demo -t tane/python-demo",
		),
	)

	r.Step(
		nil,
		demo.S(
			"docker image ls",
		),
	)

	r.Step(
		demo.S(
			"Run our image in a container",
		),
		demo.S(
			"docker run --rm tane/python-demo",
		),
	)

	r.Step(nil, nil)
	r.Cleanup(cleanupDocker)
	return r
}

func cleanupDocker() error {
	return demo.Ensure(
		"docker image rm tane/python-demo | true",
	)
}

func microservices() *demo.Run {

	r := demo.NewRun(
		"What are microservices?",
		"- Microservices are a software architecture pattern where applications built up by a collection of small independent services.",
		"",
		"- Each service typically fulfils one job or capability, and communicate with other services via an API",
		"",
		"- Services should be loosely coupled so that each can be worked on and updated without concern of effecting other services.",
	)

	r.Step(
		demo.S(
			"An example collection of microservices where each microservice is responsible for a specific task, but the entire collection represents an e-commerce platform",
		),
		demo.S(
			"cat notes/microservices-example",
		),
	)

	r.Step(nil, nil)
	return r
}

func kubernetesOverview() *demo.Run {
	r := demo.NewRun(
		"Kubernetes - An Overview",
		"- An open-source platform for deploying, scaling, and managing containerized applications.",
		"",
		"- Supports both imperative and declarative management.",
		"",
		"- Replace failed containers and manages resources.",
		"",
		"- Automatic scaling of workloads based on application demand.",
		"",
		"- Rollout updates to applications to prevent downtime.",
		"",
		"- Platform agnostic, can run on private or public cloud or bare-metal hardware.",
	)

	r.Step(
		nil,
		demo.S(
			"cat notes/kubernetes-components",
		),
	)

	r.Step(
		demo.S(
			"Get information about the running cluster",
		),
		demo.S(
			"kubectl cluster-info",
		),
	)

	r.Step(
		nil,
		demo.S(
			"kubectl get nodes",
		),
	)

	r.Step(nil, nil)
	return r
}

func kubernetesResources() *demo.Run {
	r := demo.NewRun(
		"Kubernetes - Resources",
		"- There are a variety of resources that kubernetes offers by default which typically cover 90% of the use cases",
		"",
		"- Custom resources can be offered via Custom Resource Definitions (CRDs) to expand the capabilities of your cluster",
	)

	r.Step(
		demo.S(
			"Common resources",
		),
		demo.S(
			"cat notes/kubernetes-resources",
		),
	)

	r.Step(
		demo.S(
			"List all available resources types on a cluster",
		),
		demo.S(
			"kubectl api-resources",
		),
	)

	r.Step(
		nil,
		demo.S(
			"cat notes/kubernetes-nginx-pod",
		),
	)

	r.Step(
		nil,
		demo.S(
			"cat kubernetes-nginx/kubernetes-nginx-deployment.yaml",
		),
	)

	r.Step(
		nil,
		demo.S(
			"cat kubernetes-nginx/kubernetes-nginx-service.yaml",
		),
	)

	r.Step(
		nil,
		demo.S(
			"cat kubernetes-nginx/kubernetes-nginx-ingress.yaml",
		),
	)

	r.Step(
		nil,
		demo.S(
			"kubectl apply -f kubernetes-nginx",
		),
	)

	r.Step(
		demo.S(
			"Cleanup resources",
		),
		demo.S(
			"kubectl delete -f kubernetes-nginx",
		),
	)

	r.Step(nil, nil)
	return r
}

func kubernetesHelmOperators() *demo.Run {
	r := demo.NewRun(
		"Helm, Kustomize, and Operators",
		"- Helm and Kustomize are very common ways to deploy to your kubernetes clusters",
		"",
		"- Simplifies application deployment",
		"",
		"- Allows templating and parameterization of your kubernetes manifests",
		"",
		"- Operators allow you to manage complex applications by automating operational tasks",
		"",
		"- Enables flexability and managability of complex lifecycle tasks",
	)

	r.Step(
		nil,
		demo.S(
			"cat notes/helm",
		),
	)

	r.Step(
		demo.S(
			"Add a helm repository",
		),
		demo.S(
			"helm repo add podinfo https://stefanprodan.github.io/podinfo",
			"&& helm repo list",
		),
	)

	r.Step(
		demo.S(
			"Install a helm chart",
		),
		demo.S(
			"helm upgrade --install --wait frontend",
			"--set replicaCount=2",
			"--set backend=http://backend-podinfo:9898/echo",
			"podinfo/podinfo",
		),
	)

	r.Step(
		demo.S(
			"Uninstall a helm chart",
		),
		demo.S(
			"helm uninstall frontend",
		),
	)

	r.Cleanup(cleanupHelm)
	r.Step(nil, nil)
	return r
}

func cleanupHelm() error {
	return demo.Ensure(
		"helm repo rm podinfo | true",
	)
}

func kubernetesDemo() *demo.Run {

	r := demo.NewRun(
		"Kubernetes - GCP Microservices Demo",
		"- This is a demo microservice project provided by Google for use with GCP, but it runs just fine on any cluster.",
	)

	r.Step(
		demo.S(
			"Taken from github.com/GoogleCloudPlatform/microservices-demo",
		),
		demo.S(
			"kubectl apply -f microservices-demo/release/kubernetes-manifests.yaml",
		),
	)

	r.Step(
		demo.S(
			"Forward a port to the service to allow us to view the frontend via K9S.",
			"port-forward svc/frontend-external 8080:80",
		),
		nil,
	)

	r.Step(
		demo.S(
			"Cleanup demo resources",
		),
		demo.S(
			"kubectl delete -f microservices-demo/release/kubernetes-manifests.yaml",
		),
	)

	r.Step(nil, nil)
	return r
}

func kubernetesO11Y() *demo.Run {

	r := demo.NewRun(
		"Kubernetes - O11Y",
		"- Crucial for ensuring the ongoing health of not just the cluster but the residing applications",
		"",
		"- Logging involves scraping logs from either container files or stdout with a logging agent and passing them to a logging system",
		"",
		"- Observability as whole encompasses logging, monitoring, and tracing, to ensure system performance",
	)

	r.Step(
		nil,
		demo.S(
			"cat notes/logging-diagram",
		),
	)

	r.Step(nil, nil)
	return r
}

func kubernetesSecurity() *demo.Run {

	r := demo.NewRun(
		"Kubernetes - Security",
		"- Role-Based Access Control (RBAC) is a method of regulating access to cluster resources",
		"",
		"- Roles define sets of permissions that then get bound to service accounts via role bindings",
		"",
		"- Helps enforce least-privileged principals",
	)

	r.Step(
		nil,
		demo.S(
			"cat notes/rbac-sa",
		),
	)

	r.Step(
		nil,
		demo.S(
			"cat notes/namespace",
		),
	)

	r.Step(nil, nil)
	return r
}
func EKSOverview() *demo.Run {

	r := demo.NewRun(
		"Elastic Kubernetes Service (EKS) - Overview",
		"- Managed Kubernetes service provided by AWS.",
		"",
		"- Simplifies the deployment, management, and scaling of Kubernetes clusters on AWS infrastructure.",
		"",
		"- Offers high availability, security, and scalability out of the box.",
		"",
		"- Integrates with other AWS services like IAM, VPC, and CloudWatch.",
		"",
		"- Can use managed or self-managed nodes",
	)

	r.Step(
		demo.S(
			"Connecting to an EKS cluster",
		),
		demo.S(
			"cat notes/eks-connect",
		),
	)

	r.Step(
		demo.S(
			"Using IRSA to interact with AWS resources",
		),
		demo.S(
			"cat notes/irsa",
		),
	)

	r.Step(
		nil,
		demo.S(
			"cat notes/eks-irsa-sa-dep",
		),
	)

	r.Step(nil, nil)
	return r
}
