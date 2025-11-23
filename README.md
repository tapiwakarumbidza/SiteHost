# SiteHost Recruitment Challenge

This repository contains my solution to the SiteHost recruitment challenge.

## Challenge Overview

The challenge consists of two parts:
1. **Part One**: Call API endpoints to retrieve domain and DNS record information
2. **Part Two**: Troubleshoot a customer's website accessibility issue

## Part One - API Integration

### How to Run

To execute the API client and retrieve domain/DNS information:

```bash
go run main.go
```

This will:
1. Call the `/domains/100` endpoint to get domains for Business Systems International
2. For each domain, retrieve associated DNS zones
3. For each DNS zone, retrieve DNS records
4. Print all information to the screen

### Requirements
- Go 1.13 or later
- Internet connection to access `https://api.recruitment.shq.nz`

## Part Two - Customer Issue Resolution

See `reply.md` for the analysis and response to the customer's website accessibility issue.
