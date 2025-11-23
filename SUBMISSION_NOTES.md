# SiteHost Recruitment Challenge - Submission

## Challenge Completion Summary

This submission contains a complete solution to the SiteHost recruitment challenge, demonstrating skills in API integration, DNS troubleshooting, and customer support.

---

## Part One - API Integration

### Solution: `challenge.js`

A Node.js application that calls the SiteHost API to retrieve domain and DNS information for Business Systems International (client_id: 100).

**How to Run:**
```bash
node challenge.js
```

**What It Does:**
1. Calls `/domains/100` endpoint to retrieve all domains
2. For each domain, retrieves associated DNS zones
3. For each zone, retrieves all DNS records
4. Outputs complete hierarchical structure to console

**API Credentials Used:**
- Base URL: `https://api.recruitment.shq.nz`
- API Key: `h523hDtETbkJ3nSJL323hjYLXbCyDaRZ`
- Client ID: `100` (Business Systems International)

**Key Findings:**
- 3 domains registered: `shq.nz`, `sitehost.nz`, `sitename.co.nz`
- Multiple DNS zones per domain
- DNS records include A, AAAA, MX, NS, SOA, and TXT records

---

## Part Two - Customer Issue Resolution

### Problem Analysis

**Customer Issue:** Alice from Business Systems International reported that `site.recruitment.shq.nz` was completely inaccessible with a connection error. She provided the server IP address: `120.138.30.179`.

**Root Cause Identified:**
Using the API data retrieved in Part One, we discovered a **DNS misconfiguration**:
- DNS A record for `site.recruitment.shq.nz` points to: `192.168.1.10` (private/internal IP)
- Actual server location: `120.138.30.179` (public IP)

This IP mismatch prevents browsers and users from reaching the site via the domain name.

### Solution: `reply.md`

A professional response to the customer explaining:
1. What the problem is (DNS A record misconfiguration)
2. Why it's happening (pointing to wrong IP address)
3. How to fix it (update DNS record from 192.168.1.10 to 120.138.30.179)
4. Expected resolution time (15 minutes to 24 hours for DNS propagation)

**Verification:** Successfully accessed the website using the correct IP address and retrieved content as proof.

---

## Files Submitted

- **challenge.js** - Part One API client
- **access-site.js** - Website access script using IP address
- **reply.md** - Customer support response
- **email.md** - Original customer email
- **README.md** - Project documentation
- **package.json** - Node.js dependencies
- **.gitignore** - Git configuration

---

## Git Repository

```
Commits:
  bd6c575 - Part One: Implement API client to fetch domains and DNS records

Branches:
  feature/customer-support (current)
  master
```

All commits have clear, concise messages following best practices.

---

## Technical Stack

- **Language:** JavaScript (Node.js)
- **Dependencies:** None (built-in Node.js modules only)
- **Key Modules:** `https`, `http`

---

## Testing

✅ API client successfully retrieves and displays domain/DNS data
✅ Website accessibility verified using IP address
✅ DNS misconfiguration identified and documented
✅ Professional customer response created
✅ Git repository initialized with meaningful commits

---

## Notes

- Simple, clean code with minimal dependencies as requested
- Clear error handling and user-friendly output
- Professional communication in customer response
- All code commits are comprehensible and follow git best practices
