# SYLU-EOA
基于爬虫获取正方教务系统数据(理论上适用于所有基于正方教务的教务系统)，客户端为基于Uniapp/原生小程序开发的安卓app/微信小程序。
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
- [x] 创新创业学分查询
# 未来计划
优化数据存储，利用缓存提升响应时间。
# 运行项目
修改conf下的demo_config文件为你的配置信息，重命名为config即可，数据库表结构在mysql初始化时自动建表。
# 接口文档
接口文档使用swagger框架实现，运行项目后，访问```http://host:port/swagger/index.html```即可查看。
# 正方教务相关接口文档
