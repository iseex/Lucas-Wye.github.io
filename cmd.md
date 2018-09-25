#matlab相关
*无图形界面启动*
matlab -nosplash -nodesktop
*查看工作区变量*
whos

#软件安装相关
*安装*
sudo dpkg -i + filename
*查看已经安装的包*
sudo dpkg -l
sudo dpkg -r filename
*删除包*
sudo apt-get --purge remove +包
sudo apt autoremove
sudo apt-get autoclean

#文件处理命令
mkdir 目录名         => 创建一个目录
rmdir 空目录名      => 删除一个空目录
rm 文件名 文件名   => 删除一个文件或多个文件
rm -rf 非空目录名 => 删除一个非空目录下的一切
touch 文件名        => 创建一个空文件
mv file1 file2    => 将文件 file1，更改文件名为 file2。
mv file1 dir1    => 将文件 file1，移到目录 dir1下，文件名仍为 file1。
mv dir1 dir2    => 若目录 dir2 存在，则将目录 dir1，及其所有文件和子目录，移到目录 dir2 下，新目录名称为 dir1。若目录 dir2 不存在，则将dir1，及其所有文件和子目录，更改为目录 dir2
nautilus dir => 终端中打开文件夹
sudo chmod 777 文件夹    修改文件夹权限
sudo chmod 755 ./test.sh  #使脚本具有执行权限

#系统相关
*修改环境变量*
sudo gedit ~/.profile
source ~/.profile
*访问C盘*
sudo ntfsfix /dev/sda3
*查看字体*
fc-list :lang=zh-cn
Shutdown -h now #立马关机
Shutdown -h 20:25 #系统会在今天20:25关机
Shutdown -r now #系统立马重启
Shutdown -r +10 #系统十分钟后重启
 



#python
pip3 list
python -m pydoc -p 1234
*tensorflow安装*
pip3 install -i https://pypi.tuna.tsinghua.edu.cn/simple/ --upgrade tensorflow 
#配置gedit
*添加插件*
sudo apt-get install gedit-plugins
*gedit找不到外部工具的解决方案*
sudo chown YOUR_USERNAME /home/YOUR_USERNAME/.config/gedit


#安装校园网配置
sudo dpkg -i xl2tpd.deb
*配置:*
sudo vpn-connect -c
*按照提示操作, 注意用户名后面的@a不要漏掉了*
*#连接:*
sudo vpn-connect
*断开:*
sudo vpn-connect -d
*禁止更新*
echo " xl2tpd hold" | dpkg --set-selections

#git操作
*创建空仓库：*
git init
*添加文件*
git add filename
git commit -m 说明
*-m后面输入的是本次提交的说明，可以输入任意内容，当然最好是有意义的，这样你就能从历史记录里方便地找到改动记录。*
git status
git diff
git log
*版本回退*
上一个版本就是HEAD^，上上一个版本就是HEAD^^，当然往上100个版本写100个^比较容易数不过来，所以写成HEAD~100
git reset --hard HEAD^
*记录*
git reflog
*删除*
rm test.txt
git rm test.txt
git commit -m ''

ssh-keygen -t rsa -C "229559117@qq.com"
git remote add origin git@github.com:Lustrew/notebook.git

git push -u origin master
