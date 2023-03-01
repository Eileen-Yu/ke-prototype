package tetragon

import (
	"bufio"
	"context"
	"fmt"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var NAMESPACE = "kube-system"

var POD = "tetragon-2rs6l"

var TAIL_LINES int64 = 10

func GetTetragonLogs() {
	kubeConfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		fmt.Println("Failed to load kubeconfig", err)
		return
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("Failed to initialize a k8s client", err)
		return
	}

	/*
		pods, err := clientset.CoreV1().Pods(NAMESPACE).List(context.Background(), v1.ListOptions{})
		if err != nil {
			fmt.Errorf("Failed to list pods of namespace %s", NAMESPACE)
		}

		for _, p := range pods.Items {
			fmt.Println(p)
		}
	*/

	podLogOptions := corev1.PodLogOptions{
		Follow:    true,
		TailLines: &TAIL_LINES,
		Container: "export-stdout",
	}

	podLogRequest := clientset.CoreV1().
		Pods(NAMESPACE).
		GetLogs(POD, &podLogOptions)

	logStream, err := podLogRequest.Stream(context.Background())
	if err != nil {
		fmt.Println("Failed to open pod log stream", err)
		return
	}

	if logStream == nil {
		fmt.Println("Unexpected nil logstream")
		return
	}

	fmt.Println("1")
	rd := bufio.NewScanner(logStream)
	fmt.Println("2")

	var line string
	for {
		for rd.Scan() {
			line = rd.Text()
			fmt.Printf("😄😄😄\n%s\n😉😉😉\n", line)
		}
	}
}
