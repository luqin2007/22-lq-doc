可用工具：
- `md5sum`
- `sha1sum`
- `sha256sum`
- `sha384sum`
- `sha512sum`

以上命令可接受一个输入流，返回 Hash 值和文件名（非文件则为 `-`）

```shell
echo <content> | <tools>
```

将文本内容 `<content>` 通过管道传递给目标方法

![[../../../_resources/images/Pasted image 20241128160548.png]]

```shell
<tools> <file>
```

使用目标方法计算给定文件内容的 Hash 值

![[../../../_resources/images/Pasted image 20241128160626.png]]