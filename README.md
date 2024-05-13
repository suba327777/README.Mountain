<div align="center">
   <h1>README.Mountain</h1>

   <p>
      A tool to generate your github mountain for profile README.
   </p>
   <p align="center">
      <a href="https://github.com/suba327777/README.Mountain/stargazers">
      <img alt="Stargazers" src="https://img.shields.io/github/stars/suba327777/README.Mountain?style=for-the-badge&logo=github&logoColor=D9E0EE&labelColor=302D41"></a>
      <a href="https://github.com/suba327777/README.Mountain/releases/latest">
      <img alt="Releases" src="https://img.shields.io/github/release/suba327777/README.Mountain?style=for-the-badge&logo=semantic-release&logoColor=D9E0EE&labelColor=302D41"/></a>
   </p>
</div>

## dependency, environments
- golang 1.21

## Themes

|   |   |   |   |   
|:---:|:---:|:---:|:---:|
|default|dark|sakura|maple|
|![スクリーンショット 2024-05-13 22 48 51](https://github.com/suba327777/README.Mountain/assets/84484832/0fa16586-bbf9-4ef7-ac40-b9b92c4bb35b)|![スクリーンショット 2024-05-13 22 47 30](https://github.com/suba327777/README.Mountain/assets/84484832/9af87b66-b3de-40b1-beca-114893443bc0)|![スクリーンショット 2024-05-13 22 14 26](https://github.com/suba327777/README.Mountain/assets/84484832/9b9d8ed3-cfa8-47fe-9b2c-a363c151e029)|![スクリーンショット 2024-05-13 22 15 02](https://github.com/suba327777/README.Mountain/assets/84484832/8569962e-ceea-431a-a15c-0440f312f7be)



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
          # select theme
          THEME: "default"
      - name: commit & push
        run: |
          git config user.name  "github-actions[bot]"
          git config user.email "41898282+github-actions[bot]@users.noreply.github.com"
          git add .
          git commit -m "[μRM] generate svg."
          git push
```
