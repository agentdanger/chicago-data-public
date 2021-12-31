# chicago-data | deployments

## API Status

## Deployment Steps:

1. PostgreSQL Database
2. Namespaces
3. Data Services
4. Report Builds and API Services

## Connection details:

Run the following command to configure kubectl to use our EKS clusters:

	aws eks update-kubeconfig --region <region-code> --name <cluster-name>

Run the following command to configure kubectl to use our local cluster:

	minikube start

## Microservices Available

- [ ] 311 sniffer
- [ ] Business Licenses sniffer
- [X] Crime sniffer
- [X] PostgreSQL database
- [X] Population table
- [ ] Taxi sniffer
- [ ] Traffic crashes sniffer
- [ ] Traffic Network sniffer 
- [ ] CTA "L" station sniffer
- [ ] CTA "L" stops sniffer

- [ ] Crime incidents by type, date and zip code
- [ ] Crime trends by zip code
