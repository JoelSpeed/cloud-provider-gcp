load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "openapiconv",
    srcs = ["convert.go"],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/kube-openapi/pkg/openapiconv",
    importpath = "k8s.io/kube-openapi/pkg/openapiconv",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/klog/v2:klog",
        "//vendor/k8s.io/kube-openapi/pkg/builder3/util",
        "//vendor/k8s.io/kube-openapi/pkg/spec3",
        "//vendor/k8s.io/kube-openapi/pkg/validation/spec",
    ],
)
