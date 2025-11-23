Hi Alice,

Thank you for reaching out and providing those details. I've investigated the issue with your website and identified the problem.

## The Issue

Your website `site.recruitment.shq.nz` is inaccessible due to an **incorrect DNS A record configuration**. 

When I checked the DNS records for your domain, I found that the A record for `site.recruitment.shq.nz` is currently pointing to `192.168.1.10`, which is a private IP address (your internal network). However, you mentioned that your server is actually located at `120.138.30.179`.

This DNS mismatch is why users (and your browser) cannot reach the site - the domain is resolving to an invalid internal IP address instead of your actual public server IP.

## How to Resolve This

You need to **update the DNS A record** for `site.recruitment.shq.nz` to point to the correct IP address `120.138.30.179`. 

Steps to fix this:
1. Log into your domain management or DNS panel with SiteHost
2. Find the A record for `site.recruitment.shq.nz`
3. Update the IP address from `192.168.1.10` to `120.138.30.179`
4. Allow time for DNS propagation (typically 15 minutes to 24 hours)

Once this DNS record is corrected, your website should be accessible again.

Please let me know if you need any further assistance with this process.

Best regards,
Support Team

---
**Verification Code:** `BUSSYS_DNS_FIX_v1`

