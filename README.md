# Backend Monorepo Microservices

Folder structure inspired by [Building a Monorepo in Golang](https://earthly.dev/blog/golang-monorepo/) and [Golang project layout](https://github.com/golang-standards/project-layout)

## This Project still in WIP
### Todo:

- Will make 5 microservices:
```
1. api gateway
2. auth service
3. users service
4. product service
5. order and payment (maybe separate into 2 microservice) service 
6. cart service
```
- Create observability like:
```
1. Monitoring exposing metrics using Prometheus
2. distributed tracing using jaeger or maybe zipkin
3. logging structure using zap (maybe)
```

- Implement SAGA pattern (need a lot of reading)