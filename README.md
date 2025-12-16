## AGORA

AGORA is an open, decentralized coordination protocol that allows humans to **identify real-world needs, collectively approve executable projects, fund them transparently, and verify their execution** — without relying on states, corporations, or centralized platforms.

AGORA is **not** a DAO, **not** a political system, and **not** a social network.
It is a **verifiable action engine**.

---

## Why AGORA exists

Most collective systems fail because they:
- confuse discussion with execution
- rely on trust instead of verification
- centralize power while claiming decentralization

AGORA solves a simpler problem:

> *How can humans coordinate real action at scale, where collaborating is more logical than not collaborating?*

---

## Core Ideas

### 1. Execution over promises
Projects are not text proposals. Every project must define:
- measurable goals
- locked budgets
- verification methods
- penalties for failure

If it cannot be verified, it cannot be funded.

### 2. Minimal decision-making
Humans do not manage execution.
A vote answers one question only:

> "Do we activate this executable project with these resources?"

Everything else is automatic.

### 3. Human-scale discovery
Projects declare an **impact scope** (coordinates + radius).
Clients use this to show what matters most to each human:
- projects on your street
- projects in your region
- global projects

Levels exist for **visibility and priority**, not authority.

### 4. Reputation, not power
- Reputation is earned by execution
- Reputation is non-transferable
- Higher-impact projects require higher reputation

Money alone grants no control.

---

## What AGORA is NOT

- ❌ A replacement government
- ❌ A voting app
- ❌ A token pump
- ❌ A social media platform

AGORA does not force participation.
It simply makes cooperation provable.

---

## Architecture Overview

AGORA separates **power from interfaces**.

- **Core Protocol (AGPLv3)**
  - identity
  - decision
  - execution
  - ledger

- **Clients (MIT / Apache)**
  - Web
  - Android
  - iOS
  - CLI

- **Crypto modules (Rust)**
  - signatures
  - verification

Clients are replaceable.
The protocol is not.

---

## Repository Structure

```
core/        # protocol logic (AGPLv3)
crypto/      # critical cryptography (Rust)
clients/     # user interfaces
sdk/         # developer integrations
docs/        # protocol documentation
testnet/     # runnable examples
```

---

## License

The AGORA core protocol is licensed under **AGPLv3**.

This guarantees:
- the system always remains auditable
- improvements must be shared
- no entity can capture the protocol privately

Client applications and SDKs may use permissive licenses.

---

## Status

AGORA is in **early research and development**.
Expect breaking changes.

The protocol is being designed in public.

---

## Contributing

We welcome:
- protocol engineers
- cryptographers
- distributed systems hackers
- UI developers
- documentation writers

You do not need permission to contribute.

Read `CONTRIBUTING.md` before submitting code.

---

## Running a Node (early)

```bash
go run core/cmd/agora-node/main.go
```

---

## Philosophy

AGORA assumes:
- humans are fallible
- incentives matter more than ideals
- verification beats trust

The system is intentionally:
- boring
- deterministic
- hard to corrupt

If it feels powerful, it is probably wrong.

## Project structure

Agora/
├── LICENSE
├── README.md
├── CONTRIBUTING.md
├── CODE_OF_CONDUCT.md
├── SECURITY.md
│
├── docs/
│   ├── vision.md
│   ├── discovery.md
│   ├── decision.md
│   ├── execution.md
│   ├── protocol.md
│   ├── metrics.md
│   ├── mission.md
│   ├── treasury.md
│   ├── ledger.md
│   ├── nodes.md
│   └── reputation.md
│
├── core/                  # AGPLv3 (System Core)
│   ├── cmd/
│   │   └── agora-node/
│   │       └── main.go
│   │
│   ├── node/              # P2P Node
│   │   ├── node.go
│   │   ├── networking.go
│   │   └── discovery.go
│   │
│   ├── identity/          # Human Identity
│   │   ├── keys.go
│   │   ├── proof.go
│   │   └── reputation.go
│   │
│   ├── ledger/            # Immutable Ledger
│   │   ├── block.go
│   │   ├── chain.go
│   │   ├── merkle.go
│   │   └── validation.go
│   │
│   ├── governance/        # Governance
│   │   ├── proposal.go
│   │   ├── vote.go
│   │   ├── quorum.go
│   │   └── execution.go
│   │
│   ├── economy/           # Economy
│   │   ├── assets.go
│   │   ├── escrow.go
│   │   ├── milestones.go
│   │   └── slashing.go
│   │
│   ├── api/               # Node API
│   │   ├── grpc/
│   │   └── rest/
│   │
│   └── storage/           # Local persistence
│       ├── db.go
│       ├── state.go
│       └── snapshot.go
│
├── crypto/                # Rust (Critical modules)
│   ├── Cargo.toml
│   ├── src/
│   │   ├── lib.rs
│   │   ├── signatures.rs
│   │   ├── verification.rs
│   │   └── zk.rs
│   └── ffi/
│       └── bindings.go
│
├── clients/               # MIT / Apache
│   ├── web/
│   │   ├── package.json
│   │   └── src/
│   │
│   ├── android/
│   │   └── app/
│   │
│   ├── ios/
│   │   └── AgoraApp/
│   │
│   └── cli/
│       └── agora-cli.go
│
├── sdk/                   # MIT (Third-party SDKs)
│   ├── go/
│   ├── js/
│   └── rust/
│
├── scripts/
│   ├── build.sh
│   ├── testnet.sh
│   └── lint.sh
│
├── testnet/
│   ├── genesis.json
│   ├── nodes/
│   └── scenarios/
│
└── .github/
    ├── workflows/
    └── ISSUE_TEMPLATE.md
