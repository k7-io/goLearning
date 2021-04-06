## rankList
* 排行榜

eg:
* 周排行榜/总排行榜


## redis sorted set
* [sorted set(http://www.redis.cn/commands.html#sorted_set)
* [Redis 有序集合(sorted set)](https://www.runoob.com/redis/redis-sorted-sets.html)

```shell
127.0.0.1:6379> ZADD runoobkey 1 redis
(integer) 1
127.0.0.1:6379> ZADD runoobkey 2 mongodb
(integer) 1
127.0.0.1:6379> ZADD runoobkey 3 mysql
127.0.0.1:6379> ZADD runoobkey 2.2 postgre
(integer) 1
127.0.0.1:6379> ZRANGE runoobkey 0 10
1) "redis"
2) "mongodb"
3) "postgre"
4) "mysql"
127.0.0.1:6379>
127.0.0.1:6379> ZRANGE runoobkey 0 10 WITHSCORES
1) "redis"
2) "1"
3) "mongodb"
4) "2"
5) "postgre"
6) "2.2000000000000002"
7) "mysql"
8) "3"
127.0.0.1:6379>
```