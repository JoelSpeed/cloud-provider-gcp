load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "gensupport",
    srcs = [
        "buffer.go",
        "doc.go",
        "json.go",
        "jsonfloat.go",
        "media.go",
        "params.go",
        "resumable.go",
        "retry.go",
        "retryable_linux.go",
        "send.go",
        "version.go",
    ],
    importmap = "k8s.io/cloud-provider-gcp/vendor/google.golang.org/api/internal/gensupport",
    importpath = "google.golang.org/api/internal/gensupport",
    visibility = ["//vendor/google.golang.org/api:__subpackages__"],
    deps = [
        "//vendor/github.com/googleapis/gax-go/v2:gax-go",
        "//vendor/google.golang.org/api/googleapi",
    ],
)
