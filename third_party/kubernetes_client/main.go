package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/remotecommand"
	"os"
	"path/filepath"
	"strings"
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

	//Watch监视功能的使用
	//service -> lables
	/*labels := "system=centos"
	result, err := clientset.CoreV1().Pods("default").Watch(context.TODO(), metav1.ListOptions{LabelSelector: labels})
	if err != nil {
		panic(err.Error())
	}
	resultCh := result.ResultChan()
	for ch := range resultCh {
		fmt.Printf("type:%s\n", string(ch.Type))
	}*/

	//在机器上执行kubectl命令
	/*podName := "my-centos"
	command := "cat /sys/class/net/eth0/iflink"
	err = ExecCmdExample(clientset, config, podName, command)
	if err != nil {
		fmt.Printf("ExecCmdExample exec failed:%s\n", err.Error())
	}*/

	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	for _, pod := range pods.Items {
		//fmt.Printf("Name: %s, Status: %s Namespace:%s NodeName:%s\n",
		//	pod.ObjectMeta.Name, pod.Status.Phase, pod.Namespace, pod.Spec.NodeName)

		if pod.ObjectMeta.Name == "my-centos" {
			containers := pod.Spec.Containers
			fmt.Printf("pod UUID:%s NodeName:%s\n", pod.ObjectMeta.UID, pod.Spec.NodeName)
			for _, c := range containers {
				fmt.Printf("container name:%s", c.Name)
			}
		}
	}
}

func ExecCmdExample(client kubernetes.Interface, config *restclient.Config, podName string,
	command string /*, stdin io.Reader, stdout io.Writer, stderr io.Writer*/) error {
	/*cmd := []string{
		"sh",
		"-c",
		command,
	}*/
	var stdin bytes.Buffer
	req := client.CoreV1().RESTClient().Post().Resource("pods").Name(podName).
		Namespace("default").SubResource("exec")
	option := &v1.PodExecOptions{
		Command: strings.Fields(command),
		Stdin:   true,
		Stdout:  true,
		Stderr:  true,
		TTY:     true,
	}
	/*if stdin == nil {
		option.Stdin = false
	}*/
	req.VersionedParams(
		option,
		scheme.ParameterCodec,
	)
	exec, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
	if err != nil {
		return err
	}

	//stdin,stdout,stderr
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:  &stdin,
		Stdout: &stdout,
		Stderr: &stderr,
	})
	if err != nil {
		return err
	}
	fmt.Printf("output:%s err:%s\n", stdout.String(), stderr.String())

	return nil
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
