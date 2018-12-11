#!/bin/bash

#指令执行错误的操作
exitOnErr(){
  exitCode=$?
  echo "[LINE:$1] Error: Command  exited with status $?"
  exit $exitCode
}

#捕获错误信号，并执行相应操作
trap 'exitOnErr $LINENO' ERR

export PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:~/bin
export PROJECT_HOME=/var/lib/jenkins/workspace/OspfDev/AG-OutsourcingPlatform
export SERVER_PATH=$PROJECT_HOME/server
export FRONT_PATA=$PROJECT_HOME/frontend
export GOPATH=$GOPATH:$SERVER_PATH
export SRC_PATH=$SERVER_PATH/src
export GOTEST=$SERVER_PATH/src/OutsourcingPlatform/ospf/TreadPit
export BIN_PATH=$SERVER_PATH/bin
 

#构建前端
cd $FRONT_PATA
echo $INSTALL_DEP

if [  "$INSTALL_DEP" = "TRUE"  ];then
	echo "frontend npm install ..."
	npm install
fi

if [  "$BUILD_DEP" = "TRUE"  ];then
   
   if [  "BUILD_TYPE"="DEV"  ];then
      	echo "frontend build DEV"
        ng build --aot --output-hashing=all --sourceMap=false --extractCss=true --namedChunks=false --buildOptimizer=true  
   fi
     
   if [ "BUILD_TYPE" = "PROD" ];then
        echo "frontend build PROD"
        ng build --prod --no-progress   
   fi
   
fi

rm -r ../server/bin/dist
if [ $? -eq 0 ];then
    echo "rm dist succeess!"
fi

mv dist/ ../server/bin/
if [ $? -eq 0 ];then
    echo "move dist succeess!"
fi

#构建后端
cd $SRC_PATH/OutsourcingPlatform/app

echo "go build...."
go build -o ospfdev  -ldflags "-X main.BUILDTYPE=DEPLOY"
if [  $? -eq 0 ];then
      echo "go build succeess!"
fi
 
if [ ! -z `lsof -i:6616 -t` ];then
APP_PID=$(lsof -i:6616 -t)
     kill -9 $APP_PID
     echo "kill ospfdev succeess!"
fi

if [ -f "$BIN_PATH/ospfdev" ];then
	rm $BIN_PATH/ospfdev
  echo "delete ospfdev success!"
else 
	echo "no ospfdev"
fi

mv ./ospfdev $BIN_PATH
if [ $? -eq 0 ];then
    echo "move file succeess!"
fi

nohup $BIN_PATH/ospfdev &
if [ $? -ne 0 ]; then
     exit 1
else
    echo "ospfdev run succeess !!!"
fi

if [ "BUILD_TEST" = "TRUE" ];then
echo "run go test"
cd $GOTEST
go test
  if [ $? -eq 0];then
    echo "test succeess!"
  fi
fi





