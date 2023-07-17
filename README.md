# 獬豸

### golang实现的文档相似度检测工具

## 安装方法
```go
go get https://github.com/cb252389238/xiezhi
```

## 使用方法

```go
杰卡德系数检测 返回值0-1之间，越靠近1相似度越高
jaccard := xiezhi.Jaccard(text1, text2)

余弦相似度算法 返回值0-1之间，越靠近1相似度越高
jaccard := xiezhi.CosineSim(text1, text2)

minhash算法
xiezhi.MinHash(text) //返回hash签名[]uint32类型
xiezhi.MinHashSimilarity(hash1,hash2) //对比两个hash值的相似度

simhash算法
xiezhi.SimHash(text) //返回hash值 uint64
xiezhi.SimHashSimilarity(hash1,hash2) // 返回海明距离和相似度  一般海明距离小于等于3则认为是相似文档

```