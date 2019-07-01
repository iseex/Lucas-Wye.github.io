# Github教育优惠使用之VPS

## Github教育优惠申请 
首先，注册一个[github账号](github.com)

然后在[github education](https://education.github.com/students)上进行学生认证，认证后点击Get your Pack即可
![](1.png)

申请成功后，可以看到一堆的优惠包
![](2.png)

## VPS
​        虚拟专用服务器（Virtual Private Server），是将一台服务器分区成多个虚拟专享服务器的服务。实现VPS的技术分为容器技术和虚拟化技术 。在容器或虚拟机中，每个VPS都可分配独立公网IP地址、独立操作系统、实现不同VPS间磁盘空间、内存、CPU资源、进程和系统配置的隔离，为用户和应用程序模拟出“独占”使用计算资源的体验。VPS可以像独立服务器一样，重装操作系统，安装程序，单独重启服务器。VPS为用户提供了管理配置的自由，可用于企业虚拟化，也可以用于IDC资源租用。 IDC资源租用，由VPS提供商提供。不同VPS提供商所使用的硬件VPS软件的差异，及销售策略的不同，VPS的使用体验也有较大差异。尤其是VPS提供商超卖，导致实体服务器超负荷时，VPS性能将受到极大影响。相对来说，容器技术比虚拟机技术硬件使用效率更高，更易于超卖，所以一般来说容器VPS的价格都高于虚拟机VPS的价格。 这些VPS主机以最大化的效率共享硬件、软件许可证以及管理资源。每个VPS主机都可分配独立公网IP地址、独立操作系统、独立超大空间、独立内存、独立CPU资源、独立执行程序和独立系统配置等. VPS主机用户除了可以分配多个虚拟主机及无限企业邮箱外，更具有独立主机功能, 可自行安装程序，单独重启主机。  
​        VPS本质上是一个服务器，可通过在上面搭载科学上网的代理软件(shadowsocks等)来实现科学上网。VPS首先必须要能够连接互联网，国外的服务器是可以直接上外网的。  

## DigitalOcean
​      在github学生优惠里，可以看到一个DigitalOcean的$50代金券，足够免费使用DO的服务器10个月的时间。 
DigitalOcean的使用方法如下：

### (1)注册DigitalOcean帐号
访问[DigitalOcean](https://m.do.co/c/f102147ffdd1)注册账号(使用此链接注册能得到$10的代金券)

### (2)激活
注册后，需要进行激活。激活有两种方式，一种是绑定信用卡，另一种是用paypal，都要交$5才能激活。推荐使用paypal，可以先注册一个[paypal](http://paypal.com)账号，绑定银行卡，充值$5后再解绑。  

### (3)兑换优惠码
在 Settings 里面的 Billing 会有一个兑换优惠码的地方。输入github教育优惠页面的DigitalOcean那里的优惠码即可。
若是没有找到兑换优惠码的地方，先联系[客服](https://do-support.force.com/s/createticket)
如图:  
![](3.png)

### (4)购买服务器
在DigitalOcean中服务器被称为Droplets，首先选择Create Droplets, 按如下配置即可
![](4.png)  
![](5.png)   
操作系统建议选Centos，版本就选择最新的就可以了，Centos是服务器系统的首选了，Ubuntu比较适合个人使用，在稳定性上Ubuntu不如Centos。关于节点的选择，San Francisco到大陆的网速很快，选择该节点最优。  
SSHKey不是强制的，推荐配置一下，不然每次远程登录时都需要输入密码（这个密码会在服务器搭建完成后收到的邮件中，如果配置的SSH key就不会收到密码），点击New SSH key，得到如下：  
![](6.png). 
将自己电脑上的ssh公钥复制到上面输入框中，在下面的输入框中为其起一个名字即可。  

*获取电脑ssh公钥的方法*
```shell
ssh-keygen # 然后按三次Enter
cat ～/.ssh/id_rsa.pub
```
复制文件id_rsa.pub的全部内容即可。  

### (5)配置服务器
#### 连接
在自己电脑上进行配置，首先，打开系统的终端或者cmd，使用ssh远程连接服务器。
```shell
ssh root@[your vps ip address]
```
#### 配置防火墙
```shell
# 安装firewalld
yum install firewalld firewall-config
# 启动firewalld
systemctl start firewalld
# 修改SSH端口
vi /usr/lib/firewalld/services/ssh.xml
```
此时出现
```xml
<?xml version="1.0" encoding="utf-8"?>
　　<service>
　　　　<short>SSH</short>
   　　<description>Secure Shell (SSH) is a protocol for logging into and executing commands on remote machines. It provides secure encrypted communications. If you plan on accessing your machine remotely via SSH over a firewalled interface， enable this option. You need the openssh-server package installed for this option to be useful.</description>
   　　<port protocol="tcp" port="22"/>
　　</service>
```
修改port="22"改为 port="xxx"即可。

修改后重载firewalld
```shell
firewall-cmd --permanent --add-service=ssh
firewall-cmd --reload
```

#### 安装shadowsocks
(1)安装pip
```shell
yum install m2crypto python-setuptools
easy_install pip
pip install shadowsocks
```
(2)启动shadowsocks
```shell
# 直接在前台启动shadowsocks
ssserver -p 443 -k password -m aes-256-cfb
# 在后台启动shadowsocks
sudo ssserver -p 443 -k password -m aes-256-cfb --user nobody -d start
```
(3)推荐在通过配置文件启动ss，以下命令创建了shadowsocks.json文件，并用vi打开进行编辑
```shell
vi /etc/shadowsocks.json
```
将该文件配置成下面这个样子，注意其中需要修改的地方，第一项 server 修改成vps的IP地址，第二项 server_port 可以不修改，但最好修改一下，修改后要记住这个端口，我们后面配置防火墙的时候要用到，第五项 password 修改成你的密码，其他东西不用修改。
```python
{
    "server":"165.xx.54.189",
    "server_port":8798,
    "local_address": "127.0.0.1",
    "local_port":1080,
    "password":"google",
    "timeout":300,
    "method":"aes-256-cfb",
    "fast_open": false
}
```

通过配置文件启动ss
```shell
# 后台启动
ssserver -c /etc/shadowsocks.json -d start
# 停止
ssserver -c /etc/shadowsocks.json -d stop
```
最好在后台启动ss，这样可以继续进行后续配置  
(4)防火墙添加shadowsocks端口
```shell
firewall-cmd --add-port=8798/tcp --permanent
```
这里要注意的是端口号8388要修改成上面配置文件中的端口号
#### shadowsocks开机自启设置
(1)安装supervisor
```shell
easy_install supervisor
```
(2)创建配置文件
```shell
echo_supervisord_conf > /etc/supervisord.conf
```


修改配置文件
```shell
vi /etc/supervisord.conf
```
在文件末尾加上：
```shell
[program:ssserver]
command = ssserver -c /etc/shadowsocks.json
autostart=true
autorestart=true
```
(3)设置supervisor开机自启
```shell
vi /etc/rc.local
```
在文件末尾另起一行，添加：
```python
supervisord
```
(4)为 rc.local 添加执行权限
```shell
chmod 777 /etc/rc.local
```
这样ss开启自动后台运行就配置好了，然后在DigitalOcean控制面板中重新启动服务器。









