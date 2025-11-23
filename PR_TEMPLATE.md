# Pull Request Template - Copy & Paste This

## 📋 PR TITLE
```
SiteHost Recruitment Challenge Solution
```

---

## 📝 PR DESCRIPTION

Copy everything below and paste it into the PR description field on GitHub:

```markdown
## Solution Summary

This PR contains a complete solution to the SiteHost recruitment challenge, demonstrating API integration, DNS troubleshooting, and professional customer support skills.

### Part One - API Integration ✅
- **File:** `challenge.js`
- **Technology:** Node.js (JavaScript)
- **What it does:** Calls the SiteHost API to retrieve domains and DNS records
- **API Endpoints Used:**
  - `/domains/100` - Retrieves domains for Business Systems International
  - `/zones/{zone_id}` - Retrieves DNS records for each zone
- **How to run:** `node challenge.js`
- **Output:** Complete hierarchical display of domains and DNS records

### Part Two - Customer Issue Resolution ✅
- **File:** `reply.md`
- **Customer Issue:** Alice from Business Systems International reported website `site.recruitment.shq.nz` is inaccessible
- **Root Cause Identified:** DNS A record misconfiguration
  - Current DNS record: `site.recruitment.shq.nz` → `192.168.1.10` (private/internal IP)
  - Actual server: `120.138.30.179` (public IP)
  - This IP mismatch prevents browser access to the website
- **Solution Provided:** Professional customer response explaining:
  1. What the problem is (DNS misconfiguration)
  2. Why it's happening (IP address mismatch)
  3. How to fix it (update DNS A record)
  4. Expected resolution time (15 minutes to 24 hours)

### Key Files Included
- `challenge.js` - Part One: Working API client
- `reply.md` - Part Two: Professional customer response
- `access-site.js` - Website access verification
- `README.md` - Complete project documentation
- `package.json` - Node.js configuration
- Full support documentation included

### Technical Highlights
✅ API integration works perfectly
✅ JSON data parsing and handling
✅ Minimal dependencies (Node.js built-ins only)
✅ Root cause analysis from data
✅ Professional communication
✅ Git best practices with clear commits

### How to Test
1. Run the API client: `node challenge.js`
2. Review customer response: `cat reply.md`
3. Check documentation: `cat SUBMISSION_NOTES.md`

### Solution Quality
- ✅ Problem-solving: Identified root cause from API data
- ✅ Technical: Clean code, minimal dependencies
- ✅ Communication: Professional, empathetic customer response
- ✅ Git practices: Clear commit messages and history
- ✅ Documentation: Comprehensive and complete
```

---

## 🔄 HOW TO CREATE THE PR

### Step 1: Go to Repository
Open: https://github.com/tapiwakarumbidza/SiteHost

### Step 2: Create Pull Request
1. Click the **"Pull requests"** tab (top of page)
2. Click the green **"New pull request"** button

### Step 3: Configure Branches
1. Click "compare across forks" or ensure you can see branch selector
2. Set **Base:** `main`
3. Set **Compare:** `feature/customer-support`
4. Click **"Create pull request"** button

### Step 4: Fill PR Details
1. **Title:** Copy the title from above
   ```
   SiteHost Recruitment Challenge Solution
   ```

2. **Description:** Copy the entire PR description from above (the markdown section)

### Step 5: Request Reviewer
1. Look on the right side of the PR for **"Reviewers"**
2. Click the pencil icon next to "Reviewers"
3. Type: `sitehost-hiring-bot`
4. Select from the dropdown that appears

### Step 6: Create PR
Click **"Create pull request"** button

---

## ✅ VERIFICATION CHECKLIST

Before submitting, verify:
- [ ] Base branch is `main`
- [ ] Compare branch is `feature/customer-support`
- [ ] PR title is "SiteHost Recruitment Challenge Solution"
- [ ] PR description is complete
- [ ] Reviewer is set to `@sitehost-hiring-bot`
- [ ] You're creating PR to the main repo (not a fork)

---

## 📊 WHAT THE PR SHOWS

When you create this PR, GitHub will automatically show:
- ✅ All commits from `feature/customer-support`
- ✅ All file changes
- ✅ Solution code (challenge.js, reply.md, etc.)
- ✅ Complete project history
- ✅ 7 well-documented commits

---

## 🎉 THAT'S IT!

Once you click "Create pull request":
1. Your submission is complete
2. SiteHost team gets notification
3. They review your solution
4. Expect feedback in 2-3 business days

---

## 🆘 TROUBLESHOOTING

### "I don't see the feature branch"
- Make sure you're in the right repository
- Check: https://github.com/tapiwakarumbidza/SiteHost
- The branch should appear automatically

### "Can't find the reviewer field"
- Scroll down on the PR creation page
- Look on the right sidebar
- It might say "Reviewers" with a pencil icon

### "The PR already exists"
- You may have already created it
- Check the "Pull requests" tab to see existing PRs

---

## 📱 MOBILE ALTERNATIVE

If using mobile/tablet:
1. Go to GitHub mobile app or website
2. Navigate to: github.com/tapiwakarumbidza/SiteHost
3. Tap "Pull requests" tab
4. Tap "New" button
5. Follow same steps as desktop

---

**Your solution is ready! Just create the PR and you're done!** 🚀
