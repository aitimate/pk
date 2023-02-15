# pk
```cmd
USAGE:
   pk [global options] command [command options] [arguments...]

GLOBAL OPTIONS:
   --decrypt, -d, -D         Is the decrypt mode? or default encrypt mode (default: false)
   --key value, -k value     The key for AES process
   --value value, -v value   The text value that you want to encrypt or decrypt
   --output value, -o value  The file path for the output text (default: "./.env")
   --help, -h                show help
```
## 用法1 基于.exe进行加密
> 说明: 仓库根目录下的pk.exe下载即可使用，样例如下: 
```cmd
.\pk.exe --key a475a27Bb1028f140bc2a7c843318afD --value "This is a secret!"
# .\pk.exe -k a475a27Bb1028f140bc2a7c843318afD -v "This is a secret!"
```
## 用法2 基于`go install`cmd命令进行加密
> 说明: 你需要将%GOPATH%\bin已添加到Path环境变量，因为生成的.exe程序会位于该目录下

1. 下载源码并编译
```cmd
git clone https://github.com/aitimate/pk.git && cd pk && go install
```
2. 运行示例(CMD)
```cmd
pk --key a475a27Bb1028f140bc2a7c843318afD --value "This is a secret!"
pk -k a475a27Bb1028f140bc2a7c843318afD -v "This is a secret!"
```
## [小技巧] `//go:generate`快捷执行
> 说明: `//go:generate`可以在Go文件注释位置显示一键执行图标，用法示例

```js
//go:generate ./pk.exe --key a475a27Bb1028f140bc2a7c843318afD --value "This is a secret!"
```



