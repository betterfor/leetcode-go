## 1、docker 核心原理

**限制的视图namespace**

| namespace | 隔离内容                   |
| --------- | -------------------------- |
| IPC       | 信号量、消息队列和共享内存 |
| Network   | 网络设备、网络栈、端口等   |
| Mount     | 挂载点(文件系统)           |
| PID       | 进程编号                   |
| User      | 用户和用户组               |
| UTC       | 主机名和域名               |

**限制资源cgroups：包括cpu、内存、磁盘、网络带宽等**

- cpu，cpu.cfs_quota_us容量，cpu.cfs_period_us设置使用量
- blkio，为块设备设定I/O限制，一般用于磁盘
- cpuset，为进程分配独立的cpu核和对应的内存节点
- memory，内存

将被限制的进程PID写入tasks文件，上面的设置就会对进程生效了

**rootfs**

`chroot`:change root file system。改变进程的根目录到你指定的目录

1、启动linux namespace配置
2、设置指定的cgroups参数
3、切换进程的根目录

rootfs只是一个操作系统所包含的文件、配置和目录，并不包括操作系统

**UFS(Union File System)联合文件系统**

层（layer）：用户制作镜像的每一个操作，都会生成一个层，也就是一个增量rootfs

UnionFS主要功能是将多个不同位置的目录联合挂载到同一个目录下。`mount -t aufs`

容器的rootfs

- 可读层：挂载方式是只读，`readonly+whiteout`.都以增量的方式包含了操作系统的一部分
- 可读写层： rw，在没写入文件之前，这个目录是空的。而一旦在容器里做了写操作，修改的内容会以增量的方式出现在这个层中。
可以使用docker commit和push指令，保存这个被修改过的可读写层，原先的只读层不会有任何变化。（为了实现删除可读层里的文件删除操作，aufs会创建一个whiteout文件，把只读层文件“遮挡”起来）
- Init层：docker项目单独生成的内部层，专门用来存放/etc/hosts,/etc/resolv.conf等信息，这些文件本来属于只读层，
但是用户需要在容器启动时写入一些指定的值比如hostname，这些修改只对当前容器生效，我们不希望在docker commit时把信息连同读写层一起提交掉。

## 2、pivot_root和chroot的区别
pivot_root主要是把整个系统切换到一个新的root目录，而移除对之前root文件系统的依赖，这样你能够umount原先的root文件系统。
chroot是针对进程，而系统的其他部分仍运行在老的root目录

## 3、Dockerfile内容
- FROM：指定基础镜像
- MAINTAINER：维护者信息
- RUN：构建镜像时执行的命令
- ADD：将本地文件添加到容器中，tar类型文件会自动解压(网络压缩资源不能解压)
- COPY：类似于ADD，但不能自动解压文件，也不能访问网络资源
- CMD：容器启动后调用，在`docker run`时运行
- ENTRYPOINT：类似于CMD，但不会被`docker run`的命令行参数的指令覆盖
- LABLE：为镜像添加元数据
- ENV：环境变量
- ARG：构建参数，仅在`docker build`有效，构建好的镜像不存在此环境变量
- VOLUME：持久化目录
= EXPOSE：外界交互端口
- WORKDIR：工作目录
- USER：执行后续命令的用户和用户组
- HEALTHCHECK：健康检查
- ONBUILD：延迟构建命令，有新的Dockerfile使用之前的镜像会执行ONBUILD命令

## 4、优化Dockerfile
- 构建顺序影响缓存利用率。
把不需要经常更改的行放在最前面，更改频繁的行放在最后面
- 只拷贝文件，防止溢出
- 最小化可缓存的执行层，每一个`RUN`指令都被看作可缓存的执行单元，太多的`RUN`指令会增加镜像层数，增大体积。
将更新缓存和安装文件放在同一个`RUN`指令中。
- 减小镜像体积：删除不必要的依赖 `apt --no-install-recommends`；删除包管理工具缓存
- 使用多阶段构建

## 5、docker load 加载一个镜像，docker images查看不到的原因


## 6、docker后端存储

**存储驱动**

1、 AUFS是一种Union FS，是文件级的存储驱动。AUFS能覆盖一或多个现有文件系统的层状文件系统，
把多层合并成文件系统的单层显示。简单来说就是支持将不同目录挂载到同一个虚拟文件系统下的文件系统。
这种文件系统可以一层一层地叠加修改文件。无论底下有多少层都是只读的，只有最上层的文件系统是可写的。

- 优点：性能稳定，测试完善，适用场景丰富
- 缺点：只在ubuntu和Debian，没有进内核

2、 Overlay： 一个upper文件系统和一个lower文件系统，代表Docker的镜像层和容器层

- 优点：合并进内核
- 缺点：硬连接的方式会引发inode耗尽的问题，整体不成熟

Overlay2是当前所有受支持的linux发行版的首选存储程序

3、 Device mapper：从逻辑设备到物理设备的映射框架机制，在设备创建一个资源池，然后在资源池上创建一个带有文件系统的基本设备，所有镜像都是这个基本设备的快照，
而容器是镜像的快照。所以在容器里看到文件系统是资源池上基本的文件系统的快照，并不为容器分配空间。 块级存储

- 优点：基于块设备而不是基于文件，会拥有一些内置的能力如配额支持
- 缺点：没有入门的支持

4、 Btrfs：下一代写时复制文件系统，文件级存储，把文件系统一部分配置为一个完整的子文件系统，称为subvolume。一个大的文件系统可以被划分为多个文件系统，
这些子文件系统共享底层的设备空间

- 优点：比较健壮，收到良好支持
- 缺点：没有成为linux发行版的主流选择

