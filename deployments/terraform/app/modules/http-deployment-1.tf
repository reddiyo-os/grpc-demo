resource "kubernetes_deployment" "http_microservice_1" {
  metadata {
    name = "http-microservice-1"
  }

  spec {
    replicas = 1

    selector {
      match_labels = {
        app = "http-microservice-1"
      }
    }

    template {
      metadata {
        labels ={
          app = "http-microservice-1"
        }
      }

      spec {
        container {
          name    = "http-microservice-1"
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

resource "kubernetes_service" "http_microservice_1" {
  metadata {
    name = "http-microservice-1"
  }

  spec {
    port {
      name = "http-port"
      port = 8888
    }

    selector = {
      app = "http-microservice-1"
    }

    type = "NodePort"
  }
}
