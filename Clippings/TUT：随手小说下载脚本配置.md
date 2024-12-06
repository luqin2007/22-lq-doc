---
title: "随手小说下载"
source: "https://greasyfork.org/zh-CN/scripts/451039"
author:
published:
created: 2024-12-06
description: "带图形化界面的小说下载器。任意网站，自动识别任何目录列表，自动识别正文，自由下载，简单直观。"
tags:
  - "clippings"
---
配置信息内部以json格式存放。默认不会自动保存，需要时请手动保存。在显示框显示的是本脚本规定的格式，规则如下:

开始字符以$$打头的，则后面字符认为是key，接下去的行则为这个key所对应的值。空行自动删除。

允许的key值:

- listSelector

表示链接列表选择器，多行则每一行表示一个选择器，提取时会依次合并。

当每页的目录选择器都不一样时，可以设置该值为 --，使获取剩余目录时都自动获取目录选择器。

- textSelector

表示正文选择器，多行会依次提取并合并正文。

当每章节的正文选择器都不一样时，可以设置该值为 --，使获取每章节正文时都自动获取正文选择器。

- jammerSelector

干扰码选择器，提取正文前会检索符合的元素，多行会依次删除

- nextListSelector

目录下一页选择器

- nextListDelay

目录下一页延迟时间，毫秒

- frameText

针对动态网页，利用iframe加载获取正文。单独一行，不需要设置值。只是这种方式获取文本会比较慢。个别网站会跳到章节页面，不能用。

- addedNextPageReg

自定义下一页链接的文本正则表达式。

- notUploadSaveinfo

不要上传正文保存信息。单独一行，不需要设置值。

以下的key为自定义函数，统一的行为是如果函数没有返回值，则默认会返回并更新传入的各个参数。

- customListFunc

自定义列表函数。会在列表提取前执行，参数为doc，selector。表示目录页的document和目录选择器。如果返回参数，则应该是一个对象数组表示目录列表，每个对象为｛href:xxx;text:xxx｝表示一个目录项的网址和目录名。

- customItemFunc

自定义目录链接处理函数，会在处理每一个目录链接前执行。传入doc，item，表示目录页的document和待处理的某个目录链接元素。

- customTextFunc

正文处理函数，会在每一个目录的正文处理前执行。传入doc，selector，表示正文页的document和正文选择器。如果返回值为单个文本，则直接把该文本当做正文。

- customNextPageFunc

自定义获取正文下一页链接。传入doc，href，表示当前正文页的document和网址。函数应返回下一页网址。