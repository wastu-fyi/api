package grpc

import (
	"google.golang.org/grpc"
)

type Kernel struct {
}

func (kernel Kernel) UnaryServerInterceptors() []grpc.UnaryServerInterceptor {
	return []grpc.UnaryServerInterceptor{}
}

func (kernel Kernel) UnaryClientInterceptorGroups() map[string][]grpc.UnaryClientInterceptor {
	return map[string][]grpc.UnaryClientInterceptor{}
}
