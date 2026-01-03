# Humanflow Automation 

## Overview
Humanflow Automation is an **educational proof-of-concept** demonstrating how a **resilient, resumable, and well-architected browser automation system** can be designed in **Go**, with strong emphasis on:

- Clean modular architecture  
- Robust error handling and retry mechanisms  
- State persistence and crash recovery  
- Deterministic scheduling and rate limiting  
- Observability through structured logging  

This project is intentionally focused on **system design, correctness, and engineering maturity**, rather than real-world automation execution.

---

## Safety & Compliance Disclaimer (Important)
**This project does NOT automate LinkedIn or any real platform.**

- All workflows operate on **mock HTML pages and simulated data**
- No real user accounts are accessed
- No platform safeguards are bypassed
- No fingerprinting, stealth, or evasion techniques are implemented

This repository exists **strictly for technical evaluation and educational purposes**.

---

## Architecture Summary
The system follows **clean architecture principles**, where each responsibility is isolated, testable, and independently evolvable.

internal/
├── auth/ → Authentication state machine (mocked)
├── browser/ → Browser abstraction layer
├── search/ → Profile discovery logic (mock data)
├── actions/ → Resumable executable actions
├── scheduler/ → Rate limiting, retries, execution control
├── storage/ → SQLite-backed state persistence
├── logging/ → Structured logging (Zap)
├── errors/ → Centralized error taxonomy
├── config/ → Configuration loading & validation
└── models/ → Domain models
mock/ → Simulated execution environment
data/ → Persistent application state


This separation ensures **clarity, extensibility, and safety**.

---

## Execution Model
The automation flow is deterministic and intentionally conservative:

1. Discover targets from mock data  
2. Schedule actions with strict execution limits  
3. Persist actions before execution  
4. Execute idempotent actions  
5. Retry only recoverable failures using exponential backoff  
6. Resume safely after interruptions or crashes  

No action is executed without first being persisted.

---

## Resumable Actions & Retry Strategy
Actions are designed to be **idempotent and resumable**.

- Failures are classified using a **central error taxonomy**
- Only *recoverable errors* are retried
- Fatal errors immediately abort execution
- Retry logic uses **exponential backoff**
- Context cancellation is respected at every step

This approach prevents duplication and ensures predictable behavior.

---

## State Persistence
The system uses **SQLite** to persist:

- Pending actions
- Execution status
- Historical attempts

This allows the application to:
- Survive crashes
- Resume incomplete work
- Maintain execution integrity across restarts

Persistence is treated as a **first-class concern**, not an afterthought.

---

## Scheduling & Rate Limiting
Execution is controlled by a dedicated scheduler that enforces:

- Maximum actions per run
- Delay between actions
- Controlled execution boundaries

A separate **behavior simulation module** exists solely to test scheduling logic in a mock environment and **does not attempt to bypass safeguards**.

---

## Logging & Observability
Structured logging is implemented using **Zap**, producing machine-readable JSON logs.

Logs capture:
- Action lifecycle events
- Retry attempts
- Error classification
- Scheduler decisions

This enables clear debugging and system observability.

---

## Configuration Management
Configuration is externalized via:
- `config.yaml`
- Environment variables (`.env.example`)
- Sensible defaults with validation

This allows safe tuning of system behavior **without code changes**.

---

## Running the Project
```bash
go run ./cmd




Author
Daksh Pachauri
B.Tech CSE — Final Year
