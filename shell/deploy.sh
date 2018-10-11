#!/bin/bash
container_name=(main db)
server_address=95.163.202.160
project_name="Frontend"

function deploy() {
	#  rsync的desc会自动创建一个目录，所以这样就是/root/${project_name}
	echo "maybe a little bit slow because will push this file to your-server"
	rsync -avz --delete ../${project_name} root@${server_address}:/root

	cmd="cd ${project_name}/docker;"
	for data in ${container_name[@]}
	do  
	    cmd=${cmd}"docker-compose up --build -d ${data};"
	done

	ssh root@${server_address} ${cmd}
}

function localtest() {
	echo "local debuging....."
	go run main.go
}


function getRemote() {
	echo "getting....."
	rsync -avz --delete root@${server_address}:/root/${project_name} ../
}

DBUSER=root
DBHOST=95.163.202.160
DBNAME=homework
DBPASSWORD=vimi

#get remote database sql to local
function dump() {
	mysqldump -h$DBHOST -u$DBUSER -p$DBPASSWORD $DBNAME > ./db/sql/latest_dump.sql
}

function restore() {
	mysql -h$DBHOST -u$DBUSER -p$DBPASSWORD $DBNAME < ./db/sql/latest_dump.sql
}

case "$1" in
	deploy)
		deploy
		;;

	localtest)
		localtest
		;;

	getRemote)
		getRemote
		;;

	dump)
		dump
		;;

	restore)
		restore
		;;

	*)
		echo "please choose one {dump | restore}"
		exit 1
esac