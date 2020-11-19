package template

func init() {
	Default.Add("BUILDMAKEFILE", BUILDMAKEFILE, "build/Makefile")
}

// BUILDMAKEFILE build目录下的makefile需要和根目录下的makefile合并
const BUILDMAKEFILE = `
project := {{.project}}

.PHONY: help
help: ## 显示make的目标
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: clean
clean: ## 清理
	$(MAKE) -C rpm clean
	rm -rf tgz

.PHONY: rpm	
rpm:
	$(MAKE) -C rpm 

.PHONY: tgz 
tgz: ## 将项目打包为tar.gz
	mkdir -p tgz
	docker run --rm -w / \
		-v $(realpath $(CURDIR)/../../$(project)):/$(project) \
		-v $(CURDIR)/tgz:/v \
		alpine \
		sh -c "tar -cf $(project).tar.gz $(project) && mv $(project).tar.gz /v"
`
