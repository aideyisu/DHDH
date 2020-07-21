# DHDH
Diffie-Hellman with Gin

## 使用简介:
```
go build main.go
./main
```
随后http://localhost:8080/index 打开食用即可～

## Diffie-Hellman原理

#### 简介

它是一种安全协议，让双方在完全没有对方任何预先信息的条件下通过不安全信道建立起一个密钥，这个密钥一般作为“对称加密”的密钥而被双方在后续数据传输中使用。DH数学原理是base离散对数问题。做类似事情的还有非对称加密类算法，如：RSA

DH密钥协商算法主要用于解决密钥配送问题,本身并非用来加密

构造了一个复杂计算问题,使得该问题求解在现实时间内无法快速有效求解(和RSA原理有点像,rsa利用了大素数不可约性

#### 诞生原因

DH密钥协商算法在1976年在Whitfield Diffie和Martin Hellman两人合著的论文New Directions in Cryptography（Section Ⅲ PUBLIC KEY CRYPTOGRAPHY）中被作为一种公开秘钥分发系统(public key distribution system)被提出来

在该论文中实际上提出了一些在当时很有创新性的思想。原论文重点讨论两个话题：

(1)  在公网通道上如何进行安全的秘钥分派

(2)  认证(可以细分为消息认证和用户认证)

简而言之在不知是否安全的网络通道上进行通信

#### 详细流程

###### 背景 

Alice和Bob需要彼此协商一个密钥

###### 第一步

Alice和Bob共享一个素数p,以及素数的原根g (2 <= g <= p - 1)
pg可以明文传输,谁发给谁不重要,保证二者都知晓即可

Ps: 原根

QAQ学完一年多现在有点茶壶煮饺子...这两天看看概念尽量想一个简单表述方法写在这里.先委屈看到这里的你看看维基百科Otz抱歉!

https://zh.wikipedia.org/wiki/%E5%8E%9F%E6%A0%B9

###### 第二步

Alice生成一个私有随机数A. (1 <= A <= p - 1),计算YA = g^A (mod p), 将YA发送给Bob

Bob也同理生成私有随机数B...发送YB给Alice

此时Alice 此时知道 p,g,A,YA 其中A私有

Bob       此时知道 p,g,B,YB 其中B私有

###### 第三步

Alice 通过计算KA = YB^A (mod p) 得到密钥 KA

Bob   通过计算KB = YA^B (mod p) 得到密钥 KB

此时有 KA == KB 

原理:

KA = YB^A (mod p) = (g^B (mod p))^A (mod p) = g^(A * B) (mod p)

同理 KB = YA^B (mod p) = (g^A (mod p))^B (mod p) = g^(A * B) (mod p)

###### 总结

可见alice和bob生成密钥其实是相同运算过程,因此必然KA == KB.利用椭圆曲线进行密钥协商也是相同原理

更进一步,A和B不应该选择 p - 1.只能在 {1,2....,p-2}中选择

由费马小定理可知,选择p - 1会导致情况退化为 g^(p-1) 全等于 1 (mod p).对密钥协商机密性构成威胁

#### 缺陷

中间人攻击.Alice与Bob之间存在Monitor,被中间商截获 YA YB 从而在不被发现的情况下窃取消息明文

## docx内容:
1 介绍界面每个使用顺序，每一步干什么

2 介绍核心代码

## 项目路径一览
```
 DHDH
├── 学号姓名DH大作业.docx - 脱敏=。=
├── README.md
├── bin
└── src
    ├── aes
    │   ├── aes.go
    │   ├── aes128.go
    │   ├── aes256.go
    │   └── rijndael256.go
    ├── dh
    │   ├── handler.go
    │   └── restfulPrinter.go
    ├── go.mod
    ├── go.sum
    ├── main
    ├── main.go
    ├── pkg
    ├── services
    └── templates
        ├── index.html
        ├── info.html
        ├── pre.html
        ├── talk.html
        └── test.html

7 directories, 17 files
```

