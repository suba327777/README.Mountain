# README.Mountain

## dependency, environments
- golang 1.21

## demo
<p align="center">
<img src="https://github.com/suba327777/README.Mountain/assets/84484832/2406be34-b831-4cd8-bdd4-09084eee25f6"/>
</p>

## how to use (fpr GitHub Actions)

1. Create your repository.
2. Create workflow file.
path: `.github/workflow/generate-mountain.yml`

example:
```
on:
  schedule: # run 0:00 (JST)
    - cron: "0 9 * * *"
  workflow_dispatch:

jobs:
  readme_mountain:
    runs-on: ubuntu-latest
    name: generate-mountain
    steps:
      - name: Checkout
        uses: actions/checkout@v3
      - name: Use README.Mountain
        uses: suba327777/README.Mountain@release
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
          USERNAME: ${{ github.repository_owner }}
      - name: commit & push
        run: |
          git config user.name  "github-actions[bot]"
          git config user.email "41898282+github-actions[bot]@users.noreply.github.com"
          git add .
          git commit -m "[μRM] generate svg."
          git push
```
