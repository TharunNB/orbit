# Orbit

A local-first runtime environment for AI agents and local models.

Orbit provisions and manages the infrastructure required to run local AI workloads reliably — without manually configuring models, runtimes, storage layers, observability, or process management.

```bash
orbit init analyst-agent
orbit run analyst-agent
```

Orbit handles:
- model provisioning
- runtime lifecycle management
- workspace isolation
- vector storage setup
- GPU compatibility checks
- observability foundations
- local process orchestration

---

## Why Orbit Exists

Running local AI systems today is fragmented.

Developers often need to manually configure:
- models
- inference runtimes
- vector databases
- GPU dependencies
- storage cleanup
- process management
- observability tooling

Orbit aims to provide a unified local runtime environment for AI workloads.

---

## What Orbit Is

Orbit is:
- a local AI runtime manager
- an orchestration layer for local AI workloads
- a lifecycle manager for agents and models
- a developer infrastructure tool

Orbit focuses on:
- reliability
- reproducibility
- observability
- local-first execution
- modular integrations

---

## What Orbit Is Not

Orbit is not:
- another agent framework
- a replacement for Ollama or existing inference engines
- a hosted cloud platform
- a no-code AI builder

Orbit integrates with existing ecosystems instead of replacing them.

---

## Current Scope

Orbit is currently focused on:
- local runtime management
- model provisioning
- workspace lifecycle management
- observability foundations
- runtime orchestration

Distributed execution and cloud orchestration are intentionally out of scope for now.

---

## Quick Start

Initialize a new runtime workspace:

```bash
orbit init analyst-agent
```

Run the runtime:

```bash
orbit run analyst-agent
```

Inspect active runtimes:

```bash
orbit ps
```

View logs:

```bash
orbit logs analyst-agent
```

Stop a runtime:

```bash
orbit stop analyst-agent
```

---

## Project Status

Orbit is in early development.

### Implemented

- [x] CLI structure
- [x] Agent workspace initialization
- [x] Background runtime service
- [x] Process lifecycle commands (`run`, `stop`, `logs`, `ps`)
- [x] Interactive model selection screen

### In Progress

- [ ] Reliable model execution
- [ ] Runtime isolation
- [ ] GPU compatibility detection
- [ ] Persistent runtime memory
- [ ] Structured runtime observability

### Planned

- [ ] Built-in vector storage
- [ ] Runtime snapshots
- [ ] Automatic cleanup policies
- [ ] Plugin system
- [ ] Dashboard UI
- [ ] Single-binary installation
- [ ] Linux service integration

---

## Design Goals

Orbit is designed to be:

- local-first
- reproducible
- observable
- modular
- runtime-focused
- infrastructure-oriented

The project prioritizes operational reliability over AI abstraction layers.

---

## Architecture Direction

Orbit is being designed around:

- lightweight runtime orchestration
- isolated workspaces
- local process supervision
- model lifecycle management
- structured observability
- compatibility automation

The long-term goal is to make local AI infrastructure feel predictable and operationally manageable.

---

## Planned Integrations

Orbit aims to integrate with existing ecosystems including:

- Ollama
- OpenAI-compatible APIs
- Qdrant
- OpenTelemetry
- LangGraph
- CrewAI
- Docker / Podman

---

## Development Philosophy

Orbit avoids reinventing existing ecosystems unnecessarily.

Instead of building:
- custom inference engines
- custom vector databases
- custom observability stacks

Orbit focuses on:
- orchestration
- lifecycle management
- compatibility automation
- runtime reliability
- developer experience

---

## License

MIT