[TOC]

# Unix

## (1)File System
```shell
/  
├── bin  
│   ├── sh  
│   ├── date  
│   └── csh     
├── etc  
│   ├── passwd  
│   └── group  
├── lib  
│   └── libc.so  
├── usr  
│   ├── bin  
│   ├── man  
│   ├── local  
│   │   ├── bin  
│   │   ├── man  
│   │   └── src  
├── dev  
│   ├── ttya  
│   └── null  
├── tmp  
└── home  
```

## (2)Bash

> * login/logout:
```shell
login
logout
exit
```
> * control key:  

| key             |                                              function |
| --------------- | ----------------------------------------------------: |
| ^H or Backspace |                                     erase a character |
| ^U              |                                           cancel line |
| ^S              |                                         pause display |
| ^Q              |                                       restart display |
| ^C              |                                      cancel operation |
| ^D              |                                    signal end of file |
| ^V              | treat following control character as normal character |


> * verison of unix
```shell
uname 
uname -a
```

> * change password
```shell
passwd
```

> * date and time
```shell
date
```

> * useful info
```shell
who
who am i
whoami
ps
stty # terminal control(example: stty; stty erase ^X; stty kill ^Y;)
env # show environment variables
alias # rename commane ( alias l='ls -CF')
man # manual sections
```

> * File Maintenance Commands
```shell
chmod # 
chmod u+w file
chmod g+rw file
chmod go-x file
r = 4
w = 2
x = 1


umask # set in startup files for the account to masks out permissions. 
# umask numbers added to desired permission number equals 7.


chgrp # change the group of the file
chown # change the owner of a file
```
> * link
```shell
ln #link to another file
unlink #remove the link
```

> * find and search
```shell
find 
grep
```
> * process and kill
```shell
ps
ps -l
ps -ef
top # show updating process info and system usage, in 5 seconds interval
pstree # show the hierarchical relations between processes
uptime # tell how long the system has been running
kill -9 processID
```

> * If Conditions
```shell
-r return true (1) if it exists and is readable, otherwise return false (0)
-w true if it exists and is writable
-x true if it exists and is executable
-f true if it exists and is a regular file (or for csh, exists and is not a
directory)
-d true if it exists and is a directory
-e true if the file exists
-o true if the user owns the file
-z true if the file has zero length (empty)
```

> * Shell Startup
```shell
   Bash login shell executes order:  
   /etc/profile -> ~/.bash/_profile ->  ~/.bash_login -> ~/.profile. 
   Bash non-login shell executes:  
   /etc/bashrc and ~/.bashrc
```
> * History Command
```shell
history
```

## (3)修改环境变量

```shell
sudo vim ~/.bash_profile
source ~/.bash_profile
# /etc/profile 
# /etc/paths 
# ~/.bash_profile 
# ~/.bash_login 
# ~/.profile 
# ~/.bashrc 
```
## (4)查看字体
```shell
fc-list :lang=zh-cn
Shutdown -h now #立马关机
Shutdown -h 20:25 #系统会在今天20:25关机
Shutdown -r now #系统立马重启
Shutdown -r +10 #系统十分钟后重启
```

## (5)mac的terminal操作

```shell
单词为单位移动：option+方向键
清除至当前行尾：ctrl+k
光标移到行首：ctrl+a
       行末：ctrl+e
清楚当前行：ctrl+u

```

## (6)vi

#### 1.Basic

