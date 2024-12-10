# Car Management System using Golang

## Project folder structure

```
├── Dockerfile
├── README.md
├── docker-compose.yml
├── driver
│   └── postgress.go
├── go.mod
├── go.sum
├── handler
│   ├── car
│   │   └── car.go
│   └── engine
│       └── engine.go
├── main.go
├── models
│   ├── car.go
│   └── engine.go
├── service
│   ├── car
│   │   └── car.go
│   ├── engine
│   │   └── engine.go
│   └── interface.go
└── store
    ├── car
    │   └── car.go
    ├── engine
    │   └── engine.go
    ├── interface.go
    └── schema.sql
```

## Tech Stacks
- Golang for backend
- PostgreSQL for database
- Docker Compose for containerization
- Grafana for monitoring
- JWT for authentication
- Telemetry for data collection