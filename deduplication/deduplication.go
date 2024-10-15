package deduplication

import (
	"context"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Cache to store request IDs and responses
var requestCache = sync.Map{}

const RequestIDKey = "request-id"

// CacheItem represents the cached response
type CacheItem struct {
	Response interface{}
}

// UnaryServerInterceptor is the deduplication middleware for unary gRPC calls
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// Extract request-id from metadata
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.InvalidArgument, "No metadata provided")
		}

		// Check if request-id is present
		requestIDs := md.Get(RequestIDKey)
		if len(requestIDs) == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "No request-id provided")
		}
		requestID := requestIDs[0]

		// Check if the request ID exists in the cache
		if cachedResponse, exists := requestCache.Load(requestID); exists {
			return cachedResponse.(*CacheItem).Response, nil
		}

		// Proceed with handling the request
		resp, err := handler(ctx, req)
		if err != nil {
			return nil, err
		}

		// Cache the response
		requestCache.Store(requestID, &CacheItem{Response: resp})

		return resp, nil
	}
}
