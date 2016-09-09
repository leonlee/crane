# Crane

[![Gitter](https://badges.gitter.im/Dataman-Cloud/crane.svg)](https://gitter.im/Dataman-Cloud/crane?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)

![Crane](doc/img/crane-logo-black.png)


Crane, maintained by dataman-cloud, is a docker container control plane based on docker built-in swarm. Besides swarm features, Crane implemented the functionalities usually required by the enterprise, such as private registries authentation, ACL and application stack templates share. The smart fuzzy search function is giving user a hand to skip to the desired page quickly. The stack to be deployed can carry the registry-auth info, which is stored by the private registries authentation function, to pull private registry images, without docker login needed. User can share his/her private images by clicking the *public* button in image management.

## Features

  * **Swarm features**: Portal every feature of swarm almost. Crane has highlighted the common swarm functions and improved the user experience through the friendly frontend.
  * **Stack templates management**: User can save the running stack as a template, then, others will deploy the template as soon as possible.
  * **Image management**: The private image pushed by the user can be publiced to others.
  * **Fuzzy search**: A in-memory index maintained by the backend serves the function.
  * **Node operation**: Crane is showing the detail infos about node such as kernel version, docker info, docker images or containers in the given node and so on.
  * **Overlay network management**: The overlay network CRUD.
  * **Private registries authentation**: User can save his/her private registry username and password to Crane, then, the to-be-deployed stack can use the registry-auth to pull private registry images.
  * **Webssh**: Command 'docker exec' is the magic.

## Installation



## How to use it

Visit .... to use

Please click [Crane in Chinese](https://dataman.gitbooks.io/crane/content/) for more details.

## Community

[![Gitter](https://badges.gitter.im/Dataman-Cloud/crane.svg)](https://gitter.im/Dataman-Cloud/crane?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)

wechat group: 数人云Crane技术交流群

## Contribution

Both pull-request or issue are welcomed from the community.

## convention

  ### repo branch
    * `master`: development branch in active. PR will be merged into this `master` branch.
    * `release`: version control. Tagged commits or hotfix PR will be pushed here. Maintained by the repo owners.

## License

Crane is available under the [Apache 2 license](./LICENSE).



## Quick Start

你可以用下面方式快速启动一个 Crane 环境

执行命令

  ```bash
  ROLEX_IP=192.168.59.105 ./build-and-start.sh
  ```
其中 ROLEX_IP 是运行该程序的主机IP，另外，该主机将作为 swarm manager

## Jenkins 的 exec

  ```bash
  /bin/sh -xe jenkins-job.sh

  ROLEX_IP=$IP ./build-and-start.sh
  sleep 20
  cd api_test&&./run.sh
  ```
<<<<<<< 1c360fa4e25322b3f09cd29e96c4e46fcb554193

## Open ports between the hosts

- TCP port 2377 for cluster management communications
- TCP and UDP port 7946 for communication among nodes
- TCP and UDP port 4789 for overlay network traffic


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
- 按当前的 developer 数量来说， Crane 的 PR 得到2个 +1该 PR 就可以被 merge 到 master branch。
- 负责相关模块的 merge 操作， PR 被 merge 后，原branch A 就会被删除。


## swarm stack 使用

docker stack支持compose功能，现在这个功能点在实验阶段必须要装实验版docker才能看到，
同时安装docker-compose，安装完毕编写docker-compose.yml文件,
使用docker-compose bundle功能生成root.dab文件，然后使用docker stack deploy -f root.dab stackname命令部署服务。

## 参考资料

* https://blog.docker.com/2016/06/docker-1-12-built-in-orchestration/
* swarm 中文文档： https://yeasy.gitbooks.io/docker_practice/content/swarm/usage.html
* docker 学习大杂烩：https://github.com/veggiemonk/awesome-docker
* 网络： http://collabnix.com/archives/1391
=======
>>>>>>> Updated README.md
