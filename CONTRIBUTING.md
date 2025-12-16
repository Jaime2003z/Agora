# Contributing to Agora

First of all, thank you for your interest in contributing to **Agora**.

Agora is an open, permissionless coordination engine designed to improve how humans collaborate at local, regional, and global scale.  
This is **infrastructure software**, not a product, not a startup, and not a political campaign.

Contributions are welcome from developers, researchers, auditors, designers, and systems thinkers.

---

## Project Philosophy

Before contributing, please understand the core principles of Agora:

- **Merit over authority**  
  Influence is earned through contribution and execution, not status.

- **Open by default**  
  Code, decisions, and failures are public and auditable.

- **Minimal trust, maximal verification**  
  Systems must be verifiable, deterministic, and resistant to capture.

- **Protocol-first thinking**  
  Agora is a protocol, not a platform. Avoid centralization by design.

- **Long-term resilience over short-term growth**  
  We prioritize correctness, security, and sustainability.

If these principles conflict with your goals, this may not be the right project for you — and that’s okay.

---

## Ways to Contribute

You can contribute in multiple ways:

### 1. Code
- Core protocol (Go)
- Cryptographic primitives (Rust)
- Networking (libp2p)
- Storage and state management
- Governance and execution logic
- Tooling, CLI, SDKs

### 2. Research & Design
- Governance mechanisms
- Reputation systems
- Economic models
- Attack vectors and threat modeling
- Scalability and performance analysis

### 3. Auditing & Review
- Code review
- Security analysis
- Formal verification (where applicable)
- Documentation review

### 4. Documentation
- Improve clarity and accuracy
- Write technical explanations
- Diagrams and protocol specs

### 5. Testing
- Unit tests
- Integration tests
- Adversarial scenarios
- Testnet participation

---

## Repository Structure

High-level structure:

core/ → Protocol implementation (AGPLv3)
crypto/ → Cryptographic primitives (Rust)
clients/ → User-facing applications
sdk/ → Third-party SDKs
docs/ → Specifications and design documents
testnet/ → Network testing and scenarios

If you are unsure where your contribution fits, open a discussion or issue.

---

## Contribution Workflow

1. Fork the repository
2. Create a branch from `main`

feature/<short-description>
fix/<short-description>
research/<short-description>

3. Make small, focused commits
4. Write clear commit messages
5. Open a Pull Request
6. Participate in the review process

---

## Code Guidelines

### General
- Prefer clarity over cleverness
- Deterministic behavior is mandatory
- Avoid hidden side effects
- Explicit > implicit

### Go
- Follow standard Go conventions
- Keep packages small and composable
- Avoid global state where possible
- Document public interfaces

### Rust
- Favor safety and explicitness
- No unsafe code without justification
- Cryptographic code must be reviewed carefully

### Testing
- Tests are not optional
- Security-sensitive code **must** include tests
- Regressions must be documented

---

## Governance of Contributions

Agora follows a **merit-based governance model**:

- There are no permanent maintainers by status alone
- Influence increases with consistent, high-quality contributions
- Poor-quality or malicious contributions reduce trust

Major protocol changes require:
- Clear motivation
- Backward compatibility analysis
- Security considerations
- Community review

---

## Security

If you discover a security vulnerability:

- **Do not open a public issue**
- Follow the process described in `SECURITY.md`

Responsible disclosure is critical.

---

## Licensing

- Core protocol: **AGPLv3**
- Clients & SDKs: permissive licenses (MIT / Apache)

By contributing, you agree that your work will be licensed under the project’s license.

---

## Code of Conduct

We expect:
- Respectful technical discussion
- Critique ideas, not people
- No harassment or ideological battles

See `CODE_OF_CONDUCT.md` for details.

---

## Final Notes

Agora is a long-term effort.

This project is not optimized for:
- quick wins
- hype cycles
- speculative value

It *is* optimized for:
- correctness
- legitimacy
- coordination at scale

If that excites you — welcome.
