# core

> Core AI infrastructure engine for model serving and pipeline orchestration - part of the [TensorRelay](https://github.com/TensorRelay) open-source ecosystem.

## Overview

Designed for ML engineering teams that need more than a hosted API. This project handles the hard parts of running AI in production - inference routing, scaling, failover, and observability - so you can focus on model quality instead of infrastructure.

This repository contains the `core` component, which handles the core runtime responsibilities for TensorRelay.

## Features

- Distributed inference serving with sub-50ms p99 latency
- Horizontal scaling across GPU and CPU clusters
- Built-in model versioning and A/B traffic splitting
- OpenTelemetry-native observability and tracing
- Pluggable storage backends (S3, GCS, local)

## Requirements

- Go 1.22 or later
- Make (optional, for convenience commands)

## Installation

```bash
git clone https://github.com/TensorRelay/core
cd core
go mod download
```

## Usage

```bash
# Run in development
make run

# Build a production binary
make build

# Run tests
make test
```

## Configuration

The service is configured via environment variables. Copy `.env.example` to `.env` and adjust values for your environment.

## Architecture

Written in Go for performance and low memory overhead. Designed to run as a standalone binary or inside a container. Exposes metrics in Prometheus format on `/metrics`.

## Contributing

Contributions are welcome and appreciated! Here's how to get started:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/my-feature`)
3. Make your changes with tests
4. Commit using conventional commits (`git commit -m 'feat: add my feature'`)
5. Push to your fork and open a Pull Request

Please read our contribution guidelines before submitting. All contributors are expected to follow our code of conduct.

## Contact & Support

Have a question, found a bug, or want to collaborate?

- **Email:** [rk822827@gmail.com](mailto:rk822827@gmail.com)
- **GitHub Issues:** [github.com/TensorRelay/issues](https://github.com/TensorRelay)
- **Discussions:** [github.com/TensorRelay/discussions](https://github.com/TensorRelay)

We aim to respond to all inquiries within 48 hours.

## License

Released under the [MIT License](LICENSE). See `LICENSE` for full details.
