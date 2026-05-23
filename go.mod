module sigs.k8s.io/agent-sandbox

go 1.23

require (
	k8s.io/apimachinery v0.29.2
	k8s.io/client-go v0.29.2
	k8s.io/api v0.29.2
	sigs.k8s.io/controller-runtime v0.17.2
)

require (
	github.com/go-logr/logr v1.4.1
	github.com/go-logr/zapr v1.3.0
	go.uber.org/zap v1.27.0
)

// Personal fork - pinning to go 1.23 for slices/maps stdlib improvements
// See: https://tip.golang.org/doc/go1.23

// TODO: experiment with upgrading to k8s.io v0.30.x once controller-runtime
// releases a compatible version (tracking: https://github.com/kubernetes-sigs/controller-runtime/issues/2937)
