# 微服务构建教程-CI/CD部分

持续集成（Continuous Intergration/CI) 使我们能在快速干净的环境中创建与测试我们的应用。

持续交付（Continuous Delivery/CD）可以帮助我们发布通过测试的安全代码。

# 1 创建Github仓库

1. 在Github上创建一个仓库
2. 将本地的代码上传到Github

在文件夹目录中执行如下代码

```go
$ git init
$ git remote add origin YOUR_REPOSITORY_URL
$ git add -A
$ git commit -m "initial commit"
$ git push origin main
```

# 2 添加Semaphore管理

1. 注册一个[Semaphore账号](https://semaphoreci.com/)，直接用github账户登录即可
2. 点击create new，选择关联刚刚的github项目（这里要先将Github和Semaphore做一个连接）

![Untitled](pic/Untitled.png)

![Untitled](pic/Untitled%201.png)

1. 选择Go工作流，并点击customize

![Untitled](pic/Untitled%202.png)

1. 将Go版本sem-version改为自己的版本，我的版本是1.19

![Untitled](pic/Untitled%203.png)

1. 点击run the workflow，start

![Untitled](pic/Untitled%204.png)

1. 成功启动，现在每次git push, CI pipeline都会启动测试

![Untitled](pic/Untitled%205.png)

# 3 改进pipeline

## 3.1 builder组件

![Untitled](pic/Untitled%206.png)

CI的builder主要包含以下组件：

- **pipeline:** 一个pipeline有一个共同的目标（如：测试），pipeline由blocks构成，在agent中从左到右执行。
- **agent:** agent是为pipeline提供动力的虚拟机器，该机器运行带有多种语言构建工具的unbuntu20.04镜像
- **block:** block可以将执行的job进行分组，一个block中的job通常有类似的命令和配置。当一个block中的所有job完成，下一个block开始运行。
- **job:** job中定义了执行工作的命令，他们从父block中继承配置。

## 3.2 改进pipline

我们在默认的CI pipeline上做一些改进：

- 每次集成都要重新下载依赖项，因此我们使用缓存来保存所有的依赖项
- 测试和构建在同一个job里，我们应该将他们分成不同的部分，方便之后添加测试

### 3.2.1 使用缓存保存依赖项

点击首页的`edit workflow` ，执行如下编辑

1. 点击第一个block，将其名字改为install dependencies
2. 往下拉将该block中job的名字改为install，并在框中输入如下命令

```python
sem-version go 1.19
export GO111MODULE=on
export GOPATH=~/go
export PATH=/home/semaphore/go/bin:$PATH
checkout
cache restore
go mod vendor
cache store
```

### 3.2.2 添加测试与构建block

1. 点击**+Add Block** 并将该block命名为Test
2. 在prologue框中填入以下命令，这些命令会在block执行之前执行

```python
sem-version go 1.19
export GO111MODULE=on
export GOPATH=~/go
export PATH=/home/semaphore/go/bin:$PATH
checkout
cache restore
go mod vendor
```

1. 将job命令为Test，填入以下命令

```python
go test ./...
```

1. 再新增一个block，命名为Build
2. 同样在prologue框中填入相同的命令
3. job命令为Build,填入以下命令

```python
go build -v -o go-gin-app
artifact push project --force go-gin-app
```

点击Run the workflow并启动，如图所示

![Untitled](pic/Untitled%207.png)

最后，生成的二进制可执行文件将保存再Artifacts按钮中，可以点击查看

![Untitled](pic/Untitled%208.png)