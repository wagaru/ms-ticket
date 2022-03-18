# Movie ticket selling system

It was a microservice composed by four different services:
- Show - providing movie/cinmea/show CRUD service
- Booking(developing) - providing ticker reserve / ticket cacneled / ticket waiting
- Payment(developing) - providing payment service
- User(developing) - providing basic user authentication and permission check

## Microservice framework
- go-kit

## Data storage
- MySQL

## Message broker
- RabbitMQ

## Deployment
- Docker + k8s

## Metrics
- prometheus + grafana