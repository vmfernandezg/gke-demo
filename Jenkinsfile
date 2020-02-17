pipeline {
    agent any
    environment {
        PROJECT_ID = 'endless-gamma-237509'
        LOCATION_TEST = 'us-west1-c'
        LOCATION_PROD = 'us-west1-b'
        CREDENTIALS_ID = 'testproject'
        CLUSTER_NAME_TEST = 'mycluster'
        CLUSTER_NAME_PROD = 'pruebas'  
    }
    stages {
        stage("Checkout code") {
            steps {
                checkout scm
            }
        }
        stage("Build image") {
            steps {
                script {
                    sh "whoami"
                    myapp = docker.build("vfernandezg/hello:${env.BUILD_ID}")
                }
            }
        }
        stage("Push image") {
            steps {
                script {
                    docker.withRegistry('https://registry.hub.docker.com', 'dockerhub') {
                            myapp.push("latest")
                            myapp.push("${env.BUILD_ID}")
                    }
                }
            }
        }        
        stage('Deploy to GKE test cluster') {
            steps{
                sh "sed -i 's/hello:latest/hello:${env.BUILD_ID}/g' deployment.yaml"
                step([$class: 'KubernetesEngineBuilder', projectId: env.PROJECT_ID, clusterName: env.CLUSTER_NAME_TEST, location: env.LOCATION_TEST, manifestPattern: 'deployment.yaml', credentialsId: env.CREDENTIALS_ID, verifyDeployments: false])
            }
        }
        stage('Deploy to GKE production cluster') {
            steps{
                input message:"Proceed with final deployment?"
                step([$class: 'KubernetesEngineBuilder', projectId: env.PROJECT_ID, clusterName: env.CLUSTER_NAME_PROD, location: env.LOCATION_PROD, manifestPattern: 'deployment.yaml', credentialsId: env.CREDENTIALS_ID, verifyDeployments: false])
            }
        }
    }    
}