#! /bin/sh

kill -9 $(pgrep PDSgroupon)
cd ~/golang/src/deployGroupon/
git pull https://github.com/Janetyu/PDSgroupon.git
mysql -u groupon -p groupon </root/golang/src/PDSgroupon/groupon.sql
LfC7zNreWHnGyjWn
chmod 777 PDSgroupon
./PDSgroupon -c conf/yun_config.yaml &