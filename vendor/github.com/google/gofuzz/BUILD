load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "gofuzz",
    srcs = [
        "doc.go",
        "fuzz.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/github.com/google/gofuzz",
    importpath = "github.com/google/gofuzz",
    visibility = ["//visibility:public"],
)
