package configs

import (
    "encoding/json"
    "io/ioutil"
)

// 程序配置
type Config struct {
    HttpAddress string `json:"http_address"`
    GrpcAddress string `json:"grpc_address"`
}

type Conf struct {
    *Config
    filepath string
}

// 加载配置
func (c *Conf) Load() (err error) {
    var (
        content []byte
        conf    Config
    )

    // 1. 把配置文件读进来
    if content, err = ioutil.ReadFile(c.filepath); err != nil {
        return
    }

    // 2. json做反序列化
    if err = json.Unmarshal(content, &conf); err != nil {
        return
    }

    // 3. 复制单例
    c.Config = &conf
    return
}

func NewConf(filepath string) *Conf {
    return &Conf{filepath: filepath}
}
