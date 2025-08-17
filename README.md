# Golang Custom Metrics

Uma aplicação Go simples para demonstrar como expor métricas customizadas para o Prometheus e executar em Kubernetes.

## 📋 Índice

- [Visão Geral](#visão-geral)
- [Funcionalidades](#funcionalidades)
- [Pré-requisitos](#pré-requisitos)
- [Instalação](#instalação)
- [Uso](#uso)
- [Deploy no Kubernetes](#deploy-no-kubernetes)
- [Monitoramento](#monitoramento)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Desenvolvimento](#desenvolvimento)
- [Contribuição](#contribuição)

## 🎯 Visão Geral

Esta aplicação Go é um exemplo prático de como implementar métricas customizadas do Prometheus em uma aplicação Go e integrá-la com um stack de monitoramento Kubernetes usando ServiceMonitor.

A aplicação expõe uma métrica customizada chamada `my_custom_metric` que gera valores aleatórios a cada 5 segundos, simulando dados de uma aplicação real.

## ✨ Funcionalidades

- **Métricas Customizadas**: Exposição de métricas Prometheus personalizadas
- **HTTP Server**: Endpoint `/metrics` para coleta de métricas
- **Containerização**: Aplicação totalmente containerizada com Docker
- **Kubernetes Ready**: Manifests completos para deploy em Kubernetes
- **ServiceMonitor**: Integração automática com Prometheus Operator
- **Multi-stage Build**: Dockerfile otimizado para produção

## 🔧 Pré-requisitos

### Para execução local:

- Go 1.24.5 ou superior
- Docker (opcional)

### Para deploy no Kubernetes:

- Cluster Kubernetes
- Prometheus Operator instalado
- kubectl configurado

## 🚀 Instalação

### Execução Local

1. **Clone o repositório:**

```bash
git clone https://github.com/williamkoller/golang-custom-metrics.git
cd golang-custom-metrics
```

2. **Instale as dependências:**

```bash
go mod download
```

3. **Execute a aplicação:**

```bash
go run main.go
```

4. **Acesse as métricas:**

```bash
curl http://localhost:8080/metrics
```

### Usando Docker

1. **Build da imagem:**

```bash
docker build -t golang-custom-metrics .
```

2. **Execute o container:**

```bash
docker run -p 8080:8080 golang-custom-metrics
```

## 📊 Uso

### Endpoint de Métricas

A aplicação expõe métricas no endpoint padrão do Prometheus:

```
GET http://localhost:8080/metrics
```

### Métrica Customizada

- **Nome**: `my_custom_metric`
- **Tipo**: Gauge
- **Descrição**: Exemplo de métrica customizada para Prometheus
- **Comportamento**: Gera valores aleatórios entre 0 e 100 a cada 5 segundos

Exemplo de saída:

```
# HELP my_custom_metric Exemplo de métrica customizada para Prometheus
# TYPE my_custom_metric gauge
my_custom_metric 42.1337
```

## ☸️ Deploy no Kubernetes

### Aplicar todos os manifests:

```bash
kubectl apply -f k8s/
```

### Ou aplicar individualmente:

1. **Deploy da aplicação:**

```bash
kubectl apply -f k8s/deployment.yaml
```

2. **Criar o service:**

```bash
kubectl apply -f k8s/service.yaml
```

3. **Configurar o ServiceMonitor:**

```bash
kubectl apply -f k8s/servicemonitor.yaml
```

### Verificar o deploy:

```bash
# Verificar pods
kubectl get pods -l app=metrics-addon

# Verificar service
kubectl get svc metrics-addon

# Port-forward para testar
kubectl port-forward svc/metrics-addon 8080:8080
```

## 📈 Monitoramento

### ServiceMonitor

O projeto inclui um `ServiceMonitor` que configura automaticamente o Prometheus para coletar métricas da aplicação:

- **Intervalo de coleta**: 15 segundos
- **Endpoint**: `/metrics`
- **Porta**: `http` (8080)

### Integração com Prometheus

O ServiceMonitor possui o label `release: prometheus` que permite ao Prometheus Operator descobrir automaticamente a aplicação e começar a coletar métricas.

### Consultas PromQL Úteis

```promql
# Valor atual da métrica customizada
my_custom_metric

# Taxa de mudança nos últimos 5 minutos
rate(my_custom_metric[5m])

# Valor médio na última hora
avg_over_time(my_custom_metric[1h])
```

## 📁 Estrutura do Projeto

```
golang-custom-metrics/
├── main.go                 # Aplicação principal
├── go.mod                  # Dependências do Go
├── go.sum                  # Checksums das dependências
├── Dockerfile              # Configuração Docker
├── README.md               # Documentação
└── k8s/                    # Manifests Kubernetes
    ├── deployment.yaml     # Deploy da aplicação
    ├── service.yaml        # Service Kubernetes
    └── servicemonitor.yaml # Configuração Prometheus
```

## 🛠️ Desenvolvimento

### Executar testes:

```bash
go test ./...
```

### Build local:

```bash
go build -o metrics-addon .
```

### Verificar formatação:

```bash
go fmt ./...
```

### Verificar linting:

```bash
golangci-lint run
```

### Build da imagem Docker:

```bash
docker build -t williamkoller/metrics-addon:latest .
```

### Push para registry:

```bash
docker push williamkoller/metrics-addon:latest
```

## 🔄 Ciclo de Desenvolvimento

1. **Desenvolvimento local** com `go run main.go`
2. **Teste das métricas** acessando `http://localhost:8080/metrics`
3. **Build da imagem Docker** para validação
4. **Deploy no Kubernetes** para teste completo
5. **Verificação no Prometheus** se as métricas estão sendo coletadas

## 📋 Roadmap

- [ ] Adicionar mais tipos de métricas (Counter, Histogram)
- [ ] Implementar métricas de negócio mais realistas
- [ ] Adicionar testes unitários
- [ ] Configurar CI/CD pipeline
- [ ] Adicionar Grafana dashboards
- [ ] Implementar health checks

## 🤝 Contribuição

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`)
3. Commit suas mudanças (`git commit -am 'Adiciona nova feature'`)
4. Push para a branch (`git push origin feature/nova-feature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.

## 👥 Autor

**William Koller** - [@williamkoller](https://github.com/williamkoller)

---

⭐ Se este projeto te ajudou, considere dar uma estrela no repositório!
