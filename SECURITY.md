# Security Policy

Security is a core concern of **Agora**.

Agora is designed as critical coordination infrastructure.  
Failures, vulnerabilities, or unintended behavior may have **real-world consequences**.

We take security seriously and expect the same from contributors and researchers.

---

## Scope

This security policy applies to:

- Core protocol (`core/`)
- Cryptographic modules (`crypto/`)
- Networking and P2P logic
- Ledger, validation, and execution logic
- Governance and reputation mechanisms
- Any code that affects consensus, state, or execution

Client applications and SDKs are included when they interact with core protocol logic.

---

## Responsible Disclosure

If you discover a security vulnerability:

**DO NOT** open a public GitHub issue.  
**DO NOT** disclose details publicly or on social media.

Instead, follow the responsible disclosure process below.

---

## How to Report a Vulnerability

Please send a detailed report to:

**hireyours@proton.me**  

Include:

- A clear description of the issue
- Affected components and versions
- Steps to reproduce (if applicable)
- Potential impact
- Any proof-of-concept code (optional)
- Your preferred contact information

Encrypted communication is encouraged when possible.

---

## What Qualifies as a Security Issue

Examples include, but are not limited to:

- Ledger manipulation or history rewriting
- Invalid blocks being accepted
- Identity spoofing or reputation forgery
- Governance manipulation or vote inflation
- Economic exploits (double-spend, fund draining, slashing abuse)
- P2P attacks (eclipse, sybil amplification)
- Denial-of-service vectors affecting consensus or execution
- Cryptographic weaknesses or misuse
- Inconsistent state across honest nodes

If you are unsure whether something is a vulnerability, **report it anyway**.

---

## What Is NOT a Security Issue

- Feature requests
- Performance optimizations
- Documentation issues
- UI/UX concerns
- Theoretical attacks without a plausible execution path

These should be reported via standard GitHub issues.

---

## Disclosure Process

1. Report received and acknowledged
2. Initial triage and severity assessment
3. Fix developed privately
4. Patch released
5. Public disclosure (when appropriate)
6. Credit given to reporter (if desired)

We aim to acknowledge reports within **72 hours**, but timelines may vary.

---

## Severity Levels

We roughly classify vulnerabilities as:

- **Critical** – Immediate risk to integrity, funds, or legitimacy
- **High** – Exploitable with moderate effort
- **Medium** – Limited impact or complex exploitation
- **Low** – Minor issues with minimal impact

Severity affects response time and disclosure strategy.

---

## No Bug Bounty (Yet)

Agora currently does **not** offer a paid bug bounty program.

However:
- Responsible disclosures are publicly credited (unless you request anonymity)
- Contributors who consistently help improve security gain reputation within the project
- A formal bounty program may be introduced once the protocol stabilizes

---

## Security Mindset

Agora follows these security principles:

- Explicit rules over implicit trust
- Deterministic behavior
- Minimal attack surface
- Fail-safe defaults
- Defense against capture, not just bugs

Security is not a one-time effort — it is continuous.

---

## Final Notes

Agora is early-stage software.

Do not deploy it in production environments that affect human safety, critical infrastructure, or large-scale funds until explicitly stated otherwise.

By participating in this project, you acknowledge that **security is a shared responsibility**.

---

Thank you for helping make Agora safer.