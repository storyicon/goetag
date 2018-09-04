# goetag [![Build Status](https://travis-ci.org/storyicon/goetag.svg?branch=master)](https://travis-ci.org/storyicon/goetag) [![Go Report Card](https://goreportcard.com/badge/github.com/storyicon/goetag)](https://goreportcard.com/report/github.com/storyicon/goetag)

[English Document](#English)  
[中文文档](#Chinese)

<h2 id="English">English Document</h2>    
Package etag is used to calculate the hash value of the file and string.      

The algorithm is based on the qetag of Qi Niuyun.     

Qetag address is          

    https://github.com/qiniu/qetag          
    
This package extends two export functions based on Qetag, named `GetEtagByString`, `GetEtagByBytes`  

And re-implemented `GetEtagByPath`          

> The hash value generated by this package is exactly the same as the hash value generated by the Qi Niuyun.      

```
package main
import (
    "log"
    "github.com/storyicon/goetag"
)
func main(){
    var  (
        conseq string
        err error
    )
    conseq, err = goetag.GetEtagByString("test")
    log.Println("StringHash:", conseq, err)

    conseq, err = goetag.GetEtagByBytes([]byte("test"))
    log.Println("BytesHash:", conseq, err)


    conseq, err = goetag.GetEtagByPath("./LICENSE")
    log.Println("FileHash:", conseq, err)
}
```

Output：

```
2018/09/04 12:22:58 StringHash: FqlKj-XMsZumHEwIc9OR6YeYL7vT <nil>
2018/09/04 12:22:58 BytesHash: FqlKj-XMsZumHEwIc9OR6YeYL7vT <nil>
2018/09/04 12:22:58 FileHash: FrMUx-u31ZmUSYGQi38-0zow5486 <nil>
```

### Algorithm interpretation：

For files or strings smaller than 4M, use the formula:

`Hash = UrlsafeBase64([0x16, sha1(content)])`

Formula explanation: Add a byte with a value of 0x16 before the SHA1 value (20 bytes) of the string to form 21 bytes of binary data, and do base64 encoding of urlsafe for these 21 bytes.

For files or strings larger than 4M, the calculation formula is used:
`Hash = UrlsafeBase64([0x96, sha1([sha1(Block1), sha1(Block2), ...])])`
Note: The `Block` is a separate block that divides the contents of the file or string into 4M units, ie Blockj = FileContent[j*4M:(j+1)*4M]

### Multiple programming languages for this algorithm

    https://github.com/qiniu/qetag

<h2 id="Chinese">中文文档</h2>
goetag用于计算文件和字符串的哈希值，该算法基于七牛云的qetag。    

Qetag地址是 https://github.com/qiniu/qetag    

基于Qetag的算法, goetag扩展了两个导出方法，名为`GetEtagByString`，`GetEtagByBytes`，并重新实现了`GetEtagByPath`    

> 本Package生成的哈希值与七牛云生成的哈希值是完全一样的

```
package main
import (
    "log"
    "github.com/storyicon/goetag"
)
func main(){
    var  (
        conseq string
        err error
    )
    conseq, err = goetag.GetEtagByString("test")
    log.Println("StringHash:", conseq, err)

    conseq, err = goetag.GetEtagByBytes([]byte("test"))
    log.Println("BytesHash:", conseq, err)


    conseq, err = goetag.GetEtagByPath("./LICENSE")
    log.Println("FileHash:", conseq, err)
}
```

输出：

```
2018/09/04 12:22:58 StringHash: FqlKj-XMsZumHEwIc9OR6YeYL7vT <nil>
2018/09/04 12:22:58 BytesHash: FqlKj-XMsZumHEwIc9OR6YeYL7vT <nil>
2018/09/04 12:22:58 FileHash: FrMUx-u31ZmUSYGQi38-0zow5486 <nil>
```

### 算法解释：

对于体积小于 4M 的文件或字符串，采用计算公式：  
`hash = UrlsafeBase64([0x16, sha1(content)])`  
公式解释：在字符串的 SHA1 值（20 字节）前加一个值为 0x16 的 byte，构成 21 字节的二进制数据，并对这 21 字节做 urlsafe 的 base64 编码

对于体积大于 4M 的文件或字符串，采用计算公式：  
`hash = UrlsafeBase64([0x96, sha1([sha1(Block1), sha1(Block2), ...])])`  
注：其中的`Block`是把文件或字符串内容切分成 4M 为单位的独立块，即 Blockj = FileContent[j*4M:(j+1)*4M]

> 在 sha1 值前面加一个 byte 的标志位（0x16 或 0x96）是因为：  
> 0x16=22, 而 2^22=4M, 故标志位 0x16 表示文件是按 4M 分块  
> 0x96=0x80|0x16,0x80 表示文件是有多个分块的大文件，hash 值经过了两重 sha1 计算

### 本算法的其他编程语言实现

    见：https://github.com/qiniu/qetag
