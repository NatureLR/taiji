package template

func init() {
	DefaultPool.add("k8s", K8s, "k8s.yaml")
}

// K8s 文件模板
const K8s = `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{.project}}
spec:
  selector:
    matchLabels:
      app: {{.project}}
  template:
    metadata:
      labels:
        app: {{.project}}
    spec:
      containers:
      - name: {{.project}}
        # 此处需要替换为docker镜像地址
        image: <Image>
        resources:
          limits:
            memory: "128Mi"
            cpu: "500m"
        #ports:
        #- containerPort: <Port>
---
apiVersion: v1
kind: Service
metadata:
  name: {{.project}}
spec:
  selector:
    app: {{.project}}
  ports:
  #- port: <Port>
  #  targetPort: <Target Port>
`
