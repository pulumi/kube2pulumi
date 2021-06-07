resource "foo" "kubernetes:core/v1:Namespace" {
apiVersion = "v1"
kind = "Namespace"
metadata = {
name = "foo"
}
}
