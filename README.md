# parseRedisString
解析Redis中string字符串

使用go实现
Redis中字符串类型如下：（如果字符串为空，直接输出空格）
    第一个字节是“+”代表简单字符串类型。     "+OK\r\n"
    第一个字节是“-”代表错误类型。           "-Error message\r\n"
    第一个字节是“:”代表整型。               ":1000\r\n"*   
    第一个字节是“$”代表块字符串。           "$6\r\nfoobar\r\n"
    第一个自己是“*”代表数组。               "*2\r\n$3\r\nfoo\r\n$3\r\nbar\r\n"
