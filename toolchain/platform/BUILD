load("@bazel-zig-cc//toolchain/platform:defs.bzl", "declare_platforms")

package(
    default_visibility = ["//visibility:public"],
)

constraint_setting(
    name = "using_zig_cc",
    default_constraint_value = "zig_cc_disabled",
)

constraint_value(
    name = "zig_cc_disabled",
    constraint_setting = "using_zig_cc",
)

constraint_value(
    name = "zig_cc_enabled",
    constraint_setting = "using_zig_cc",
)

declare_platforms()
