# idgenerator

#### 介绍
用一种全新的雪花漂移算法，让ID更短、生成速度更快。
核心在于缩短ID长度的同时，还能保持每毫秒并发处理量（50W/0.1s），且能保持伸缩能力。

支持QQ群：646049993


#### 功能说明：

1.ID更短，是传统算法的几倍，用50年都不会超过js（Number）的最大值。（默认配置WorkerId是6bit，自增数是6bit）

2.生成速度更快，0.1秒可生成50万个。（i7笔记本，默认算法配置6bit+6bit）

3.支持时间回拨处理。

4.支持手工插入新ID。

5.漂移时对外异步发通知事件。

6.不依赖任何外部缓存和数据库。

7.目前是C#版，很快会出java、php等版本。


#### 文件说明：

1.SnowWorkerM1.cs 是雪花漂移算法。

2.SnowWorkerM2.cs 是传统雪花算法。

