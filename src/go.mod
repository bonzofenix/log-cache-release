module code.cloudfoundry.org/log-cache

go 1.21.0

toolchain go1.21.10

require (
	code.cloudfoundry.org/go-batching v0.0.0-20240522165838-74c69d79243c
	code.cloudfoundry.org/go-diodes v0.0.0-20240515174142-71582f284718
	code.cloudfoundry.org/go-envstruct v1.7.0
	code.cloudfoundry.org/go-log-cache/v2 v2.0.7
	code.cloudfoundry.org/go-loggregator/v9 v9.2.1
	code.cloudfoundry.org/go-metric-registry v0.0.0-20240523160243-6c152ef80e25
	code.cloudfoundry.org/tlsconfig v0.0.0-20240530171334-2593348de0c6
	github.com/Benjamintf1/unmarshalledmatchers v0.0.0-20190408201839-bb1c1f34eaea
	github.com/benbjohnson/jmphash v0.0.0-20141216154655-2d58f234cd86
	github.com/cespare/xxhash v1.1.0
	github.com/cloudfoundry/gosigar v1.3.59
	github.com/dvsekhvalnov/jose2go v1.7.0
	github.com/emirpasic/gods v1.18.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.20.0
	github.com/influxdata/go-syslog/v3 v3.0.1-0.20210608084020-ac565dc76ba6
	github.com/onsi/ginkgo/v2 v2.19.0
	github.com/onsi/gomega v1.33.1
	github.com/prometheus/client_golang v1.19.1
	github.com/prometheus/client_model v0.6.1
	github.com/prometheus/common v0.53.0
	github.com/prometheus/prometheus v1.99.0
	github.com/shirou/gopsutil/v3 v3.24.4
	golang.org/x/net v0.25.0
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.34.1
)

require github.com/go-chi/chi/v5 v5.0.12

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/OneOfOne/xxhash v1.2.8 // indirect
	github.com/alecthomas/units v0.0.0-20231202071711-9a357b53e9c9 // indirect
	github.com/benjamintf1/unmarshalledmatchers v0.0.0-20190408201839-bb1c1f34eaea // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/edsrzf/mmap-go v1.1.0 // indirect
	github.com/go-kit/kit v0.13.0 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/go-task/slim-sprig/v3 v3.0.0 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/pprof v0.0.0-20240528025155-186aa0362fba // indirect
	github.com/lufia/plan9stats v0.0.0-20240513124658-fba389f38bae // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/power-devops/perfstat v0.0.0-20240221224432-82ca36839d55 // indirect
	github.com/prometheus/procfs v0.15.0 // indirect
	github.com/shoenig/go-m1cpu v0.1.6 // indirect
	github.com/square/certstrap v1.3.0 // indirect
	github.com/tklauser/go-sysconf v0.3.14 // indirect
	github.com/tklauser/numcpus v0.8.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.step.sm/crypto v0.46.0 // indirect
	golang.org/x/crypto v0.23.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	golang.org/x/tools v0.21.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240528184218-531527333157 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/prometheus/prometheus => github.com/prometheus/prometheus v2.13.1+incompatible // pinned
