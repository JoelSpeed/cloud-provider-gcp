load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "managedfields",
    srcs = [
        "extract.go",
        "gvkparser.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/apimachinery/pkg/util/managedfields",
    importpath = "k8s.io/apimachinery/pkg/util/managedfields",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/apimachinery/pkg/api/meta",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:meta",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1/unstructured",
        "//vendor/k8s.io/apimachinery/pkg/runtime",
        "//vendor/k8s.io/apimachinery/pkg/runtime/schema",
        "//vendor/k8s.io/kube-openapi/pkg/schemaconv",
        "//vendor/k8s.io/kube-openapi/pkg/util/proto",
        "//vendor/sigs.k8s.io/structured-merge-diff/v4/fieldpath",
        "//vendor/sigs.k8s.io/structured-merge-diff/v4/schema",
        "//vendor/sigs.k8s.io/structured-merge-diff/v4/typed",
    ],
)
