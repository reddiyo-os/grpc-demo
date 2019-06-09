resource "kubernetes_deployment" "http_microservice_2" {
  metadata {
    name = "http-microservice-2"
  }

  spec {
    replicas = 1

    selector {
      match_labels= {
        app = "http-microservice-2"
      }
    }

    template {
      metadata {
        labels = {
          app = "http-microservice-2"
        }
      }

      spec {
        container {
          name    = "http-microservice-2"
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

resource "kubernetes_service" "http_microservice_2" {
  metadata {
    name = "http-microservice-2"
  }

  spec {
    port {
      name = "http-port"
      port = 8888
    }

    selector ={
      app = "http-microservice-2"
    }

    type = "NodePort"
  }
}

