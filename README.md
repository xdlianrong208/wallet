# 匿名计算链钱包

匿名计算链钱包为匿名计算链的一个子项目，服务于用户，用于安全计算过程和匿名交易。钱包每个用户单独使用，无中心化的网络服务器参与，完成用户使用匿名计算链的基本功能。

![wallet](/img/wallet.png)

## 功能模块

钱包包括一个初始化模块和四个功能模块，其中购币交易和转账交易是钱包的核心功能模块。

- 初始化
- 购币交易（兑换积分）
- 转账交易
- 收款
- 财务显示

## 初始化

初始化模块提供两个功能：新建钱包和加载钱包

- 新建钱包：指定一个本地路径**持久化保存**钱包，同时加载、进行数据读写。（可以使用单个数据文件存储）。

- 加载钱包：制定一个本地路径加载已有钱包，进行数据读写。

## 购币交易（兑换）

受到用户的兑换请求后，前端将公钥、金额参数传给钱包后端，钱包后端和交易所接口交互，兑换积分，在收到交易所的承诺后还要验证下是否是有效承诺（用CM_o, o, r），并把结果返还给前端，同时在用户钱包生成**一条账本数据**（保存CM_o和金钱value）。

## 转账交易（链上交互）

用户选择使用的承诺，使用的金额和找零金额，填写接受者，和链上提供的接口交互（解锁账户、发起交易），发起转账，生成**转账交易记录**( v r s 和交易哈希)

交互方式：向链上发起一个 HTTP 请求，返回转账交易记录。

## 收款

- 接收者接收转账交易：用户填写交易的哈希、交易的接收者的公钥，用哈希拿取交易，用接收者私钥解密，获取其中的转账金额，存到自己的钱包里。

- 发送者拿到找零金额：用户填写交易的哈希，系统上链中找交易哈希是否存在，如果存在，和用户钱包转账交易记录比对，如果有相同交易哈希，则判定转账交易已经上链。钱包拿取转账交易记录，删去用户花费的承诺，新存取找零金额到自己的钱包里。

## 财务显示（前端）

将当前用户的钱包内财务数据加载到前端界面可视化，采用 Vue 方案。

## FAQ

- 本钱包采用中心化方案，私钥及财产数据都保存在平台方。
- 业界的区块链钱包保存用户的地址，由于匿名计算链并不具有地址，而是使用一条财产数据来替代（财产数据格式：随机数+公钥）。

## 关联项目
- [Geth-Comments](https://github.com/xdlianrong/Geth-Comments)