# 外观模式

API 为 facade 模块的外观接口，大部分代码使用此接口简化对 facade 类的访问。

facade 模块同时暴露了a和b两个module 的NewXXX 和interface，其他代码如果需要使用细节功能时可以直接调用