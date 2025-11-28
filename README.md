# ContainerConnect
A CLI tool written in Go that allows you to interact with running containers in **AWS ECS and EKS**. Connect to containers, list clusters and services/pods, and debug with ease.

## Features

- Connect to running ECS containers via AWS ExecuteCommand
- List ECS clusters and services
- (Future) Connect to EKS pods
- Modular and extensible CLI structure built with [Cobra](https://github.com/spf13/cobra)
- Written in Go for performance and simplicity

## Prerequisites

1. **Go 1.20+** installed on your system
2. **AWS CLI configured** with credentials:
    ```bash
    aws configure
    ```

**Note:** Ensure your IAM user/role has ECS and EKS permissions. For ECS ExecuteCommand, enable SSM and ECS Exec on your services.

## Installation

Clone the repository:
```bash
git clone https://github.com/Parth252/ContainerConnect.git
cd ContainerConnect
```

Build the CLI:
```bash
go build -o containerconnect main.go
```

Run the CLI:
```bash
./containerconnect --help
```

## Usage

### ECS Commands

List ECS clusters:
```bash
./containerconnect ecs list-clusters
```

List services in a cluster:
```bash
./containerconnect ecs list-services --cluster my-cluster
```

Connect to a running container:
```bash
./containerconnect ecs connect --cluster my-cluster --service my-service
```

### EKS Commands (Planned)

List pods in a namespace:
```bash
./containerconnect eks list-pods --namespace my-namespace
```

Connect to a pod:
```bash
./containerconnect eks connect --namespace my-namespace --pod my-pod
```

## Project Structure

```
cmd/
  root.go        # Root command
  ecs/           # ECS commands
     ecs.go
     list_clusters.go
     list_services.go
     connect.go
  eks/           # EKS commands (planned)
     eks.go
     list_pods.go
     connect.go
```

- Each command is implemented as a Cobra command
- Subcommands are modular and easy to extend

## Examples

List clusters and connect:
```bash
./containerconnect ecs list-clusters
./containerconnect ecs list-services --cluster my-cluster
./containerconnect ecs connect --cluster my-cluster --service my-service
```

Expected CLI tree:
```
cc
└── ecs
     ├── list-clusters
     ├── list-services
     └── connect
```

## Contributing

1. Fork the repository
2. Create a new branch for your feature:
    ```bash
    git checkout -b feature/my-feature
    ```
3. Make your changes
4. Submit a pull request

## License

This project is licensed under the Apache License 2.0 — see LICENSE for details.