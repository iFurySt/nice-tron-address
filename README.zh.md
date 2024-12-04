# Tron牛逼地址生成器

这个工具可以生成带有特定后缀的Tron地址。比如：
```plaintext
TBvjAaMRHwvzFbtM3GvxC9mg8888888888
```

## 使用教程
到这里下载最新版本：[releases page](https://github.com/iFurySt/nice-tron-address/releases).

运行二进制文件并传入参数

### 生成一个带有后缀`888`的地址:
```bash
./tron-addr-gen --suffix 888 --num 1
```

### 生成5个带有后缀`abc`的地址:
```bash
./tron-addr-gen -s abc -n 5
```

## 输出示例
```plaintext
Private Key: e3d9b4c8f2f5a...bcdef, Address: TBvjAaMRHwvzFbtM3GvxC9mg8888888888
```

## 证书
本项目基于 **Creative Commons Legal Code CC0 1.0 Universal** 协议。更多信息请参考 [LICENSE](LICENSE)

注：除非另有明确声明，本协议项下的作品关联人，在可适用法律所允许的最大限度内，不对本作品提供担保，不承担因本作品使用产生的责任。
