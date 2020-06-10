package grpc

import (
	"context"
        "fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func forwardMetadataHeaders(ctx context.Context, headers []string) context.Context {
	fmt.Println("forwardMetadata",ctx)

	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Println(md)
	if !ok {
		return ctx
	}

	// Turn all matching headers into a flat sequence of key-value
	// pairs, as required by metadata.AppendToOutgoingContext().
	var pairs []string
	for _, key := range headers {
		values := md.Get(key)
		for _, value := range values {
			pairs = append(pairs, key, value)
		}
	}

	if len(pairs) == 0 {
		return ctx
	}
	return metadata.AppendToOutgoingContext(ctx, pairs...)
}

// NewMetadataForwardingUnaryClientInterceptor creates a gRPC request
// interceptor for unary calls that extracts a set of incoming metadata
// headers from the calling context and copies them into the outgoing
// metadata headers. This may, for example, be used to perform
// credential forwarding.
func NewMetadataForwardingUnaryClientInterceptor(headers []string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req interface{}, resp interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		return invoker(forwardMetadataHeaders(ctx, headers), method, req, resp, cc, opts...)
	}
}

// NewMetadataForwardingStreamClientInterceptor creates a gRPC request
// interceptor for streaming calls that extracts a set of incoming
// metadata headers from the calling context and copies them into the
// outgoing metadata headers. This may, for example, be used to perform
// credential forwarding.
func NewMetadataForwardingStreamClientInterceptor(headers []string) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		return streamer(forwardMetadataHeaders(ctx, headers), desc, cc, method, opts...)
	}
}
