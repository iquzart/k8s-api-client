package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var kubeconfig *string

	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// uses the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, _ := kubernetes.NewForConfig(config)
	check_err(err)

	// access the API to list pods
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), meta_v1.ListOptions{})
	check_err(err)
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	// List Pods
	podList, err := clientset.CoreV1().Pods("").List(context.TODO(), meta_v1.ListOptions{})
	check_err(err)

	fmt.Printf("List Of Pods:\n")
	for _, p := range podList.Items {
		fmt.Printf("%s   %s   %s  \n", p.Name, p.Namespace, p.Status.Phase)
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func check_err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
