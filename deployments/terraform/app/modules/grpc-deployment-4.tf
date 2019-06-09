resource "kubernetes_deployment" "grpc_microservice_4" {
  metadata {
    name = "grpc-microservice-4"
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = "grpc-microservice-4"
      }
    }

    template {
      metadata {
        labels = {
          app = "grpc-microservice-4"
        }
      }

      spec {
        container {
          name    = "grpc-microservice-4"
          image   = "mornindew/grpc-demo-microservice:latest"
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

resource "kubernetes_service" "grpc_microservice_4" {
  metadata {
    name = "grpc-microservice-4"
  }

  spec {
    port {
      name = "grpc-port"
      port = 50051
    }

    selector = {
      app = "grpc-microservice-4"
    }

    type = "NodePort"
  }
}
