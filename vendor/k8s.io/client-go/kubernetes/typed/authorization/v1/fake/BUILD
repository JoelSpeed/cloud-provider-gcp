load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "fake",
    srcs = [
        "doc.go",
        "fake_authorization_client.go",
        "fake_localsubjectaccessreview.go",
        "fake_selfsubjectaccessreview.go",
        "fake_selfsubjectrulesreview.go",
        "fake_subjectaccessreview.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/k8s.io/client-go/kubernetes/typed/authorization/v1/fake",
    importpath = "k8s.io/client-go/kubernetes/typed/authorization/v1/fake",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/api/authorization/v1:authorization",
        "//vendor/k8s.io/apimachinery/pkg/apis/meta/v1:meta",
        "//vendor/k8s.io/apimachinery/pkg/runtime/schema",
        "//vendor/k8s.io/client-go/kubernetes/typed/authorization/v1:authorization",
        "//vendor/k8s.io/client-go/rest",
        "//vendor/k8s.io/client-go/testing",
    ],
)
