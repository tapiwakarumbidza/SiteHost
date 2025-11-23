# GitHub Submission Guide - Step by Step

## ⚠️ IMPORTANT: Manual Steps Required

Since authentication requires your GitHub credentials, here are the exact commands to run:

---

## Step 1: Create Repository on GitHub

1. Go to **https://github.com/new**
2. Fill in:
   - **Repository name:** `SiteHost` (or `hiring-test`)
   - **Description:** "SiteHost Recruitment Challenge Solution"
   - **Visibility:** Public
3. Click "Create repository"
4. You'll see a screen with commands - **use the commands below instead**

---

## Step 2: Add Remote and Configure

Copy and paste these commands into PowerShell:

```powershell
cd "c:\Users\RC_Student_Lab\Downloads\SiteHost"

# Replace YOUR_USERNAME with your actual GitHub username
git remote add origin https://github.com/YOUR_USERNAME/SiteHost.git

# Verify the remote was added
git remote -v
```

**Expected output:**
```
origin  https://github.com/YOUR_USERNAME/SiteHost.git (fetch)
origin  https://github.com/YOUR_USERNAME/SiteHost.git (push)
```

---

## Step 3: Verify Branches

```powershell
git branch -a
```

**Expected output:**
```
* feature/customer-support
  master
```

---

## Step 4: Create Main Branch

First, create and push to the main branch:

```powershell
# Create main branch from master
git checkout -b main master

# Push main branch to GitHub (your first push - may need authentication)
git push -u origin main
```

**First time:** GitHub will prompt for authentication
- Use your GitHub username
- Use a Personal Access Token (PAT) as password (not your GitHub password)

---

## Step 5: Push Feature Branch

```powershell
# Switch to feature branch
git checkout feature/customer-support

# Push the feature branch
git push -u origin feature/customer-support
```

---

## Step 6: Create Pull Request on GitHub

1. Go to your repository: `https://github.com/YOUR_USERNAME/SiteHost`
2. Click "Pull requests" tab
3. Click "New pull request" button
4. Set:
   - **Base branch:** `main`
   - **Compare branch:** `feature/customer-support`
5. Fill in the PR details:

**Title:**
```
SiteHost Recruitment Challenge Solution
```

**Description:**
```
## Solution Summary

This PR contains a complete solution to the SiteHost recruitment challenge.

### What's Included

**Part One - API Integration:**
- Node.js application that calls SiteHost API endpoints
- Retrieves domains and DNS records for Business Systems International (client_id: 100)
- Run: `node challenge.js`

**Part Two - Customer Support:**
- Analysis of customer's website accessibility issue
- Root cause identified: DNS A record misconfiguration
- Professional customer response with resolution steps

### Problem Identified

The domain `site.recruitment.shq.nz` DNS A record points to `192.168.1.10` (private IP) instead of the actual server at `120.138.30.179`.

### Solution Provided

Updated customer response explaining:
1. What the problem is
2. Why it's happening
3. How to fix it
4. Expected resolution timeline

### How to Test

- Run `node challenge.js` to see API output
- Review `reply.md` for customer communication

### Files

- `challenge.js` - Part One API client
- `reply.md` - Part Two customer response
- `access-site.js` - Website access verification
- `README.md` - Project documentation
- `SUBMISSION_NOTES.md` - Technical details
```

6. Click "Create pull request"

---

## Step 7: Request Review

1. On the PR page, look for "Reviewers" section on the right
2. Click "Reviewers"
3. Type: `sitehost-hiring-bot`
4. Select it from the dropdown (or it will be noted in PR description)

---

## ✅ All Done!

Your submission will then be:
- ✅ Visible on GitHub
- ✅ Ready for review by SiteHost team
- ✅ Waiting for feedback (typically 2-3 business days)

---

## 📌 Quick Reference

| Item | Value |
|------|-------|
| Repository | https://github.com/YOUR_USERNAME/SiteHost |
| Base Branch | main |
| Feature Branch | feature/customer-support |
| Reviewer | @sitehost-hiring-bot |
| Files | 11 total (see dir listing) |
| Commits | 4 clean commits |

---

## 🆘 Troubleshooting

### "Authentication failed"
- Use a Personal Access Token (PAT) instead of password
- Get PAT at: https://github.com/settings/tokens

### "Repository already exists"
- Either use a different repo name
- Or delete the repo on GitHub first

### "Branch already exists on remote"
- The branch was already pushed
- Just create the PR directly

---

## 📝 Next Steps After Submission

1. **Share PR Link** - Keep the PR URL handy
2. **Monitor Notifications** - GitHub will notify you of comments/reviews
3. **Respond to Feedback** - If requested, push additional commits
4. **Do NOT merge** - Let SiteHost team do the final review

---

**Your solution is ready! Follow these steps to submit.** 🚀
