# hugo
https://my.oschina.net/u/4169309/blog/4319906
https://my.oschina.net/u/4364580/blog/3338018






```shell
# use WSL ubuntu, install hugo
sudo apt install hugo -y

hugo new site .

# theme; http://www.flysnow.org/2018/07/29/from-hexo-to-hugo.html
git clone https://github.com/rujews/maupassant-hugo themes/maupassant
# copy themes/maupassant/exampleSite

# local test
hugo server --buildDrafts --baseURL=http://127.0.0.1  --bind=0.0.0.0 --port=12345 -w 
-t hyde        使用hyde主题，如果使用-t 选择了主题会将当前默认的主题覆盖；
 --buildDrafts参数将生成被标记为草稿的页面，是否发布：hugo 会忽略所有通过 draft: true 标记为草稿的文件。必须改为 draft: false 才会编译进 HTML 文件。
 --baseURL=http://www.datals.com   站点监听域名
 --bind=0.0.0.0   监听全部网段
 --port=80        服务监听端口
 -w               如果修改了网站内的信息，会直接显示在浏览器的页面上，不需要重新运行hugo server，方便我们进行修改。

http://127.0.0.1:12345/


# vi config.toml
contentDir = "../gnote/"

# search


# push gitee
git clone https://gitee.com/gameoverboss/gameoverboss.git
# 创建同名仓库，开启pages服务
http://gameoverboss.gitee.io/


```
