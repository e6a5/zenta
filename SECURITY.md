# Security Policy

## Reporting Security Vulnerabilities

The zenta team takes security seriously. We appreciate your efforts to responsibly disclose your findings.

### Reporting Process

**Please do NOT report security vulnerabilities through public GitHub issues.**

Instead, please report security vulnerabilities by email to:

Create a GitHub Security Advisory at https://github.com/e6a5/zenta/security/advisories

Please include as much of the following information as possible:

- Type of issue (e.g. buffer overflow, SQL injection, cross-site scripting, etc.)
- Full paths of source file(s) related to the manifestation of the issue
- The location of the affected source code (tag/branch/commit or direct URL)
- Any special configuration required to reproduce the issue
- Step-by-step instructions to reproduce the issue
- Proof-of-concept or exploit code (if possible)
- Impact of the issue, including how an attacker might exploit it

### Response Timeline

- **Initial Response**: Within 48 hours
- **Status Update**: Within 7 days with more detailed response
- **Resolution**: Security patches will be prioritized and released as quickly as possible

### Disclosure Policy

- We will acknowledge receipt of your vulnerability report and send you regular updates about our progress
- We will confirm the problem and determine affected versions
- We will audit code to find similar problems
- We will prepare fixes for all still-supported releases
- We will release new versions with security fixes
- We will publicly announce the vulnerability after patches are available

### Security Considerations for zenta

zenta is designed with security in mind:

**Local Data Only**
- All data is stored locally on your machine
- No data is transmitted over the network
- No external dependencies for core functionality

**File System Access**
- Currently creates no persistent files or directories
- All operations are stateless and in-memory only
- Does not require elevated privileges

**Input Validation**
- All user inputs are validated and sanitized
- No code execution from user input
- No shell command injection vulnerabilities

**Dependencies**
- Minimal external dependencies (only Go standard library)
- Manual dependency reviews during development
- No automated dependency scanning currently implemented

### Supported Versions

We provide security updates for the following versions:

| Version | Supported          |
| ------- | ------------------ |
| 0.3.x   | :white_check_mark: |
| 0.2.x   | :x:                |
| 0.1.x   | :x:                |

### Security Best Practices for Users

When using zenta:

1. **Download from official sources**
   - Use official GitHub releases
   - Verify checksums when possible
   - Avoid third-party redistributions

2. **Privacy**
   - zenta currently stores no persistent data
   - All operations are performed in-memory only

3. **System security**
   - Keep your operating system updated
   - Use standard security practices for your terminal environment

### Attribution

We believe in responsible disclosure and will acknowledge security researchers who help improve zenta's security:

- Security researchers who report vulnerabilities will be credited in release notes
- We maintain a hall of fame for security contributors
- Significant findings may be eligible for a small token of appreciation

Thank you for helping keep zenta and our users safe! 