# grpc-request-dedup

This project provides a gRPC middleware for request deduplication, helping achieve idempotency in gRPC-based systems. It's designed as a Proof of Concept (PoC) to demonstrate how to handle duplicate requests.

## Features

- Intercepts gRPC unary calls to detect and handle duplicate requests
- Uses request IDs to uniquely identify each request
- Caches responses for duplicate requests, reducing unnecessary processing
- Can helps achieve idempotency in distributed systems if I ever make it useable

## How It Works

1. The middleware intercepts incoming gRPC requests.
2. It extracts a unique request ID from the request metadata.
3. If the request ID is found in the cache, the cached response is returned immediately.
4. For new requests, the handler is called, and the response is cached before being returned.

## Benefits

- **Idempotency**: Ensures that multiple identical requests produce the same result.
- **Efficiency**: Reduces server load by avoiding redundant processing of duplicate requests.
- **Consistency**: Helps maintain data consistency in distributed systems.

## Use Cases

- Retry scenarios: Safely handle client retries without side effects.
- Network issues: Gracefully handle requests that may be sent multiple times due to network problems.

## Limitations

- This is a Proof of Concept and may need further refinement for production use.
- The current implementation uses an in-memory cache, which may not be suitable for all distributed scenarios.
- Cache expiration and cleanup strategies are not implemented in this PoC.

## Getting Started

[Include instructions on how to integrate this middleware into a gRPC server]

## Contributing

This project is open for contributions. Feel free to submit issues, feature requests, or pull requests to help improve this PoC.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
