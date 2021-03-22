ps -ef|grep alist | grep -v grep | awk '{print $2}' | xargs kill -9 2>&1 &
chmod +x ./alist
nohup ./alist > log.log 2>&1 &

echo "服务已启动，如果访问3000端口无响应，请在宝塔面板放行3000端口"