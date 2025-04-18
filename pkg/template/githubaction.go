package template

func init() {
	Default.Add("action", GITHUBACTION, ".github/workflows/build-image.yaml")
}

const GITHUBACTION = `name: Docker Image CI

on:
  push:
    tags:
      - 'v*' # 触发条件为v开头的tags
    branches:
      - master # 触发分支
env:
  REPO: <repo> #  docker仓库类似 naturelingran/net-echo

jobs:
  docker:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v3
    
      - name: Display Tag or Commit ID
        run: | 
              if [ "${{ .LeftDoubleBrace }} github.event_name {{ .RightDoubleBrace }}" == "push" ]; then
                if [ "${{ .LeftDoubleBrace }} startsWith(github.ref, 'refs/tags/') {{ .RightDoubleBrace }}" == "true" ]; then
                  echo "Tag: ${{ .LeftDoubleBrace }} github.ref }}"
                  echo "TAG_NAME=${{ .LeftDoubleBrace }} github.ref {{ .RightDoubleBrace }}" >> $GITHUB_ENV
                else
                  short_commit_id=$(git rev-parse --short "${{ .LeftDoubleBrace }} github.sha {{ .RightDoubleBrace }}")
                  echo "Short Commit ID: $short_commit_id"
                  echo "TAG_NAME=$short_commit_id" >> $GITHUB_ENV
                fi
              fi

      - name: Set up QEMU
        uses: docker/setup-qemu-action@v3

      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v3
  
      - name: Login to Docker Hub
        uses: docker/login-action@v3
        with:
          username: ${{ .LeftDoubleBrace }} secrets.DOCKERHUB_USERNAME {{ .RightDoubleBrace }}
          password: ${{ .LeftDoubleBrace }} secrets.DOCKERHUB_TOKEN {{ .RightDoubleBrace }}

      - name: Build and push
        uses: docker/build-push-action@v5
        with:
          context: .
          push: true
          tags: ${{ .LeftDoubleBrace }} env.REPO {{ .RightDoubleBrace }}:${{ .LeftDoubleBrace }} env.TAG_NAME {{ .RightDoubleBrace }},${{ .LeftDoubleBrace }} env.REPO {{ .RightDoubleBrace }}:latest
          file: ./build/Dockerfile
          platforms: linux/amd64,linux/arm64

      - name: Update repo description
        uses: peter-evans/dockerhub-description@v3
        with:
          username: ${{ .LeftDoubleBrace }} secrets.DOCKERHUB_USERNAME {{ .RightDoubleBrace }}
          password: ${{ .LeftDoubleBrace }} secrets.DOCKERHUB_TOKEN {{ .RightDoubleBrace }}
          repository: ${{ .LeftDoubleBrace }} env.REPO {{ .RightDoubleBrace }}
`
