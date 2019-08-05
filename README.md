# 依赖组件
1:godep:https://github.com/tools/godep  

# 安装  
1：安装模块  
go get github.com/greatming/realgo  

2：创建app  
${GOPATH}/src/github.com/greatming/realgo/install/install.sh  {$appname}  

3:下载app依赖模块  
cd  ${GOPATH}/src/{$appname} && godep restore  

4:运行app  
go run ${GOPATH}/src/{$appname}/main.go  
