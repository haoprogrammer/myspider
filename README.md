#go语言爬虫项目

#技术
golang


#简介
1. fetcher:获取数据范围信息
2. parser:解析器
3. engine:爬虫引擎
4. schduler:调度器


#工作
1. 优化代码
2. 支持城市翻页


#时间
2019417 如何存数据，只要生成有价值的item即可
2019419 数据存储到es中;
        重构代码；
            1. 维护一个es client，不用每次save都有一个client
            2. 拆分出worker
            3. index名字可以从外面传入