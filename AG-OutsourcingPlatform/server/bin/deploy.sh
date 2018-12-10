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
export SRC_HOME=/var/lib/jenkins/workspace/OutsourcingPlatform/AG-OutsourcingPlatform
export SERVER_PATH=$SRC_HOME/server
export FRONT_PATA=$SRC_HOME/frontend
export GOPATH=$GOPATH:$SERVER_PATH
export SRC_PATH=$SERVER_PATH/src
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
        ng build --no-progress    
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
go build -o ospf  -ldflags "-X main.BUILDTYPE=DEPLOY"
if [  $? -eq 0 ];then
      echo "go build succeess!"
fi
 
if [ ! -z `lsof -i:6616 -t` ];then
APP_PID=$(lsof -i:6616 -t)
     kill -9 $APP_PID
     echo "kill ospf succeess!"
fi

if [ -f "$BIN_PATH/ospf" ];then
	rm $BIN_PATH/ospf
  echo "delete ospf success!"
else 
	echo "no ospf"
fi

mv ./ospf $BIN_PATH
if [ $? -eq 0 ];then
    echo "move file succeess!"
fi

nohup $BIN_PATH/ospf &
if [ $? -ne 0 ]; then
     exit 1
else
    echo "ospf run succeess !!!"
fi
