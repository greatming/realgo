
# 依赖组件
~~1:godep:https://github.com/tools/godep~~

# 安装  
1：安装模块  
go get github.com/greatming/realgo  

2：创建app  
${GOPATH}/src/github.com/greatming/realgo/install/install.sh  {$app_name}  {$app_path}

例如： /Users/hmreal/mygo/lib/src/github.com/greatming/realgo/install/install.sh  mysite  /Users/baidu/opt/

3:下载app依赖模块  
~~cd  ${GOPATH}/src/{$appname} && godep restore~~   
cd {$app_path}{$app_name} && go mod init {$app_name}

4:运行app  
go run main.go  







