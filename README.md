# SYLU-EOA
沈阳理工大学校园一体化服务平台--基于爬虫获取正方教务系统数据(理论上适用于所有基于正方教务的教务系统)，客户端为基于Uniapp/原生小程序开发的安卓app/微信小程序，预计2023年8月下旬上线。
# 框架
- 主体：[gin](https://github.com/gin-gonic/gin) [gorm](https://github.com/go-gorm/gorm)
- 爬虫部分: [colly](https://github.com/gocolly/colly) [resty](https://github.com/go-resty/resty)
# 存储(待定中)
基于腾讯云COS服务保存静态文件。
# 功能(开发中，按优先级排序)
- [x] 教务绑定
- [x] 成绩查询
- [x] 绩点查询
- [x] 课表查询
- [x] 校历查询
- [ ] 课程云标签
- [ ] 课程云标签
- [ ] 青年大截图
- [ ] 吃瓜园地
- [ ] 教务通知
- [ ] 绩点排行榜
- [ ] 课程讨论区
- [ ] 二手交易
- [ ] 失物招领
# 未来计划
优化数据存储，利用缓存提升响应时间。
# 运行项目
修改conf下的demo_config文件为你的配置信息，重命名为config即可，数据库表结构在mysql初始化时自动建表。
# 接口文档
接口文档使用swagger框架实现，运行项目后，访问```http://host:port/swagger/index.html```即可查看。
# 正方教务相关爬虫教程
- [正方教务系统数据爬取详解（一）登录cookie获取](https://www.hackerxiao.online/archives/schooldata)
- [正方教务系统数据爬取详解（二）数据获取](https://www.hackerxiao.online/archives/school2)