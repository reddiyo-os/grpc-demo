resource "kubernetes_deployment" "orchestrator" {
  metadata {
    name = "orchestrator"
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = "orchestrator"
      }
    }

    template {
      metadata {
        labels ={
          app = "orchestrator"
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
          name    = "orchestrator"
          image   = "mornindew/reddiyo-os-orchestrator-example:latest"
          command = ["./server"]

          port {
            name           = "http-port"
            container_port = 8888
          }

          resources {
            limits {
              cpu    = "100m"
              memory = "500Mi"
            }
          }

          readiness_probe {
            http_get {
              path = "/healthz"
              port = "8888"
            }

            initial_delay_seconds = 5
            period_seconds        = 5
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

resource "kubernetes_service" "orchestrator_service" {
  metadata {
    name = "orchestrator-service"
  }

  spec {
    port {
      name = "http-port"
      port = 8888
    }

    selector ={
      app = "orchestrator"
    }

    type = "LoadBalancer"
  }
}

