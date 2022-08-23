# NFT.com Analytics Infra

![image](https://user-images.githubusercontent.com/5006941/178542374-97b8d319-7e55-439f-826d-a45d43039d37.png)
Our analytics infrastructure is deployed using Pulumi. This infrastructure complements the Indexer infrastructure (https://github.com/NFT-com/indexer/tree/master/infra) by sharing networking and database resources.

## Analytics Infrastructure

- CI/CD Pipeline with GitHub, GitHub Actions, Node/Typescript and Pulumi
- Multi-env: Dev, Prod (manual permission required for deployments to any env)
- Secrets managed in Doppler, flow into GitHub Secrets and used in GitHub actions (secrets —> env variables)

### GitHub Deployment Process 

- Pushed branches starting with `fix/` or `feat/` triggers a deployment to the dev environment (nftcom-analytics-dev)
- Merged branches starting with `fix/` or `feat/` into main triggers a deployment to the staging environment (nftcom-analytics-staging)
- Tagging the main branch starting with `v` triggers deployment to the prod environment (nftcom-analytics-prod)

### Analytics AWS Infrastructure Components

- Elastic Container Service (ECS) Cluster & Task Definitions
- ECS Service & Load Balancer for each of the Analytic Components (Graph, Event, Jobs)
- ECS EC2 Capacity Provider (w/ASG & LaunchConfig)
- Elastic Container Registry (ECR)

### Analytics Deployment Notes

- After deployment is triggered, GitHub Actions executes the `action.yml` script, deploys the shared infra, builds the latest images and pushes them to AWS ECR, and finally deploys the ECS cluster including the task definitions to instantiate the analytic API services on ECS.
