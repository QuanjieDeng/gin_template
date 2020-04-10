## NFS服务器
- mysql的数据需要落地，所以需要一台NFS服务器，需要事先将NFS服务器搭建好 
- NFS对外公布的路径为 /nfs_data   mysql的挂载路径为  /nfs_data/mysql  注意需要配置模式为no_root_squash


## 部署
- 部署脚本为 start.sh 
- conf为对应的配置文件，用来生成ConfigMap 创建之前需要根据实际情况配置好对应的信息

- mysql 的service类型为NodePort，方便直接通过其他工具软件链接操作

## 镜像启动时报错 
- chown: changing ownership of '/var/lib/mysql/': Operation not permitted  
在使用K8S安装msql 磁盘挂载到nfs  原因是nfs默认的配置是 squash_all  
修改/etc/exports中的模式为  no_root_squash
