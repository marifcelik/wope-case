#!/bin/env bash

if [ "$1" == "install" ]; then
    if ! command -v minikube &> /dev/null; then
            echo "Minikube yüklü değil"
            echo -n "Minikube'i yüklemek ister misiniz? (y/n)"
            read answer
            if [ "$answer" == "y" ]; then
                echo "Minikube yükleniyor..."
                curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
                sudo install minikube-linux-amd64 /usr/local/bin/minikube
            fi
    else
        echo "Minikube yüklü"
    fi
    minikube start --network-plugin=cni --cni=calico
    
    if ! command -v kubectl &> /dev/null; then
        echo "Kubectl yüklü değil"
        echo -n "Kubectl'i yüklemek ister misiniz? (y/n)"
        read answer
        if [ "$answer" == "y" ]; then
            echo "Kubectl yükleniyor..."
            kubeInstallCommand=""
            if [ "$(uname -m)" == "x86_64" ]; then
                kubeInstallCommand="curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl""
            elif [ "$(uname -m)" == "aarch64" ]; then
                kubeInstallCommand="curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/arm64/kubectl""
            else
                echo "Sistem desteklenmiyor."
                exit 1
            fi
            $kubeInstallCommand
            sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
        fi
    else
        echo "Kubectl yüklü"
    fi
    exit
fi

case $1 in
    docker)
        echo "Docker Compose ile uygulama çalıştırılıyor..."
        docker-compose up
        ;;
    helm)
        if ! command -v helm &> /dev/null
        then
            echo "Helm yüklü değil"
            echo -n "Helm'i yüklemek ister misiniz? (y/n)"
            read answer
            if [ "$answer" == "y" ]; then
                echo "Helm yükleniyor..."
                curl https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 | bash
            else
                echo "Helm yüklenmediği için uygulama deploy edilemedi."
                exit 1
            fi
        fi
        cd helm
        helm install --generate-name .
        ;;
    kubernetes|k8s)
        echo "Kubernetes ile uygulama çalıştırılıyor..."
        kubectl apply -f kubernetes
        ;;
    *)
        echo "Geçersiz komut"
        ;;
esac
