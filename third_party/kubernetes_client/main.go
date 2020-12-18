package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	//net.InterfaceByIndex()

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	//service -> lables
	labels := "system=centos"
	result, err := clientset.CoreV1().Pods("default").Watch(context.TODO(), metav1.ListOptions{LabelSelector: labels})
	if err != nil {
		panic(err.Error())
	}
	resultCh := result.ResultChan()
	for ch := range resultCh {
		fmt.Printf("type:%s\n", string(ch.Type))
	}
	/*pods, err := clientset.CoreV1().Pods("default").List(context.TOD	O(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	for _, pod := range pods.Items {
		//fmt.Printf("Name: %s, Status: %s Namespace:%s NodeName:%s\n",
		//	pod.ObjectMeta.Name, pod.Status.Phase, pod.Namespace, pod.Spec.NodeName)

		//if pod.ObjectMeta.Name == "my-centos" {
		containers := pod.Spec.Containers
		fmt.Printf("pod UUID:%s NodeName:%s\n", pod.ObjectMeta.UID, pod.Spec.NodeName)
		for _, c := range containers {
			fmt.Printf("Name:%s", c.Name)
		}
		//}
	}*/
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
