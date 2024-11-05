---
title: "手把手教大家做包（超级小白） 来自 是喵喵丫 - 酷安"
source: "https://www.coolapk.com/feed/59655705?shareKey=ZWE1NDQ2YjQ0OWJiNjcxNWM0MjE~"
author:
published:
created: 2024-11-06
description:
tags:
  - "clippings"
---
此方法必须要root，没root试试虚拟机，没尝试过。  
此教程，以11u/Pro为例，其他机型做包方法也应该通用。

1.官改包的制作教程  
首先我们要有做包工具，小白做官改不推荐DNA，最好用hais工具来制作。

![](https://image.coolapk.com/feed/2024/1019/18/21971321_bf9758ae_2999_9391_354@579x301.jpeg.m.jpg)

我们工具下载的都是这种压缩的包，我们要将他解压在data这个路径下

![](https://image.coolapk.com/feed/2024/1019/18/21971321_65a84453_2999_94_312@1080x2400.jpeg.m.jpg)

执行Linux.sh这个脚本，记得给root权限。

![](https://image.coolapk.com/feed/2024/1019/18/21971321_4d2876b0_2999_9406_672@1080x2400.jpeg.m.jpg)

执行文件的时候，它会提示下什么文件，下完之后是这样子的。

![](https://image.coolapk.com/feed/2024/1019/18/21971321_8b3b8654_2999_9415_288@1080x2400.jpeg.m.jpg)

我们要将要做的包移动到OriginalPackage这个文件下

![](https://image.coolapk.com/feed/2024/1019/18/21971321_3117a79a_2999_9423_858@1080x2400.jpeg.m.jpg)

![](https://image.coolapk.com/feed/2024/1019/18/21971321_de36d8ea_2999_9432_357@1080x2400.jpeg.m.jpg)

![](https://image.coolapk.com/feed/2024/1019/18/21971321_0a043fb3_2999_9437_787@1080x2400.jpeg.m.jpg)

接着我们再打开工具，创建工程

![](https://image.coolapk.com/feed/2024/1019/18/21971321_d9e35fed_2999_9441_526@1080x2400.jpeg.m.jpg)

![](https://image.coolapk.com/feed/2024/1019/18/21971321_2c1ae720_2999_945_886@1080x2400.jpeg.m.jpg)

我们要等待一些时间，第一次他会下一些东西。

![](https://image.coolapk.com/feed/2024/1019/18/21971321_83cf80e4_3003_9021_913@1080x2400.jpeg.m.jpg)

解压好了会是这个界面。

![](https://image.coolapk.com/feed/2024/1019/18/21971321_8a171110_3003_9025_687@1080x2400.jpeg.m.jpg)

一定要去除avb校验，否则百分之百开不了机，选40就可以。

![](https://image.coolapk.com/feed/2024/1019/18/21971321_a5663a1f_3003_9028_32@1080x2400.jpeg.m.jpg)

如果自己想要精简什么，我们可以选择退出。

![](https://image.coolapk.com/feed/2024/1019/18/21971321_1ae165a6_3003_9035_73@1080x2400.jpeg.m.jpg)

选HaisWorker这个文件

![](https://image.coolapk.com/feed/2024/1019/18/21971321_68b17121_3003_9042_444@1080x2400.jpeg.m.jpg)

点开你分解的包。MIUI14以下大部分软件在/system/system/里面 MIUI14级以上大部分软件在product。

![](https://image.coolapk.com/feed/2024/1019/18/21971321_d329d668_3003_9051_563@1080x2400.jpeg.m.jpg)

![](https://image.coolapk.com/feed/2024/1019/18/21971321_2a933706_3003_9053_521@1080x2400.jpeg.m.jpg)

千万别乱删，否则会卡米，如果追求稳定最好用脚本自带的。  
如果你要做的包是MIUI12.5及以下，或MIUI13以上（不包括MIUI13）都要破解卡米或system校验  
注意:再次进入脚本时候要看好选择进入工程，而不是创建工程，选择创建工程的话，你之前做的都毁掉了。

![](https://image.coolapk.com/feed/2024/1019/18/21971321_2b0c8aeb_3003_9056_214@1080x2400.jpeg.m.jpg)

破解卡米和system校验选30就行

![](https://image.coolapk.com/feed/2024/1019/18/21971321_d078d905_3003_9062_54@1080x2400.jpeg.m.jpg)

做完一切后就可以选90打包了。

![](https://image.coolapk.com/feed/2024/1019/18/21971321_d57b1d06_3007_3063_266@1080x2400.jpeg.m.jpg)

弄好的包一般都在Complete这个文件里面。  
这样一个小白官改包就做好了。  
################################

2.移植包制作教程  
小白可以先选同系列的先下手  
以11u为例，可以选移植11青春版和11清活版等。一般都可以直接开机。

我一般都是用syt和dna来做移植包  
以syt为例，

![](https://image.coolapk.com/feed/2024/1019/18/21971321_a48346f0_3007_3067_838@1080x2400.jpeg.m.jpg)

我们要继续用到mt，也是在data这个文件下面找到  
syt-tool这个文件点开，同样也是root运行。

![](https://image.coolapk.com/feed/2024/1019/18/21971321_1f37a3fc_3007_3076_982@1080x2400.jpeg.m.jpg)

然后要创建一个文件，将包移进去。

![](https://image.coolapk.com/feed/2024/1019/18/21971321_86006e10_3007_3081_768@1080x2400.jpeg.m.jpg)

然后选择解压zip，之后再解压plyload这样我们会得到几个文件

![](https://image.coolapk.com/feed/2024/1019/18/21971321_1b82301c_3007_3091_290@1080x2400.jpeg.m.jpg)

将原包也就是自己机型的vendor.img，和odm.img（老机型是那些br文件，要先合并成img再进行替换）替换到自己要开的包那里，把out整个文件和boot顺便也替换过去，方便之后的打包。

![](https://image.coolapk.com/feed/2024/1019/18/21971321_ba34a646_3007_3102_299@1080x2400.jpeg.m.jpg)

之后我们选三解压br/dat/img（原包的和移植包的都要解）

![](https://image.coolapk.com/feed/2024/1019/18/21971321_e9270dac_3007_3107_941@1080x2400.jpeg.m.jpg)

MIUI12.5的机型文件在vendor.img里面自己慢慢找。  
MIUI12.5以上的在/product/etc/device\_features/这里面打开device\_features将本机的机型文件复制到要移植的device\_features文件里

![](https://image.coolapk.com/feed/2024/1019/18/21971321_2036c447_3007_3116_101@1080x2400.jpeg.m.jpg)

也要将官方的displayconfig里面的文件复制到另一个要移植的displayconfig里面。

![](https://image.coolapk.com/feed/2024/1019/18/21971321_bb24068b_3007_3125_231@1080x2400.jpeg.m.jpg)

在product/overlay/找到DevicesAndroidOverlay.apk和DevicesOverlay.apk将原机的复制到同个路径下要移植的包里面。  
相机也要替换，否则可能会出现相机打不开的情况。

![](https://image.coolapk.com/feed/2024/1019/18/21971321_6529c589_3011_0478_625@1080x2400.jpeg.m.jpg)

我们还要打开移植包的product/etc/文件的build.prop文件MIUI13可能在system/system的build.prop的文件下。

![](https://image.coolapk.com/feed/2024/1019/18/21971321_9199dbde_3011_0483_783@1080x2400.jpeg.m.jpg)

把屏幕dpi改成本机的。

如果移植的系统是12.5要破解卡米，MIUI14级以上要破解system校验  
解除Android 13及以上的签名校验的教程是  
找到services.jar  
搜索getMinimumSignature  
把  
.line 1114

invoke-static {p0}, Landroid/util/apk/ApkSignatureVerifier;->getMinimumSignatureSchemeVersionForTargetSdk(I)I

move-result p0  
修改为  
.line 1114  
const/4 p0, 0x0  
修改另一条  
line 345  
iget v0, v0, Landroid/content/pm/ApplicationInfo;->targetSdkVersion:I

invoke-static {v0}, Landroid/util/apk/ApkSignatureVerifier;->getMinimumSignatureSchemeVersionForTargetSdk(I)I

move-result v0  
修改为  
.line 345  
iget v0, v0, Landroid/content/pm/ApplicationInfo;->targetSdkVersion:I

const/4 v0, 0x0  
下一条  
.line 696  
invoke-static {p1}, Landroid/util/apk/ApkSignatureVerifier;->getMinimumSignatureSchemeVersionForTargetSdk(I)I

move-result p1  
修改为  
.line 696  
const/4 p1, 0x0  
下一条  
.line 3123  
invoke-static {p5}, Landroid/util/apk/ApkSignatureVerifier;->getMinimumSignatureSchemeVersionForTargetSdk(I)I

move-result p5  
修改为  
.line 3123  
const/4 p5, 0x0

或者  
找/system/framework/framework.jar，在其dex中搜索“getMinimumSignature”，应当可以找到一个名为"getMinimumSignatureSchemeVersionForTargetSdk"的函数实现，找到其中"const/4 v0, 0x2"这行，把"2"改为"1"。保存，重编译，替换回去即可。注意不要重新签名

![](https://image.coolapk.com/feed/2024/1019/18/21971321_46c40fac_3011_0489_391@1080x2400.jpeg.m.jpg)

推荐第一种方法

12.5破解卡米的教程是  
system/system/framework/找到services.jar，点击查看，classes.dex文件选择DEX++编辑器，并全选dex文件。

依次选择com→miui→server→SecurityManagerService

![](https://image.coolapk.com/feed/2024/1019/18/21971321_8af4b1fd_3011_0499_896@3240x2400.png.m.jpg)

搜索method private checkSystemSelfProtection(Z)V，将line里面的内容都删除，之后保存退出

![](https://image.coolapk.com/feed/2024/1019/18/21971321_297e5adf_3011_0501_37@3240x2400.jpeg.m.jpg)

以上方法来源于网络……  
接下来就可以打包了。如果卡一屏清除数据试试，如果一屏都没进可能avb没去，也可能包本身就有问题。切记做包有成🧱风险，请自行承担。  
小白做做官改就行了，移植包让大佬去做因为大佬可以用dsu来测试![笑哭](http://static.coolapk.com/emoticons/v9/coolapk_emotion_31_xiaoku.png)，不想写dsu测试教程了。注意文件千万别刷错，否则9008伺候  
成品:

![](https://image.coolapk.com/feed/2024/1019/18/21971321_c3139a3c_3011_0504_908@1736x2412.jpeg.m.jpg)

![](https://image.coolapk.com/feed/2024/1019/18/21971321_5e88a6a2_3011_051_209@1736x2412.jpeg.m.jpg)

再好的安卓之光也处处遗憾，火龙888也烧不透你冰封的心，纯洁的白色后盖也留不住花心的你，67w闪充也充不进你断电的心，X轴震动马达也震不动绝情的你，5000万超光感主摄也拍不美自我毁容的你，杜比立体双扬也藏不住我们的秘密，再好的哈曼卡顿也放不出我们的回忆，万能的红外遥控也遥控不了我们最初的距离，全场景NFC也刷不动不带NFC的你。  
[#ProjectTreble#](https://www.coolapk.com/t/ProjectTreble?type=12) [#小米11#](https://www.coolapk.com/t/%E5%B0%8F%E7%B1%B311?type=12) [#HyperOS#](https://www.coolapk.com/t/HyperOS?type=12)