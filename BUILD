go_get(
    name = 'gosigar',
    get = [
        'github.com/cloudfoundry/gosigar'
    ],
    revision = ['50ddd08d81d770b1e0ce2c969a46e46c73580e2a'],
)

go_binary(
    name = 'zfetch',
    srcs = ['src/zfetch.go', 'src/utils.go'],
    deps = [
        ':gosigar'
    ],
    visibility = ['PUBLIC'],
)
