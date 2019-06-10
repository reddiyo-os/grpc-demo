resource "kubernetes_deployment" "http_microservice_5" {
  metadata {
    name = "http-microservice-5"
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = "http-microservice-5"
      }
    }

    template {
      metadata {
        labels = {
          app = "http-microservice-5"
        }
      }

      spec {
        container {
          name    = "grpc-health-check"
          image   = "mornindew/grpc-demo-healthcheck-sidecar:latest"
          command = ["./server"]

          port {
            name           = "http-api"
            container_port = 8080
          }

          liveness_probe {
            http_get {
              path = "/health/liveness"
              port = "8080"
            }

            initial_delay_seconds = 5
            period_seconds        = 3
          }

          readiness_probe {
            http_get {
              path = "/health/readiness"
              port = "8080"
            }

            initial_delay_seconds = 5
            period_seconds        = 5
            success_threshold     = 1
          }

          image_pull_policy = "Always"
        }        
        container {
          name    = "http-microservice-5"
          image   = "mornindew/grpc-http-demo-microservice:latest"
          command = ["./server"]

          port {
            name           = "http-api"
            container_port = 8888
          }

          resources {
            limits {
              cpu    = "100m"
              memory = "500Mi"
            }
          }

          image_pull_policy = "Always"
        }
      }
    }

    strategy {
      type = "RollingUpdate"

      rolling_update {
        max_unavailable = "0"
        max_surge       = "4"
      }
    }

    min_ready_seconds = 15
  }
}

resource "kubernetes_service" "http_microservice_5" {
  metadata {
    name = "http-microservice-5"
  }

  spec {
    port {
      name = "http-port"
      port = 8888
    }

    selector ={
      app = "http-microservice-5"
    }

    type = "NodePort"
  }
}

