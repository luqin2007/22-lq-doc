---
title: Elasticsearch：官方分布式搜索和分析引擎 | Elastic
source: https://www.elastic.co/cn/elasticsearch
author:
  - "[[Elastic]]"
published: 
created: 2024-11-23
description: 在 RESTful 风格的分布式开源搜索和分析引擎中，Elasticsearch 处于领先地位，速度快，可实现水平可扩展性和可靠性，并能让您轻松进行管理。免费启用。...
tags:
  - clippings
---
![](https://www.elastic.co/static-res/images/hero/generic-a-light-left.svg)

![](https://www.elastic.co/static-res/images/hero/generic-d-light-right.svg)

![](https://static-www.elastic.co/v3/assets/bltefdd0b53724fa2ce/blt36f2da8d650732a0/5d0823c3d8ff351753cbc99f/logo-elasticsearch-32-color.svg)

## Elastic Stack 的核心

Elasticsearch 是一个开源的分布式 RESTful 搜索和分析引擎、可扩展的数据存储和向量数据库，能够解决不断涌现出的各种用例。作为 Elastic Stack 的核心，Elasticsearch 会集中存储您的数据，让您飞快完成搜索，微调相关性，进行强大的分析，并轻松缩放规模。

[下载 Elasticsearch](https://www.elastic.co/cn/downloads/elasticsearch)

刚接触 Elasticsearch？立即部署并运行。

[观看视频](https://www.elastic.co/cn/virtual-events/getting-started-elasticsearch)

参加我们的 Elasticsearch 工程师培训，为使用 Elasticsearch 奠定良好的基础。

[查看培训](https://www.elastic.co/cn/training/elasticsearch-engineer-1)

培养高级 Elasicsearch 技能，例如相关度调整、文本分析等等。

[查看培训](https://www.elastic.co/cn/training/)

## 深入了解全新 Elasticsearch Relevance Engine

[Elasticsearch Relevance Engine](https://www.elastic.co/cn/elasticsearch/elasticsearch-relevance-engine)™ (ESRE) 旨在为基于人工智能的搜索应用程序提供强大支持。使用 ESRE，您可以应用具有卓越相关性的开箱即用型语义搜索（无需域适应），与外部大型语言模型 (LLM) 集成，实现混合搜索，并使用第三方或您自己的转换器模型。

![](https://static-www.elastic.co/v3/assets/bltefdd0b53724fa2ce/bltc00f65307e5db868/646ba194a21aba2dc3f0d739/illustration-hero-esre.png)

## 查询和分析

## 从数据中探寻各种问题的答案

- ![](https://images.contentstack.io/v3/assets/bltefdd0b53724fa2ce/blt5cfd03e50d91fe0f/5e6157c725d22d7db56a574d/icon-search-ui-32-color.svg)

### 定义您自己的搜索方式

通过 Elasticsearch，您能够执行和合并多种类型的搜索（结构化数据、非结构化数据、地理位置、指标，以及从管道查询语言搜索）。先从一个简单的问题出发，试试看能够从中发现些什么。
- ![](https://images.contentstack.io/v3/assets/bltefdd0b53724fa2ce/blta28bdc6896c47075/5d0ca0cd77f34fd55839aeae/icon-scale-32-color.svg)

### 大规模存储和分析

找到与查询最匹配的 10 个文档并不困难。但是如果面对的是十亿行日志，又该如何解读呢？Elasticsearch 聚合让您能够从大处着眼，探索数据的趋势和模式。

## 速度

## Elasticsearch 很快。快到不可思议。

- ### 快速获得结果

如果您能够立即获得答案，您与数据的关系就会发生变化。这样您就有条件进行迭代并涵盖更大的范围。
- ### 强大的设计

但是要达到这样的速度并非易事。我们通过有限状态机实现了用于全文检索的倒排索引，实现了用于存储数值数据和位置数据的 BKD 树， 以及用于分析的列存储。
- ### 无所不包

而且由于每个数据都被编入了索引，因此您再也不用因为某些数据没有索引而烦心。您可以用快到令人发指的速度使用和访问您的所有数据。

## 可扩展性

## 可以在笔记本电脑上运行。也可以在承载了 PB 级数据的成百上千台服务器上运行。

原型环境和生产环境可无缝切换；无论 Elasticsearch 是在一个节点上运行，还是在一个包含 300 个节点的集群上运行，您都能够以相同的方式与 Elasticsearch 进行通信。

它能够水平扩展，每秒钟可处理海量事件，同时能够自动管理索引和查询在集群中的分布方式，以实现极其流畅的操作。

![Illustration](https://static-www.elastic.co/v3/assets/bltefdd0b53724fa2ce/blt8c328002d82e303e/5d0d573477f34fd55839b61f/illustration-elasticsearch-scalability-555.png)

## 弹性

## 我们在您高飞的时候保驾护航。

硬件故障。网络分割。Elasticsearch 为您检测这些故障并确保您的集群（和数据）的安全性和可用性。通过跨集群复制功能，辅助集群可以作为热备份随时投入使用。Elasticsearch 运行在一个分布式的环境中，从设计之初就考虑到了这一点，能够让您永久高枕无忧。

![](https://static-www.elastic.co/v3/assets/bltefdd0b53724fa2ce/blt61799e12d10f4581/5e6158f8dc0f1706df255d1c/illustration-elasticsearch-resiliency-555.png)

## 灵活性

## 存储和探索数据以满足自身需求。

数据是不断变化的，这使得存储和搜索全部数据变得非常昂贵。Elasticsearch 能让您在性能和成本之间取得平衡。您可以将数据存储在本地以实现快速查询，也可以将无限量的数据[远程存储于低成本的 S3](https://www.elastic.co/cn/elasticsearch/elasticsearch-searchable-snapshots) 上。借助[运行时字段](https://www.elastic.co/cn/elasticsearch/elasticsearch-runtime-fields)，您还可以快速加载数据并针对变化做出相应调整。

![Illustration](https://static-www.elastic.co/v3/assets/bltefdd0b53724fa2ce/blt50671ed5bbad50f3/608fbbdc2d1d221032193ff1/illustration-balance-cost.png)

### Elasticsearch - 部署最广泛的向量数据库

#### 在两分钟内复制到本地进行试用

```
curl -fsSL https://elastic.co/start-local | sh
```

[阅读文档](https://www.elastic.co/guide/en/elasticsearch/reference/current/elasticsearch-intro.html)

## 用例

## 我到底能够使用 Elasticsearch 做什么？

数字、文本、地理位置、结构化、非结构化。欢迎使用所有数据类型。全文本搜索只是全球众多公司利用 Elasticsearch 解决各种挑战的冰山一角。查看直接依托 Elastic Stack 所构建解决方案的完整列表。

- ![](https://images.contentstack.io/v3/assets/bltefdd0b53724fa2ce/blt6c24c7dfdefbe990/5d046a2c0b20e952523c9642/logo-logging-64-color.svg)

### 日志监测
- ![](https://images.contentstack.io/v3/assets/bltefdd0b53724fa2ce/blt27971019eb9ad70f/5d046a516a73714d525f13a9/logo-metrics-64-color.svg)

### 基础架构监测
- ![](https://images.contentstack.io/v3/assets/bltefdd0b53724fa2ce/blt8f9a891b1832bac4/5d07fc7797f2babb5af90736/logo-apm-64-color.svg)

### APM
- ![](https://images.contentstack.io/v3/assets/bltefdd0b53724fa2ce/bltbe80b46a5391fbd7/5d082a3e616162aa5a85706d/logo-uptime-32-color.svg)

### 合成监测
- ![](https://images.contentstack.io/v3/assets/bltefdd0b53724fa2ce/blt549f7d977c2a88f4/5d082d34616162aa5a85707d/logo-enterprise-search-32-color.svg)

### Search
- ![](https://images.contentstack.io/v3/assets/bltefdd0b53724fa2ce/blt9fcd982a6d4e49f8/5d097c21970556dd5800d0af/logo-maps-64-color.svg)

### 地图
- ![](https://images.contentstack.io/v3/assets/bltefdd0b53724fa2ce/blt8bb3bd866432b54f/5da5be770e045e6904a28bc9/logo-siem-64-color.svg)

### SIEM
- ![](https://images.contentstack.io/v3/assets/bltefdd0b53724fa2ce/blt93d9e66fdbb25c4a/5d9e8be911db0711236249bd/logo-endpoint-64-color.svg)

### 终端安全

## 下列客户已在使用，并深受他们的信任和喜爱

## 真正的 Elasticsearch 体验

## 这与 Amazon 的 Elasticsearch Service 一样吗？

- ### 不。仅此一家。
- ### 专属功能

使用多项专属功能，例如 Machine Learning、针对 BI 连接性的 ODBC 驱动程序、自动管理时序型数据，以及告警。

## 增强

## Elasticsearch 功能

- ![](https://images.contentstack.io/v3/assets/bltefdd0b53724fa2ce/bltfe4b512a22783724/5d082c605fe8e2af5a559a9e/icon-security-settings-32-color.svg)

### 安全性

在数据粒度层面为 Elasticsearch 提供强大保护。
- ![](https://images.contentstack.io/v3/assets/bltefdd0b53724fa2ce/blt78bc0a900d2c485f/5d082c3297f2babb5af907fc/icon-monitoring-32-color.svg)

### Monitoring

随时了解 Elastic Stack 的动态，确保其处于最佳状态。
- ![](https://images.contentstack.io/v3/assets/bltefdd0b53724fa2ce/blt73db4351e80113f1/5d085611616162aa5a8570e7/icon-alerting-32-color.svg)

### 告警
- ![](https://images.contentstack.io/v3/assets/bltefdd0b53724fa2ce/bltf04dd93e9da6275a/5d082c3ed8ff351753cbc9dd/icon-sql-32-color.svg)

### Elasticsearch SQL

使用 SQL 与您的数据进行交互，还可使用 ODBC 和 JDBC 驱动程序访问这些数据。
- ![](https://images.contentstack.io/v3/assets/bltefdd0b53724fa2ce/blt03a57f1b95872b92/5e615b2cb28469061c944332/icon-real-time-32-color.svg)

### 时序数据管理

通过索引生命周期管理、冻结索引和汇总等方式，自动完成进程。
- ![](https://images.contentstack.io/v3/assets/bltefdd0b53724fa2ce/blt3638e20310e2d6b3/5d082be797f2babb5af907f6/icon-machine-learning-32-color.svg)

### 机器学习

自动对 Elasticsearch 数据进行异常检测。

## 客户端库

## 使用您自己的编程语言与 Elasticsearch 进行交互

Elasticsearch 使用的是标准的 RESTful API 和 JSON。此外，我们还构建和维护了很多其他语言的客户端，例如 [Java、Python、.NET、SQL 和 PHP](https://www.elastic.co/cn/guide/en/elasticsearch/client/index.html)。与此同时，我们的[社区也贡献了很多客户端](https://www.elastic.co/cn/guide/en/elasticsearch/client/community/current/index.html)。这些客户端使用起来简单自然，而且就像 Elasticsearch 一样，不会对您的使用方式进行限制。

```
curl -H "Content-Type: application/json" -XGET
'http://localhost:9200/social-*/_search' -d '{
  "query": {
    "match": {
      "message": "myProduct"
    }
  },
  "aggregations": {
    "top_10_states": {
      "terms": {
        "field": "state",
        "size": 10
      }
    }
  }
}'
```

## 分布式

## 您的部署您做主

无论您在哪里进行搜索，我们都能满足您的需求。

- ## Elastic Cloud

### 在 AWS、Google Cloud 和 Azure 上部署托管型 Elasticsearch 和 Kibana

在您所选的云服务提供商平台上迅速完成全套技术栈的部署。作为 Elasticsearch 的开发公司，我们为您在云端的 Elastic 集群提供多项功能和贴心支持。
- ## 本地部署

### 下载 Elasticsearch

安装最新版本后，只需几步即可开始在您的设备上运行 Elasticsearch。