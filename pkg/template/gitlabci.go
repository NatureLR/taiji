package template

func init() {
	Default.Add("gitlabci", GitlabCI, ".gitlab-ci.yml")
}

// GitlabCI 文件模板
const GitlabCI = `# 定义全局变量
variables:
  DOCKER_REGISTRY_URL: "$REGISTRY"

# 定义阶段
stages:
  - build
  - deploy

# 所有stage脚本之前都会执行,用来定义一些要用的变量
before_script:
  - echo "$CI_JOB_NAME由$GITLAB_USER_NAME($GITLAB_USER_ID)"触发
  - echo ====================================before_script开始执行=========================================================== >/dev/null

  # 打印一些环境变量用于调试
  - envs="CI_JOB_NAME PWD CI_PROJECT_NAME CI_PROJECT_ID CI_PROJECT_URL"
  - for v in $envs ;do echo "$v---------->$(printenv $v)";done

  - IMAGE_REPO="$DOCKER_REGISTRY_URL/$CI_PROJECT_NAMESPACE/$CI_PROJECT_NAME/$CI_COMMIT_REF_NAME"
  - IMAGE_REPO="$(echo $IMAGE_REPO|tr "[:upper:]" "[:lower:]")"
  - echo IMAGE_REPO:$IMAGE_REPO

  - IMAGE_TAG=$IMAGE_REPO:${CI_COMMIT_SHA:0:8}
  - echo IMAGE_TAG:$IMAGE_TAG
  - IMAGE_TAG_LATEST=$IMAGE_REPO:latest
  - echo IMAGE_TAG_LATEST:$IMAGE_TAG_LATEST

  - echo ====================================before_script执行完毕=========================================================== >/dev/null

# 定义每个阶段的按钮
image:
  # 绑定的stage
  stage: build
  # 只在那些分支上生效
  only:
    - tags
    - master
    - dev
    
  # 要用到的镜像dokcer in docker
  #image: docker:git
  # 加入的服务
  #services:
  #  - docker:18.09.7-dind
  # 执行方式为手动执行
  when: manual
  # 允许失败
  allow_failure: true
  # 以下所有脚本都是为了执行docker build
  script:
    - echo "Build on $CI_COMMIT_REF_NAME"
    - echo "HEAD commit SHA $CI_COMMIT_SHA"

    # 编译docker镜像 
    # TODO 改为调用makefile
    #- docker build -f ./Dockerfile -t $IMAGE_TAG -t $IMAGE_TAG_LATEST .
    #- docker push $IMAGE_TAG
    #- docker push $IMAGE_TAG_LATEST
    - echo "The build is sucessful,The image is $IMAGE_TAG"
    - echo "The build is sucessful,The image is $IMAGE_TAG_LATEST"

binary:
  stage: build
  when: manual
  allow_failure: true
  image: golang:1.17
  only:
    - tags
    - master
    - dev
  artifacts:
    paths:
      - ./artifacts
  script:
    - make all

# 发布模板
.deploy:
  stage: deploy 
  when: manual
  allow_failure: true
  script:
    - '<执行发布脚本>'
`
