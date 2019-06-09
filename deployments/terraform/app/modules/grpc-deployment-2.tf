resource "kubernetes_deployment" "grpc_microservice_2" {
  metadata {
    name = "grpc-microservice-2"
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = "grpc-microservice-2"
      }
    }

    template {
      metadata {
        labels = {
          app = "grpc-microservice-2"
        }
      }

      spec {
        container {
          name    = "grpc-microservice-2"
          image   = "mornindew/grpc-demo-microservice:latest"
          command = ["./server"]

          port {
            name           = "http-api"
            container_port = 8888
          }

          resources {
            limits {
              memory = "500Mi"
              cpu    = "100m"
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

resource "kubernetes_service" "grpc_microservice_2" {
  metadata {
    name = "grpc-microservice-2"
  }

  spec {
    port {
      name = "grpc-port"
      port = 50051
    }

    selector = {
      app = "grpc-microservice-2"
    }

    type = "NodePort"
  }
}

