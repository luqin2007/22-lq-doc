---
title: "使用-Zeal-打造属于自己的文档"
source: "https://zhuanlan.zhihu.com/p/396577255"
author:
  - "[[知乎专栏]]"
published:
created: 2024-12-04
description: "引言作为一名程序开发者，翻查文档成为了每天都必须去做的事情。然而，由于使用到的知识和工具不相同，所以几乎每次都需要重新打开官网的文档进行搜索，而且不同知识和工具的文档还需要切换浏览器标签来浏览，十分…"
tags:
  - "clippings"
---
## Zeal 创建自己专属的文档项目

---

经过上面的步骤后，这个文档系统已经可以很好地为我们日常开发服务了，但是除了下载常用的 api 文档之外，能不能编写我们自己的文档呢？

Zeal 官方写了一个教我们如何去编写自己的 Docsets 的[文档](https://link.zhihu.com/?target=https%3A//kapeli.com/docsets)，但是该文档写得比较简单，并且没有详细地操作指引，操作起来比较复杂。

> 经过实验之后，Zeal 的 Docsets 其实是 html 的集合，那么我们可以先用文档工具，生成一些静态的 html 文档。然后通过 Docsets 官方提供的 Docsets 生成器来把 html 生成 Docsets，这样就可以生成出属于我们自己的 Docsets 了。  

### 合适的文档生成器

目前各种开发语言都有文档生成器，我比较熟悉的 Node.js 就有数十个像 Gitbook、Docsify、Vuepress 等等。但是并不是每一个都适合用来制作 Docsets，举个例子：

> Docsify 是一个很棒的生成器，但是用于 Docsets 的话就会有问题。原因是因为 Vuepress 是通过 js 读取 Markdown 来实现的，而 Zeal 内部是一个浏览器，并没有静态服务器，所以制作出来的 Docsets 会出现跨域的问题。  

最终我选择了使用 Gitbook 来制作 Docsets，它能生成静态的 Html 文件，并且能够通过本地双击打开，能够跟 Zeal 完美融合。

### 编写文档

确定了使用 Gitbook 之后，先安装 Gitbook： `npm install gitbook-cli -g`

然后，使用 Gitbook 创建项目： `gitbook init`

创建完之后，会看到如下目录结构：

```text
├── README.md
├── SUMMARY.md
├── chapter1
│   ├── README.md
│   ├── section1.1.md
│   └── section1.2.md
└── chapter2
    └── README.md
```

创建完成之后，通过命令 `gitbook serve` 即可看到文档的样式：

![](https://pic4.zhimg.com/v2-4971af97b21ce90f20b9f77aa6810623_1440w.jpg)

  

接下来就是每个开发者根据自己的需要，编写自己的 markdown 文件和修改 SUMMARY.md。由于文档是离线的，所以不用担心会被其他人盗取，可以把日常用到的账号密码、代码片段、常用的资料都分类写进去。

写完之后，通过命令 `gitbook build` 生成 html 文件。生成后会见到一个 `_book` 的目录，所有的 html 文件都在里面，这样文档就算编写完成了。

## 把 html 生成 Docsets

---

有了文档对应的 html 之后，需要把 html 生成 Docsets。我使用 Node.js 生成，在 npm 上面找了一个叫 [docset-generator](https://link.zhihu.com/?target=https%3A//www.npmjs.com/package/docset-generator) 的插件。通过以下代码，即可生成对应的 Docsets：

```text
let DocSetGenerator = require("docset-generator").DocSetGenerator;
let docSetGenerator = new DocSetGenerator({
  destination: "./output/",
  name: "Season",
  documentation: "./personal-doc/_book",
  icon: './icon.png',
  entries: [ // entries 可以设置 Docsets 的分类，一般一个分类对应一个 html
    {
      name: "个人资料",
      type: "Guide",
      path: "./index.html"
    },
    ...
  ]
});
docSetGenerator.create();
```

这里我把 Gitbook 和 docset-generator 整合到一个项目中，通过 [npm-run-all](https://link.zhihu.com/?target=https%3A//www.npmjs.com/package/npm-run-all) 插件，先调用 Gitbook 生成 html，然后再生成 Docsets：

```text
...
"scripts": {
  "start": "gitbook build ./personal-doc && node index.js"
},
...
```

生成完之后，将会得到一个 XXX.docsets 的文件夹，接下来就是最后的一步，把这个文件夹放进 Zeal 里面。

在 Zeal 里面有一个叫 docsets 的文件夹，进去之后会看到下载的 docsets 都在里面，把刚刚生成的 docsets 放进去，重启一下 Zeal 就可以看到自己的文档了。

![](https://pic4.zhimg.com/v2-6b7ff978f1a8de6ee8f6d855963513c3_1440w.jpg)