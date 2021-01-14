# 钱包接口

包含了钱包需要与区块链交互的所有接口，见 require.go

### 1.UnlockAccount

功能说明：解锁以太坊上的账户，在发送任何交易前必须要做。

输入：

```go
ethaccount  string // 以太坊账户地址
ethkey      string // 以太坊账户密码
```

输出：

- 解锁成功：ture
- 解锁失败：false

应用：在钱包视图发起转账交易前使用。

### 2.SendTransaction

功能说明：向以太坊发送转账交易

输入：

```go
spk string // 发送者公钥 g1+g2+p+h的拼接
rpk string // 接收者公钥 g1+g2+p+h的拼接
s   string // 发送金额
r   string // 找零金额
vor string // 使用本次货币承诺的随机数
cmo string // 使用本次货币的承诺
```

输出：

- 解锁成功：ture
- 解锁失败：false

应用：在钱包视图发起转账交易时使用。

### 3.GetTransaction

功能说明：获取交易内容

输入：

```go
txhash string // 交易的hash值
```

输出：

- 解锁成功：ture
- 解锁失败：false

应用：接收者接收转账金额，发送者接收找零金额时使用。







