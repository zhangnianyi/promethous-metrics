package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	log "promethous-metrics/log"
	"strings"
)
//如果你想为你的metrics设置其他的label  就要设置label 和value 必须是切片variableLabels []string
var vLable = []string{}
var vValue = []string{}
var constLabel = prometheus.Labels{"zhaoshangbank": "访问量"}  //设置一个常量用来上做metrics的label
var constLabel2 = prometheus.Labels{"promtthous": "qps"}
// define a struct   from prometheus's struct  named Desc
//下面的是固定写法
type ovsCollector struct {
	ovsMetric *prometheus.Desc
	barMetric *prometheus.Desc

}
//初始化一个 ovsCollector {
func newOvsCollector() *ovsCollector {
	var rm = make(map[string]string)
	rm = getBondStatus()

		for k, _ := range rm {
			// get the net
			vLable = append(vLable, k)
		}
	return &ovsCollector{
		//variableLabels []string  也必须传递一个 []string 进去
		ovsMetric: prometheus.NewDesc("ovs_bond_zhang",
			"Show ovs bond status-zhangchao123",vLable ,
			constLabel),
			//如果你不是设置多余的label 可以设置为nil
		barMetric: prometheus.NewDesc("bbb_metric",
			"test promethous for the cacaca",
			nil, constLabel2),
	}
}

//默认方法
func (collector *ovsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.ovsMetric
	ch <- collector.barMetric
}



//  get the value of the metric from a function who would execute a command and return a float64 value
//收集方法
func (collector *ovsCollector) Collect(ch chan<- prometheus.Metric) {
	var metricValue float64
	var test1 float64
	var rm = make(map[string]string)
	rm = getBondStatus()
	//map[a1b1:enabled a2b2:disabled]
		vValue = vValue[0:0]
		fmt.Println(vLable)
		for _, v := range vLable {
			// get the net
			vValue = append(vValue, rm[v])
			//vvale0444444 [enabled disabled]
			if v == "disabled" {
				metricValue++
			}
		}
		//如果你是有label的话这里就需要进行匹配 匹配方法是内部方法，newOvsCollector 这个函数传递了label 进去，Collect 传递了value进去
		//a[a1b1,a2b2] b[enabled,disabled]
	    //a[0]=>b[0]
	    //a[1]=>b[1]
		metricValue =1000
		test1 =10

		log.SugarLogger.Info("collector.ovsMetric",collector.ovsMetric,"prometheus.CounterValue",prometheus.CounterValue,"metricValue",metricValue)
		//直接塞进去就就可以了
		ch <- prometheus.MustNewConstMetric(collector.ovsMetric, prometheus.CounterValue, metricValue,vValue...)
	    ch <- prometheus.MustNewConstMetric(collector.barMetric, prometheus.CounterValue,test1 )
	}

// define metric's name、help


func getBondStatus() (m map[string]string) {
	tt := []string{"zhang:enabled","zeng:disabled"}
	var nMap = make(map[string]string)
	for i := 0; i < len(tt); i++ {
		// if key contains "-"åå
		 //tt := []string{"a1-b1:enabled","a2-b2:disabled"}
		if strings.Contains(tt[i], "-") == true {

			nKey := strings.Split(strings.Split(tt[i], ":")[0], "-")

			// [a1 b1]
			nMap[strings.Join(nKey, "")] = (strings.Split(tt[i], ":"))[1]

		} else {
			nMap[(strings.Split(tt[i], ":"))[0]] = (strings.Split(tt[i], ":"))[1]
		}
	}

	return nMap
	//a1b1-111 map[a1b1:enabled a2b2:disabled]
}

func main() {
	log.InitLogger()
	defer log.SugarLogger.Sync()
	ovs := newOvsCollector()
	prometheus.MustRegister(ovs)

	http.Handle("/metrics", promhttp.Handler())
	log.SugarLogger.Info("begin to server on port 8080")
	//log.Info("begin to server on port 8080")
	// listen on port 8080
	log.SugarLogger.Info(http.ListenAndServe(":8080", nil))
}


