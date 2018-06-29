package readconfig

import (
	"os"
	"bufio"
	"strings"
	"io"
	"log"
	"flag"
)

var (
	enviroment string
	configDir string
	configMap = make(map[string]interface{})
)

func ReadConfigUtils() {
	flag.StringVar(&enviroment, "enviroment", "dev", "enviroment")
	flag.Parse()
	log.Printf("enviroment : %v \n",enviroment)
	//打开这个ini文件
	f, _ := os.Open(getConfigDir())
	//读取文件到buffer里边
	buf := bufio.NewReader(f)
	for {
		//按照换行读取每一行
		l, err := buf.ReadString('\n')
		//相当于PHP的trim
		line := strings.TrimSpace(l)
		//判断退出循环
		if err != nil {
			if err != io.EOF {
				//return err
				panic(err)
			}
			if len(line) == 0 {
				break
			}
		}
		switch {
		case len(line) == 0:
			//匹配[db]然后存储
		case line[0] == '[' && line[len(line)-1] == ']':
			section := strings.TrimSpace(line[1 : len(line)-1])
			log.Println(section)
		default:
			//dnusername = xiaowei 这种的可以匹配存储
			i := strings.IndexAny(line, "=")
			configMap[strings.TrimSpace(line[0:i])] = strings.TrimSpace(line[i+1:])
		}
	}
	//循环输出结果
	log.Println("--------------- config start ---------------")
	for k, v := range configMap {
		log.Println(k,":", v)
	}
	log.Println("--------------- config end ---------------")
}

func getConfigDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	//configDir = dir + "/" + getConfigName()
	//configDir = dir + "/project/message/" + getConfigName()
	if enviroment == "dev"{
		configDir = dir + "/project/message/" + getConfigName()
	}else{
		configDir = dir + "/" + getConfigName()
	}
	log.Printf("configDir : [%v]\n",configDir)
	return configDir
}

func getConfigName() string {
	return "config-"+enviroment+".ini"
}