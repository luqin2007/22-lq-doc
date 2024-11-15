LangChain 是一个基于大语言模型（LLMs）用于构建端到端语言模型应用的框架，让 AI 调用具有以下特性：

- 数据感知：将语言模型与其他数据联系起来，实现丰富、多样化的数据利用
- 代理性：实现语言模型与环境的交互，使模型能够深入理解环境并响应

> [!note] 个人感觉有点类似于 Spring AI
# 组件

`````col
````col-md
flexGrow=2
===
- 模型：Models，各大语言模型的 LangChain 接口和调用细节，以及输出解析机制。
- 提示模板：Prompts，提示工程流线化
- 数据检索：Indexes，构建并操作文档的方法，接受用户的查询并返回最相关的文档，搭建本地知识库
- 记忆：Memory，通过短时记忆和长时记忆，在对话过程中存储和检索数据
- 链：Chains，以特定方式封装各种功能，并通过一系列的组合，自动而灵活地完成常见用例
- 代理：Agents，让大模型自主调用内外部工具
````
````col-md
flexGrow=1
===
![[../../../_resources/images/76619cf2f73ef200dd57cd16c0d55ec4.png]]
````
`````
# 使用

1. 配置环境，安装 `langchain` 和相关其他配置
2. 创建 `llm` 对象，使用相关工具

---

- [[内部知识库]]
- 核心组件
	- [[核心组件/模型/模型|模型]]
	- [[核心组件/链|链]]
	- [[核心组件/记忆|记忆]]
	- [[核心组件/代理|代理]]
- [[工具和工具箱]]
- [[检索增强生成]]

---

- [[通过链和代理查询鲜花信息]]
- [[在AI应用中引入异步通信机制]]
- [[通过角色扮演脑暴一个鲜花营销方案]]
- [[根据气候变化自动制定鲜花存储策略]]
- [[部署一个鲜花网络电商的人脉工具]]
- [[易速鲜花聊天客服机器人的开发]]
# 参考

```cardlink
url: https://juejin.cn/book/7387702347436130304
title: "LangChain 实战课 - 掘金小册 - 掘金小册"
description: "手把手带你开发专属的 ChatGPT 应用。「LangChain 实战课」由掘金小册撰写，2114人购买"
host: juejin.cn
favicon: https://lf-web-assets.juejin.cn/obj/juejin-web/xitu_juejin_web/static/favicons/favicon-32x32.png
```

```cardlink
url: https://python.langchain.ac.cn/docs/introduction/
title: "简介 | 🦜️🔗 LangChain 中文"
host: python.langchain.ac.cn
favicon: https://python.langchain.ac.cn/img/brand/favicon.png
image: https://python.langchain.ac.cn/img/brand/theme-image.png
```

接入不同 AI 模型可查询：

```cardlink
url: https://python.langchain.ac.cn/v0.2/docs/integrations/llms/
title: "大型语言模型 | 🦜️🔗 LangChain 中文"
host: python.langchain.ac.cn
favicon: https://python.langchain.ac.cn/v0.2/img/brand/favicon.png
image: https://python.langchain.ac.cn/v0.2/img/brand/theme-image.png
```
