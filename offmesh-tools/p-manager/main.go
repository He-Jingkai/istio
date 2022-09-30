package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	tools "p_manager/p-manager-tools"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"net/http"
)

var clientSet *kubernetes.Clientset
var PodProxyMap map[tools.PodMeta]*tools.PodMeta
var ProxyPool []*tools.PodMeta

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	Init()
	e.GET("/distribute_proxy/:namespace/:podName", DistributeProxy)
	e.GET("/return_proxy/:namespace/:podName", ReturnProxy)

	e.Logger.Fatal(e.Start(":80"))
}

func Init() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}

func DistributeProxy(c echo.Context) error {
	//get a new proxy from proxy pool
	podMeta := tools.PodMeta{
		NameSpace: c.Param("namespace"),
		Name:      c.Param("podName"),
	}
	proxy, err := PopTopProxyFromPool(&podMeta)
	if err != nil {
		return err
	}
	PodProxyMap[podMeta] = proxy
	str, err := json.Marshal(*proxy)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, string(str))
}

func ReturnProxy(c echo.Context) error {
	pod := tools.PodMeta{
		NameSpace: c.Param("namespace"),
		Name:      c.Param("podName"),
	}
	proxy := PodProxyMap[pod]
	delete(PodProxyMap, pod)
	ReturnProxyToPool(proxy)
	return c.String(http.StatusOK, "")
}

//TODO: ProxyPool管理策略

func AddNewProxyToPool(pod *tools.PodMeta) error {
	proxy, err := tools.CreateNewProxy(pod, clientSet)
	if err != nil {
		log.Error("CreateNewProxy Error: ", err)
	} else {
		ProxyPool = append(ProxyPool, proxy)
	}
	return err
}

func PopTopProxyFromPool(pod *tools.PodMeta) (*tools.PodMeta, error) {
	var proxy *tools.PodMeta
	if len(ProxyPool) == 0 {
		err := AddNewProxyToPool(pod)
		if err != nil {
			return nil, err
		}
	}
	proxy = ProxyPool[0]
	ProxyPool = ProxyPool[1:]
	return proxy, nil
}

func ReturnProxyToPool(proxy *tools.PodMeta) {
	err := tools.DeleteProxy(clientSet, proxy.Name)
	if err != nil {
		return
	}
	//ProxyPool = append(ProxyPool, proxy)
}