5、ZFS：使用zfs驱动需要有ZFS格式化的块设备挂载到graphdriver路径（默认/var/lib/docker）,以快照的克隆作为分享层的途径。它不是基于文件的实现。

- 优点：拥有较好的性能，有配额的支持
- 缺点：没有基于文件(inode)的共享达到内库共享

**总结**
+ overlay2，aufs，overlay都在文件级别，这样可以更有效地使用内存，但在写繁重的工作负载中，容器的可写层可能会变得非常大
+ 块级存储器：devicemapper，btrfs，zfs更好地为写繁重的工作负载
+ 对于许多小型写入或具有多层或深文件系统的容器，overlay的性能比overlay2更好，但会消耗更多的inode，这可能导致inode耗尽
+ btrfs和zfs需要大量的内存
+ zfs对于高密度工作负载是一个不错的选择

[docker 面试](https://www.jianshu.com/p/2de643caefc1)
[docker storage driver](https://docs.docker.com/storage/storagedriver/select-storage-driver/)

## 7、docker预热

没有平台之前，运维人员安装微服务的时候，需要先拉取镜像传到内网，然后再执行安装

## 8、OCI镜像规范

- 镜像索引（Image Index）:可选部分，主要作用是指向镜像的不同平台版本
- 清单（Manifest）：镜像包含的配置和内容文件。
  主要作用是
  * 支持内容可寻址的镜像模型，在该模型中可以对镜像的配置进行哈希处理，生成镜像及其唯一标识
  * 通过镜像索引包含的多体系结构镜像，通过引用镜像清单获取特定平台的镜像版本
  * 可转为OCI运行时规范以运行容器
- 配置（Configuration）：描述容器的根文件系统和容器运行时使用的执行参数，还有一些镜像的元数据。
  * 在配置规范里定义了镜像的文件系统的组成方式。镜像文件系统由若干镜像层组成，每一层都是tar包，除了底层（base image），其余层都是记录了父层向下一层文件系统的变化集（changeset）
- 层文件（Layers）：镜像的根文件由多个层文件叠加而成，每个层文件在分发时都被打成了tar包

## 9、网络

- host模式：容器和宿主机共享network namespace
- container模式：容器和另外一个容器共享一个network namespace
- none模式：独立的network namespace，但没有对其进行任何的网络配置
- bridge模式：默认，docker启动时会在主机创建一个docker0的虚拟网桥，此主机上的所有容器都会连接在虚拟网桥上。虚拟网桥的工作方式与物理交换机类似，这样主机上所有的容器通过交换机连在了一个二层网络中。
从docker0分配一个IP给容器使用，并配置docker0的IP地址为容器的默认网关。
docker在创建一个容器时，会创建一对veth pair，一端置于容器中，一端置于docker0的虚拟网络中，从而实现网桥和容器的通信。
将容器中的veth(virtual ETHernet,将一个network namespace发出的数据包转发到另一个namespace)命名为eth0，docker0为网桥中的veth指定唯一的名字，从docker0中可用的地址段选取一个分配给容器，并配置默认路由。

同一个Node上Pod通信

poda的eth0->poda的veth->bridge0->podb的veth->podb的eth0

跨主机访问，网络插件cni：
flannel的三种实现方式：VXLAN；host-gw；UDP。

CNI的设计思想：kubernetes在启动infra容器后，可以直接调用CNI网络插件，为这个infra容器的network namespace配置符合预期的网络栈。

- UDP：性能差

操作系统将一个IP包发送给flannel0设备后，flannel0会把这个IP包交给创建这个设备的flannel进程，是内核态向用户态流动的方向。
fannel进程向flannel0设备发送了一个IP包，那么这个IP包会出现在宿主机的网络栈中，然后根据宿主机的路由表进行下一步处理，是用户态向内核态流动的方向。

三层overlay网络：首先对发出端的IP包进行UDP封装，然后在接收端进行解封装拿到原始IP包，进而把IP包转发给pod。

相比两台宿主机的直接通信，fannel UDP模式的容器多了flanneld的处理，在发出IP包的过程中，经过了三次内核态与用户态的数据拷贝。

* 第一次，用户态的容器进程发出的IP包经过docker0网桥进入内核态
* IP包根据路由表进入TUN(flannel0)设备，从而回到用户态flanneld进程
* flanneld进行UDP封包之后重新进入内核态，将UDP包通过宿主机的eth0发出去

- VXLAN(Virtual Extensible LAN)虚拟可扩展局域网：Linux内核本身支持的一种网络虚拟化技术。

设计思想：在现有的三层网络之上，“覆盖”一层虚拟的、由内核VXLAN模块负责的二层网络，使得连接在VXLAN二层网络上的“主机”之间可以像在同一个局域网(LAN)通信。

VXLAN会在宿主机上设置一个特殊的网络设备VTEP(VXLAN Tunnel End Point)虚拟隧道断点，对二层数据帧进行封装和解封装，这个执行流程是在内核中完成的

- host-gw：就是将每个flannel子网的“下一跳”设置成该子网对应的宿主机的IP地址

Calico提供的网络解决方案和flannel的host-gw几乎一样。不同于flannel通过etcd和宿主机上的flanneld来维护路由信息的做法，
calico通过BGP(Border Gateway Protocol边界网关协议)来自动在集群里分发路由信息。
calico不会再宿主机上创建任何网桥设备，只会为每个容器设置一个veth pair设备，然后把其中一端放置在宿主机上。所以calico的cni插件需要在宿主机上为每个容器的veth pair设备配置一条路由规则，用于接收传入的IP包。

+ 三层网络的优点：不需要封包、拆包，传递效率高，是可以设置复杂的路由规则
+ 隧道模式的优点：不需要在主机间的网关维护容器的路由信息，只需要主机有三层网络的连通性