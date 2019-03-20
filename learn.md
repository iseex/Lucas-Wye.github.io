# Unix

## 修改环境变量
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
## 查看字体
```shell
fc-list :lang=zh-cn
Shutdown -h now #立马关机
Shutdown -h 20:25 #系统会在今天20:25关机
Shutdown -r now #系统立马重启
Shutdown -r +10 #系统十分钟后重启
```

## mac的terminal操作

```shell
单词为单位移动：option+方向键
清除至当前行尾：ctrl+k
光标移到行首：ctrl+a
       行末：ctrl+e
清楚当前行：ctrl+u

```



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
```
### 克隆
```shell
git clone git@github.com:xxx/xxx.git
```



# Vim

```txt
i → Insert 模式，按 ESC 回到 Normal 模式.
x → 删当前光标所在的一个字符。
:wq → 存盘 + 退出 (:w 存盘, :q 退出)   （陈皓注：:w 后可以跟文件名）
dd → 删除当前行，并把删除的行存到剪贴板里
p → 粘贴剪贴板
a → 在光标后插入
o → 在当前行后插入一个新行
O → 在当前行前插入一个新行
^ → 到本行第一个不是blank字符的位置（所谓blank字符就是空格，tab，换行，回车等）	
/pattern → 搜索 pattern 的字符串（陈皓注：如果搜索出多个匹配，可按n键到下一个）
 P → 粘贴
 yy → 拷贝当前行当行于 ddP
 u → undo
 <Ctrl-r> → redo
:saveas <path/to/file> → 另存为 <path/to/file>
:q! → 退出不保存 :qa! 强行退出所有的正在编辑的文件，就算别的文件有更改。
:bn 和 :bp → 你可以同时打开很多文件，使用这两个命令来切换下一个或上一个文件
. → (小数点) 可以重复上一次的命令
N<command> → 重复某个命令N次
:N ->到第N行
gg 到第一行
G ->到最后一行
W → 到下一个单词的开头。
E → 到下一个单词的结尾。
% : 匹配括号移动，包括 (, {, [. 
0y$ 拷贝当前行
fa → 到下一个为a的字符处
t, → 到逗号前的第一个字符。逗号可以变成其它字符
F 和 T → 和 f 和 t 一样，只不过是相反方向
v 可视化
<C-p>或是<C-n>，自动补齐
qa 把你的操作记录在寄存器 a,于是 @a 会replay被录制的宏, @@ 是一个快捷键用来replay最新录制的宏。
:split → 创建分屏 (:vsplit创建垂直分屏)
<C-w> : 切换分屏。
<C-w>_ : 减小尺寸 
<C-w>+ : 增加尺寸
```
