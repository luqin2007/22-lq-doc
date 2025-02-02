---
title: "PhotoPrism - 可私有部署，基于机器学习 TensorFlow 的图像自动分类、开源照片管理工具，类本地化 Google Photos 服务 - 小众软件"
source: "https://www.appinn.com/photoprism/"
author:
  - "[[青小蛙]]"
  - "[[青小蛙]]"
published: 2021-05-30T12:05:45+08:00, 2021-05-30T12:05:45+08:00, 2021-05-30T12:05:45+0800
created: 2024-11-24
description: "PhotoPrism 是一款开源的，基于机器学习软件 Google TensorFlow 的开源照片管理工具，它能够实现自动图像分类，可检测颜色、色度、亮度、质量、全景照片、位置等图片属性，支持直接上传照片，也可以挂载 WebDAV 目录，Docker 傻瓜式安装方式，Web"
tags:
  - "clippings"
---
> [!danger] 该应用仅可以通过 Docker 安装

**PhotoPrism** 是一款开源的，基于机器学习软件 Google TensorFlow 的开源照片管理工具，它能够实现自动图像分类，可检测颜色、色度、亮度、质量、全景照片、位置等图片属性，支持直接上传照片，也可以挂载 WebDAV 目录，Docker 傻瓜式安装方式，Web 界面访问和管理照片，可以在家中或服务器上托管。@[Appinn](https://www.appinn.com/photoprism/)

![PhotoPrism - 基于机器学习 TensorFlow 的自动图像分类、开源照片管理工具](https://static2.appinn.com/images/202105/photoprism.jpg)

PhotoPrism 有点像本地化的 Google Photos，虽然没那么强大，但已经有点那个意思了。而开发者也强调 PhotoPrism 就是一个可以不使用云端相册的解决方案，它是一个私人托管的应用程序，用于浏览、组织和分享您的照片集。它利用最新的技术，在不妨碍你的情况下自动标记和查找图片。

## PhotoPrism 示例

PhotoPrism 提供了一个官方示例，里面已经有了不少照片，以及自动分类，包括标签、地理位置，并且还能识别黑白照片、全景照片、相似照片（Stacks），直接访问下面的地址就能看到：

- [https://demo.photoprism.org/](https://demo.photoprism.org/)

支持中文界面，需要在左下角设置中，将语言修改为简体中文即可。

在左侧菜单中，你可以找到相册、硬盘、我的最爱、瞬间（自动生成的相册）、日历、地点、标签（自动生成）、文件夹等内容：

![PhotoPrism - 基于机器学习 TensorFlow 的自动图像分类、开源照片管理工具](https://static2.appinn.com/images/202105/screen-appinn2021-05-30_11_24_59.jpg)

PhotoPrism 的设置内容也很丰富，有多种主题界面，权限设置，可以完全私用，也可以只读提供给别人观看，或者做一个公共相册，让任何人（或有权限）上传照片。

拥有三种上传图片的方式：

- 原生储存（部署时直接挂载在 Docker 里）
- 浏览器上传
- WebDAV

![PhotoPrism - 基于机器学习 TensorFlow 的自动图像分类、开源照片管理工具](https://static2.appinn.com/images/202105/screen-appinn2021-05-30_11_26_20.jpg)

## PhotoPrism 部署

相对于其他类似的项目，PhotoPrism 的部署可能是最简单的一个了。官方提供了基于 docker-compose 的部署模板，竟然简单到完全不需要修改，直接安装：

| 1  2 | `wget https:``//dl``.photoprism.org``/docker/docker-compose``.yml`  `docker-compose up -d` |
| --- | --- |

等 30 秒，就能用 ip:2342 访问了。当然，前提是安装好了 docker 和 docker-compose（青小蛙提供了一个 Ubuntu 下的[简明教程](https://meta.appinn.net/t/topic/3284)）。

如果想要定制化 PhotoPrism，就需要[阅读文档](https://docs.photoprism.org/getting-started/)了。

当然，一切私有化部署都有技术门槛，这也是为什么公共云服务不可被取代。想要部署 PhotoPrism 至少需要一定的电脑操作能力和一台运行 Docker 的服务器（可以是普通电脑），这在一切皆手机的时代，已经拒绝了大部分人。还好这个年代的公有云也提供了廉价的电脑（没错，你在公有云购买的服务器，实际上可以理解为租了一台电脑）。

## PhotoPrism 适合什么人？

PhotoPrism 适合极少的一部分人，是的，极少。但它的确没有隐私担忧。

官网地址：[https://photoprism.app/](https://photoprism.app/?ref=appinn)