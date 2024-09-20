# go-learning
the note for learning golang

## 线程
### 内核态线程
### 用户态线程（协程）
#### 用户态线程是绑定在内核态线程上的，CPU并不知道用户态线程的存在，只知道它运行的是内核态线程（Linux的PCB进程控制块）
### 映射关系
#### N:1（N个协程绑定一个线程）
##### 优点：协程在用户态线程即完成切换，不会陷入到内核态，切换非常轻量快速
##### 缺点：一个进程的所有协程都绑定在一个线程上，某个程序用不了硬件的多核加速能力；一旦某个协程阻塞，造成线程阻塞，本进程的其他协程都无法执行了，根本就没有并发能力了
#### 1:1（一个协程绑定一个线程，最容易实现）
##### 优点：协程的调度都由CPU完成，不存在N:1的缺点
##### 缺点：协程的创建、删除和切换的代价都由CPU完成，有点浪费
#### M:N关系（M个协程绑定N个线程）
##### 优点：是N:1和1:1类型的结合，克服了前两种模型的缺点
##### 缺点：实现最复杂
### 协程跟线程的区别
#### 线程由CPU调度，是抢占式的；协程由用户态调度，是协作式的，一个协程让出CPU后，才执行下一个协程

## goroutine调度
## GMP
### G: goroutine协程
### P: Processor处理器
#### 包含了goroutine的资源，如果线程想运行goroutine，必须先获取P，P中还包含了可运行的G队列
### M: Thread线程
### 在Go中，线程是运行goroutine的实体，调度器的功能是把可运行的goroutine分配到工作线程上
1. 全局队列(Global Queue)：存放等待运行的G
2. P的本地队列：和全局队列类似，存放的也是等待运行的G，存的数量有限，不超过256个。新建G'时，G'优先加入到P的本地队列，如果队列满了，则会把本地队列中一半的G移动到全局队列。
3. P列表：所有的P都在程序启动时创建，并保存在数组中，最多有**GOMAXPROCS**（可配置）个
4. M：线程想运行任务就得获取P，从P的本地队列获取G，P队列为空时，M也会尝试从全局队列拿一批G放到P的本地队列，或从其他P的本地队列**偷**一半放到P的本地队列。M运行G，G执行之后，M会从P获取下一个G，不断重复。
### Goroutine调度器和OS调度器是通过M结合起来的，每个M都代表了一个内核线程，OS调度器负责把内核线程分配到CPU的核上执行

### P和M的个数问题
1. P的数量
   * 由启动时环境变量$GOMAXPROCS或者是由runtime的方法GOMAXPROCS()决定。意味着在程序执行的任意时刻都只有$GOMAXPROCS个goroutine在同时执行
2. M的数量
   * go语言本身的限制：go程序启动时，会设置M的最大数量，默认10000，但是内核很难支持这么多线程数，所以该限制可以忽略
   * runtime/debug中的SetMaxThreads函数，设置M的最大数量
   * 一个M阻塞了，会创建新的M

M与P的数量没有绝对关系，一个M阻塞，P就会去创建或者切换到另一个M，所以，即使P的默认数量是1，也有可能会创建多个M出来
### P和M何时被创建
1. P何时创建：在确定了P的最大数值n后，运行时系统会根据这个数量创建n个P
2. M何时创建：没有足够的M来关联P并运行其中的可运行的G。比如所有的M此时都阻塞住了，而P中还以后很多就绪任务，就会去寻找空闲的M，如果没有空闲的，就会去创建新的M

## 调度器设计策略
### 复用线程：避免频繁地创建、销毁线程，而是对线程的复用
1. work stealing机制
当本线程无可运行的G时，尝试从其他线程绑定的P偷取G，而不是销毁线程
2. hand off机智
当本线程因为G进行系统调用阻塞时，线程释放绑定的P，把P转移给其他空闲的线程执行
**利用并行：** GOMAXPROCS设置P的数量，最多有GOMAXPROCS个线程分布在多个CPU上同时运行。GOMAXPROCS也限制了并发的速度，比如GOMAXPROCS = 核数/2，则最多利用了一版的CPU核进行并行
**抢占：** 在coroutine中要等待一个协程主动让出CPU才执行下一个协程，在Go中，一个goroutine最多占用CPU10ms,防止其他goroutine被饿死，这就是goroutine不同于coroutine的一个地方
**全局G队列：** 在新的调度器中依然有全局G队列，但功能已经被弱化了，当M执行work stealing从其他P偷不到G时，它可以从全局G队列获取G

### go func()调度流程