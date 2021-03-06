# BlackMamba


## 代码 branch 规划

- master 分支是 开发分支， review 过的代码会放到 master 分支， 大家共同维护。
- release 分支是 版本控制分支，master 的每次发版都会放到 release 分支， 周伟涛维护。

## Code Review


- 凡是需要跟产品经理，设计师等所有非本项目的 developer 沟通时，我们使用 jira 。 它主要用来跟踪项目进度，用户故事， 较大 feature 划分，管理等。
- 纯代码相关的任务、不需要非-developer参与的，如bug, code refactor 等，我们使用github issue 进行管理。
- developer 承担某一个具体的任务A后，它需要在这里新建一个解决该任务的 git branch A。
- developer 将实现A的 commit 提交到branch A。不要将与项目无关的内容 commit，以免困扰你的合作者。
- 另外，提交的 commit 不要引入已知的 bug，所谓已知的 bug 就是该 commit 可能会导致项目无法运行， 或导致某些以前正常的模块失效，而这些是你开发的未完成的功能导致的。 提交时将这些地方注释掉，或者让项目先不触发它，又或者留到下个commit完成后提交。
- developer 完成该具体任务后，他/她需要发一个 PR(pull request) 以告诉其他 developer 进行 code review。
- reviewer(别的 developer) 每天需要拿出部分时间 review 他人的 PR， 发现问题就及时在 PR 里面 comment。
- developer 需要对 reviewer 的 comment 作出解释，或者提交新的 commit 去改进它。 后者你需要向 PR 里追加 commit。
- reviewer 如果对 developer 的 PR 无异议， 他会在 PR 的 comment 里面回复 +1.
- 按当前的 developer 数量来说， BlackMamba 的 PR 得到2个 +1该 PR 就可以被 merge 到 master branch。
- 负责相关模块的 merge 操作， PR 被 merge 后，原branch A 就会被删除。

## 环境启动

### 通用启动方案

   - 项目目录下执行 ``` npm install ``` (前提全局安装了 npm, brew install node)
   - 项目目录下执行 ``` bower install ```(前提全局安装了 bower, npm install -g bower)
   - 修改项目下的 conf.js
   - 项目目录下执行 ``` gulp serve ```(前提全局安装了 gulp, npm install -g gulp)
   - 打开浏览器 默认指定 5000 端口, 例如:http://localhost:5000
    
### 在 Linux 下面 (如果通用方案不行)

  - 第一种方法: 项目根目录下执行 ```python3 -m http.server 8001 ``` 具体端口自己定
  - 第二种方法:
       - 编辑 conf/nginx.conf , 将其中的 `192.168.59.104` 改成本机的 IP
       - `./start.sh` , 该命令会启动一个 nginx container， 并监听 80 端口

### 在 MacOS 下面 (如果通用方案不行)

  - 第一种方法: 项目根目录下执行 ```python3 -m http.server 8001 ``` 具体端口自己定
  - 第二种方法: 单独为项目配置一个 ngnix 文件并启动:

  ```shell
  #user  nginx;
  worker_processes  1;

  events {
      worker_connections  1024;
  }


  http {
      include       /usr/local/etc/nginx/mime.types;
      default_type  application/octet-stream;

      index    index.html;


      sendfile        on;
      #tcp_nopush     on;

      keepalive_timeout  65;

      #gzip  on;

      server {
          listen       8001;
          server_name  blackmamba;
          root    /Users/my9074/work/blackmamba;


          # serve static files
          location ~ ^/(images|javascript|js|css|flash|media|static)/  {
            expires 30d;
          }

          location / {
            try_files $uri /index.html;
          }

          location /auth {
            try_files $uri /auth-index.html;
          }

      }
  }
  ```



## 参考资料

* swarm 中文文档： https://yeasy.gitbooks.io/docker_practice/content/swarm/usage.html
* docker 学习大杂烩：https://github.com/veggiemonk/awesome-docker
