
# readme

关于 token

>https://blog.csdn.net/subfate/article/details/86147645

用户名是 penglonghua

仓库是 go_web


如果是 公有仓库，那么就不需要 <token> ，如果是私有仓库，添加上即可.

clone

```shell
git clone https://<token>@github.com/<用户名>/<仓库名>

git clone https://github.com/penglonghua/go_web

```

push

```shell
git push https://<token>@github.com/<用户名>/<仓库名>

git push https://github.com/penglonghua/go_web

```

如果要更新本地已经有的代码请使用 pull

```shell
git pull https://github.com/penglonghua/go_web

git pull https://github.com/penglonghua/go_web

```


## …or create a new repository on the command line
```shell
git init
git add README.md
git commit -m "first commit"
git remote add origin git@github.com:penglonghua/go_web.git
git push -u origin master
```
这个地方替换成具体的就是

```shell
git remote add origin https://@github.com/penglonghua/go_web
```

## …or push an existing repository from the command line
```shell
git remote add origin git@github.com:penglonghua/go_web.git
git push -u origin master
```

```shell
git remote add origin https://github.com/penglonghua/go_web
```

## …or import code from another repository
You can initialize this repository with code from a Subversion, Mercurial, or TFS project.



***

关于大文件提交的问题：

有两个解决方案：
1 是实现真正的大文件提交.
>https://blog.csdn.net/u012390519/article/details/79441706


2 删除大文件，不做大文件，下面是解决方法. 注意替换成实际的大文件路径.
>https://blog.csdn.net/sinat_26114733/article/details/88121752


官网
>https://git-lfs.github.com/

两个方法都测试，成功.
