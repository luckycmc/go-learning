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