| key                 | function             |
| --------            | -----:               |
|i | Insert 模式，按 ESC 回到 Normal 模式.|
|x | 删当前光标所在的一个字符。|
|:wq | 存盘 + 退出 (:w 存盘, :q 退出)(:w 后可以跟文件名）|
|dd | 删除当前行，并把删除的行存到剪贴板里|
|p | 粘贴剪贴板|
|a | 在光标后插入|
|o | 在当前行后插入一个新行|
|O | 在当前行前插入一个新行|
|^ |到本行第一个不是blank字符的位置（所谓blank字符就是空格，tab，换行，回车等） |
|/pattern | 搜索 pattern 的字符串（如果搜索出多个匹配，可按n键到下一个）|
|P | 粘贴|
|yy | 拷贝当前行当行于 ddP|
|u | undo|
|^r | redo|
|:saveas <path/to/file> | 另存为 <path/to/file>|
|:q! | 退出不保存 :qa! 强行退出所有的正在编辑的文件，就算别的文件有更改。|
|:bn 和 :bp | 你可以同时打开很多文件，使用这两个命令来切换下一个或上一个文件|
|. | (小数点) 可以重复上一次的命令|
|N<command> | 重复某个命令N次|
|:N | 到第N行|
|gg  | 到第一行|
|G | 到最后一行|
|W | 到下一个单词的开头。|
|E | 到下一个单词的结尾。|
|% | 匹配括号移动，包括 (, {, [. |
|0y$ | 拷贝当前行|
|fa | 到下一个为a的字符处|
|t, | 到逗号前的第一个字符。逗号可以变成其它字符|
|F 和 T | 和 f 和 t 一样，只不过是相反方向|
|v |可视化|
|^p,^n | 自动补齐|
|qa |把你的操作记录在寄存器 a,于是 @a 会replay被录制的宏, @@ 是一个快捷键用来replay最新录制的宏。|
|:split | 创建分屏 (:vsplit创建垂直分屏)|
| ^w | 切换分屏。 |
|^w_ | 减小尺寸 |
| ^w+ | 增加尺寸 |

#### 2. Cursor Movement
| key                 | function             |
| --------            | -----:               |
| hjkl  | move  |
| ^F | forward one screen  |
| ^B | back one screen  |
| ^D | down half screen  |
| ^U | up half screen  |
| G | go to last line of file  |
| [n] G | go to last line or line [n]  |
| $ | end of current line  |
| ^ | beginning of text on current line  |
| 0 | beginning of current line  |
| [n] w | forward [n] word(s)  |
| [n] b | back [n] word(s)  |
| e | end of word  |

#### 3. Cursor Movement by Matching
| key                 | function             |
| --------            | -----:               |
| f character      | forward to next this character |
| / pattern string      |  go to matched pattern downward |
| ? pattern string      |  go to matched pattern upward  |
| n and N      |  next matched |
| %      |  match parentheses  |

#### 4. Inserting Text
| key                 | function             |
| --------            | -----:               |
| i | insert text before the cursor |
| a | append text after the cursor |
| I | insert text at beginning of line |
| A | append text at end of line |
| o | open new line after current line |
| O | open new line before current line |

#### 5. Deleting Text
| key                 | function             |
| --------            | -----:               |
| dd | delete current line|
|[n] dd or d [n] d |delete [n] line(s)|
|[n] dw | delete [n] word(s) |
| D | delete from cursor to end of line|
| x | delete current character |
| [n] x | delete [n] characters |
| [n] X | delete previous [n] character (like backspace) |

#### 6. Changing Commands 
| key                 | function             |
| --------            | -----:               |
| s/S  | substitute  |
| cw  | change word  |
| [n]cw  | change next [n] word(s)  |
| c$  | change from cursor to end of line  |
| ~  | change case of character  |
| J  | join current line and next line  |
| u  | undo the last command just done  |
| .  | repeat last change |
| [n]yy or y[n]y | yank [n] line(s) to buffer |
| [n] yw | yank [n] word(s) to buffer |
| p  | put yanked or deleted text after cursor |
| P  | put yanked or deleted text before cursor |

#### 7.Command Line Mode
| key | function |
| --------            | -----:               |
| :number |  跳到那一行 |
| -/+ 行数 | 向前或者向后移动若干行 |
| :![cmd] | 暂时退出命令行执行cmd |
| :set all | display all option settings  |
| :[Addr]s/old expr/new string/[g] | 替换当前行/Addr的old expr为new string,[g]加上则替换所有的，否则只替换第一个 |
#### 8. 正则表达式

| key                 | function             |
| --------            | -----:               |
| ? | match any single character at the indicated position |
| * |  match any string of zero or more characters |
| [abc...] | match any of the enclosed characters |
| [a-e] | match any characters in the range a,b,c,d,e |
| [!def] | match any characters not one of the enclosed characters, sh/bash |
| {abc,bcd,cde} | match any set of characters separated by comma (,) (no spaces), bash/csh |
| ~|  home directory of the current user, bash/csh |
| ~ user | home directory of the specified user, bash/csh |
| . | match any single character except newline |
| [^abc] | match any character NOT in the enclosed set |
| ^exp | regular expression must start at the beginning of the line |
| exp$ | regular expression must end at the end of the line |
| \ | treat the next character literally 转义字符 |
| xy*z | xy开头，z结尾的字符串 |

#### 9. More on vi
| key                 | function             |
| --------            | -----:               |
| :n | move to next file |
| :rew | 回到第一个文件 |

## (7)I/O Redirection and Piping

> * stdin: 0
> * stdout: 1
> * stderr: 2  

| key                 | function             |
| --------            | -----:               |
| >,>!,>/| stdout重定向到file |
| >> | stdout重定向到file(不覆盖) |
| < | stdin从file重定向 |
| \| | 管道 |
| tee | 复制stdout |
| >/dev/null | 扔掉输出 |
| 2>file | stderr to file |
| >file 2>&1 | redirect stdout and stderr to file |
| >>file 2>&1 | append stdout and stderr to file |
| 2>&1/|command |
| command>outfile 2>errfile | 分离stdout, stderr到file |


# git操作

### 创建空仓库
```shell
git init
```
### 添加文件
```shell
git add filename
git commit -m 说明
*-m后面输入的是本次提交的说明，可以输入任意内容，当然最好是有意义的，这样你就能从历史记录里方便地找到改动记录。*
```
```shell
git status
git diff
git log
```
### 版本回退
```shell
上一个版本就是HEAD^，上上一个版本就是HEAD^^，当然往上100个版本写100个^比较容易数不过来，所以写成HEAD~100
git reset --hard HEAD^
```
### 记录
```shell
git reflog
```
### 删除
```shell
rm test.txt
git rm test.txt
git commit -m ''
```
### 创建远程
```shell
ssh-keygen -t rsa -C "xxxxxxxxx@xx.com"
```
### 推送至github
```shell
git remote add origin git@github.com:xxx/xxx.git
```
### 第一次推送
```shell
git push -u origin master
```
### 后续推送
```shell
git push origin master
git push
```
### 克隆
```shell
git clone git@github.com:xxx/xxx.git
```
### pull
```shell
git pull # 用于从另一个存储库或本地分支获取并集成(整合)。git pull命令的作用是：取回远程主机某个分支的更新，再与本地的指定分支合并杂。
# git pull <远程主机名> <远程分支名>:<本地分支名>

git pull origin next:master
# 取回origin主机的next分支，与本地的master分支合并

git pull origin next
# 远程分支(next)与当前分支合并，则冒号后面的部分可以省略

git pull origin
# 当前分支与远程分支存在追踪关系，git pull就可以省略远程分支名。

git pull
# 如果当前分支只有一个追踪分支，连远程主机名都可以省略。
```

# Network
## (1)IP Address(5 Classes)

| key                 | function             |
| --------            | -----:               |
| Bits                | 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 |
| Class A             | 0 prefix        suffix|
| Class B             | 1 0 prefix                             suffix    |
| Class C             | 1 1 0 prefix                                                   suffix|
| Class D             | 1 1 1 0 multicast address|
| Class E             | 1 1 1 1 reserved for future use|

## (2)Commands
```shell
ping 
domainname
hostname

cat /etc/hosts # ip address
sudo systemctl restart NetworkManager # hosts生效

cat /etc/resolv.conf # dns server
ip # TCP/IP interface configuration and routing utility
ifconfig # configure a network interface
route # show / manipulate the IP routing table
netstat # show network status (network connections, routing tables, interface statistics, masquerade connections, and multicast memberships)

sudo ping ip地址 -i 0.01 -s 65500 # 每0.01秒给ip地址对应的机器发送65500字节的数据包
```



## (3)Useful Remote Connection Utilities

```shell
• ftp [options] host
transfer file(s) using file transfer protocol
• telnet [host [port]]
communicate with host using telnet protocol
• Are they secured? Any user/passwd information on network?

• rlogin/rsh, login/shell execution remotely on another machine; ssh, remote login or remote execution using secure shell
• rcp/scp, remotely copy files from this machine to another machine; rsync, smartly copy files over network after checking contents.
• curl transfer a URL via HTTP, FTP, IMAP, etc.
• wget download files over the Internet via HTTP or FTP.
• lynx / links text-mode (mini) web browser
```


# gcc 

```shell
gcc -c filename.c # compile only, produce .o
gcc -g # compile for debugging
gcc -o filename.o # 
gcc -O1,2,3,4,s,fast # for optimization level
gcc -Ipathname 
gcc -Dsymbol # define preprocessor symbol
gcc -Ldirectory # add directory to the library search path
gcc -lxyz # link with library libxyz.a or libxyz.so
```

# gdb

```shell
gdb a.out # debug
list
br 8 # breakpoint in line 8
run # 
print value 
next
where
help
quit
```

[more of gdb](https://blog.csdn.net/gatieme/article/details/51671430)

# make

#### Predefined Macros
用make -p 查看
```shell
AS - assembler (as)
CC - C compiler command (cc)
FC - Fortran compiler command (fc)
CPP - C++ preprocessing command ($(CC) -E)
CXX - C++ compiler command (g++)
CFLAGS - C compiler option flags (e.g. -g)
FFLAGS - Fortran compiler option flags (e.g. -g)
LDFLAGS - Linking option flags (e.g. –L /usr/share/lib)
LDLIBS – Linking libraries (e.g. -lm)
```

#### Special Internal Macros
```shell
$*
# The basename of the current target
$<
# The name of a dependency file, as we see on last page
$@
# The name of the current target.
$?
# The list of dependencies that are newer than the target.
```

# ssh
安装 *SSH(Secure Shell)* 服务以提供远程管理服务 
```shell
sudo apt install ssh
```
启动ssh服务
```shell
/etc/init.d/ssh start
sudo service ssh start
```
检测是否已启动
```shell
ps -e | grep ssh
```

SSH远程登录
```shell
ssh username@192.168.0.1
```
将*文件/文件夹*从远程机*下载*到本地(scp) 
```shell
scp -r username@192.168.0.1:/home/username/remotefile.txt
```
将*文件/文件夹*从本地*推送*到远程机(scp) 
```shell
scp -r localfile.txt username@192.168.0.1:/home/username/
```

# tmux
[参考](https://www.jianshu.com/p/48b5b61e1c38)

```shell
# 终端环境查看会话列表
tmux ls 

# 垂直分屏
ctrl+b % 

# 水平分屏
ctrl+b " 

# 根据按箭方向选择切换到某个pane。
ctrl+b Up|Down|Left|Right 
# 滚屏
# 要进入copy-mode，即PREFIX+[，然后就可以用上下键来滚动屏幕，配置了vi快捷键模式，就可以像操作vi一样来滚动屏幕，非常的方便。退出直接按‘q’键即可。
```


