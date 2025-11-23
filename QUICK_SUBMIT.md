# GitHub Submission - Commands to Execute

## Prerequisites
1. Have a GitHub account
2. Have a Personal Access Token (get at https://github.com/settings/tokens)

---

## Commands to Run (Copy & Paste)

### Step 1: Verify Current State
```powershell
cd "c:\Users\RC_Student_Lab\Downloads\SiteHost"
git status
git log --oneline -5
```

### Step 2: Create Main Branch (if not exists)
```powershell
git checkout -b main master
```

### Step 3: Add GitHub Remote
```powershell
# REPLACE YOUR_USERNAME with your actual GitHub username
git remote add origin https://github.com/YOUR_USERNAME/SiteHost.git
git remote -v
```

### Step 4: Push Main Branch
```powershell
git push -u origin main
```
(First time will ask for GitHub username and Personal Access Token)

### Step 5: Push Feature Branch
```powershell
git push -u origin feature/customer-support
```

### Step 6: Verify on GitHub
1. Go to: https://github.com/YOUR_USERNAME/SiteHost
2. You should see both branches
3. feature/customer-support should have your 5 commits

### Step 7: Create PR on GitHub Website
1. Go to: https://github.com/YOUR_USERNAME/SiteHost
2. Click "Pull requests" tab
3. Click "New pull request"
4. Base: main → Compare: feature/customer-support
5. Click "Create pull request"
6. Add the PR description (see GITHUB_SUBMISSION_STEPS.md)
7. Request review from: @sitehost-hiring-bot

---

## Full Command Block (Can Copy All at Once)

```powershell
cd "c:\Users\RC_Student_Lab\Downloads\SiteHost"

# Create main branch
git checkout -b main master

# Add remote (CHANGE YOUR_USERNAME TO YOUR USERNAME)
git remote add origin https://github.com/YOUR_USERNAME/SiteHost.git

# Push branches
git push -u origin main
git push -u origin feature/customer-support

# Verify
git branch -a
git remote -v
```

---

## What Happens Next

✅ Your code is on GitHub
✅ Ready for PR review
✅ SiteHost team reviews
✅ Expect feedback in 2-3 business days

---

## Important Notes

- **Replace YOUR_USERNAME** with your actual GitHub username
- **Main branch** is created from master for GitHub standard
- **Feature branch** contains all your solution commits
- **Do NOT merge** - let SiteHost review first
- **Keep the PR open** - they'll comment/review there

---

Ready to go! 🚀
