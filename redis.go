package redisgo

type RedisClient struct {
}

//getConn从连接池获取一个新的连接
func (this *RedisClient) getConn() {

}

//Do 执行redis方法
func (this *RedisClient) Do() {

}

//Do 执行redis方法
func (this *RedisClient) Auth(pass string) {

}

func (this *RedisClient) Get(key string) {

}

func (this *RedisClient) Set(key string) {

}

func (this *RedisClient) Exists() {

}

func (this *RedisClient) Del() {

}

func (this *RedisClient) Ttl() {

}

func (this *RedisClient) Expire() {

}

func (this *RedisClient) ExpireAt() {

}

func (this *RedisClient) Incr() (int64, error) {
	return 0, nil
}

func (this *RedisClient) Decr() (int64, error) {
	return 0, nil
}
