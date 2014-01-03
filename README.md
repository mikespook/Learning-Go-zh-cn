# `学习 Go 语言` - 一本学习 Go 语言的免费电子书。

本书当前包含：

1. 简介：演示了如何安装 Go，并通过一个源代码逐行展示了 Go 语言。
2. 基础：类型，变量和控制流程。
3. 函数：如何编写和使用函数。
4. 包：函数和数据通过包组织在一起。这里将会看到如何编写包。对如何在包中使用单元测试也进行了介绍。
5. 进阶：学习如何创建自定义数据类型，并在其上定义函数（在 Go 中叫做方法）。
6. 接口：Go 并不支持传统意义上的面向对象。在 Go 中核心概念是接口。
7. 并发：通过关键字 go，可以在独立的调度中运行函数（叫做 goroutine）。在这些 goroutine 之间进行通讯是通过 channel 完成的。
8. 通讯：如何建立/读取/写如文件。以及网络相关内容。

每章包含若干个提供了答案的练习来帮助你提升。

* [Go 项目][1]
* [学习 Go 语言][2]
* [Learning Go][4]

## 构建

### 依赖的包 

在 Ubuntu 上需要用到下面的包来从 LaTeX 源文件中构建本书（12.04 已测试）。

* `inkscape`
* `gnumeric`
* `ttf-droid`
* `ttf-dejavu`
* `ttf-sazanami-gothic`	（日文）
* `ttf-arphic-ukai`     
* `texlive-fonts-recommended`
* `texlive-extra-utils`
* `texlive-xetex`
* `texlive-latex-extra`
* `texlive-latex-recommended`
* `latex-cjk-xcjk`     	（中文）
* `ttf-wqy-microhei`	（中文）
* `git-core`
* `GNU make`



### 使用下面的 shell 脚本自动构建和安装。

你可以复制和粘贴下面的代码到 vt100 完成若干包的安装。

```
# 在 Ubuntu 12.04 通过测试
for i in inkscape \
gnumeric \
ttf-droid \
ttf-dejavu \
ttf-sazanami-gothic \
ttf-arphic-ukai \
texlive-fonts-recommended \
texlive-extra-utils \
texlive-xetex \
texlive-latex-extra \
texlive-latex-recommended \
latex-cjk-xcjk \
ttf-wqy-microhei \
git-core \
make \
do 
sudo apt-get install $i -y
done
```

### 检出 `学习 Go 语言` LaTeX 源文件。

使用 http 协议。

```
me@ubuntu1204:~$git clone https://github.com/mikespook/Learning-Go-zh-cn.git 
Cloning into 'Learning-Go-zh-cn'...
remote: Counting objects: 4515, done.
remote: Compressing objects: 100% (1385/1385), done.
remote: Total 4515 (delta 3106), reused 4512 (delta 3104)
Receiving objects: 100% (4515/4515), 1.53 MiB | 1.17 MiB/s, done.
Resolving deltas: 100% (3106/3106), done.
me@ubuntu1204:~$
```

如果已经有了 github 账户，可以使用 git 协议。

```
git clone git@github.com:mikespook/Learning-Go-zh-cn.git
```


### 生成 `学习 Go 语言` 

```
make
```

### 下载已经编译好的版本

最新编译的 pdf 发布在[七牛存储][3]


[1]: http://golang.org  "Go 项目"
[2]: http://www.mikespook.com/learning-go/ "学习 Go 语言"
[3]: http://mikespook.qiniudn.com/%E5%AD%A6%E4%B9%A0%20Go%20%E8%AF%AD%E8%A8%80(Golang).pdf?download "七牛存储"
[4]: http://www.miek.nl/projects/learninggo/index.html
