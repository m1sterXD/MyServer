# 🚧 UNDER DEVELOPMENT 🚧
This project is currently a work in progress. Features and functionality may change as development continues.
## AS IS
### Microservices:
There are two microservices:
- One starts up a server and listens for requests.
- The other makes requests to the server.
Both are containerized using Docker.

### Two more microservices: (added in version 0.2)
- one for grafana
- one for prometheus

### Metrics System:
 Includes Prometheus for metrics collection and Grafana for visualization.


## TO BE
Add a database to store user data. 

# functionality
## Getting Started
1. Clone the Repository
`git clone https://github.com/m1sterXD/MyServer`
2. Build and Run the Containers
Use Docker Compose to build and run the containers:
`docker-compose up --build`
3. Access the Application
- Server: Accessible at http://localhost:8080

- Prometheus: Accessible at http://localhost:9090

- Grafana: Accessible at http://localhost:3000
4. Configure Grafana
- Log in to Grafana ( admin / MyGrafana).

- Add Prometheus as a data source:

- `URL: http://prometheus:9090`

- Create dashboards to visualize your custom metrics.
5. _Enjoy_ (❁´◡`❁)
