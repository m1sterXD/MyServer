# 🚧 UNDER DEVELOPMENT 🚧
This project is currently a work in progress. Features and functionality may change as development continues.
## AS IS
Microservices:
    There are two microservices:
        One starts up a server and listens for requests.
        The other makes requests to the server.
    Both are containerized using Docker.
Metrics System:
    Includes Prometheus for metrics collection and Grafana for visualization.
    Currently, Prometheus and Grafana are not containerized.

## TO BE
Add a database to store user data.
Containerize Grafana and Prometheus to integrate them with the existing Docker setup.