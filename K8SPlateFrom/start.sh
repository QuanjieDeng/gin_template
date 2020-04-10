#! /bin/bash  


NAMESPACE=gvp

#Create Namespace  
kubectl   create  namespace  ${NAMESPACE}
sleep 3
#Create  ConfigMap
kubectl     create   configmap   gvp-config   --from-file=./conf      -n    ${NAMESPACE}
sleep 3
#Create msql
kubectl      apply  -f     mysql-pv.yaml    -n    ${NAMESPACE}
kubectl      apply  -f     mysql-pvc.yaml    -n    ${NAMESPACE}
kubectl      apply  -f     mysql.yaml    -n    ${NAMESPACE}
kubectl      apply  -f     mysql-service.yaml    -n    ${NAMESPACE}
sleep  10
#create  databases
mysqlpods=`kubectl get pods    -o=name     -n     ${NAMESPACE}    | sed "s/^.\{4\}//" |   grep  mysql`
kubectl    exec   -ti    ${mysqlpods}   -n  ${NAMESPACE}    --  mysql  -uroot    -p123456   <   ./sql/init.sql

#Create gvp
kubectl      apply -f   gvp.yaml      -n    ${NAMESPACE}
#Create gvp-service
kubectl      apply -f   gvp-service.yaml      -n    ${NAMESPACE}


#create gvp-ingress
kubectl      apply -f   gvp-ingress.yaml      -n    ${NAMESPACE}


