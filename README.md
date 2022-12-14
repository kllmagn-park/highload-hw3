AWS, Golang backend, L7, nginx, Grafana + Prometheus

# Порядок запуска
1. Запустить медленный backend на трех серверах через Dockerfile в папке backend командой `docker build -t backend .` и `docker run -p 5000:5000 backend`
2. Запустить nginx через docker-compose.yaml командой `docker-compose up` в папке balancer
3. Запустить Grafana и Prometheus через docker-compose.yaml командой `docker-compose up` в папке metrics