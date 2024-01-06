pipeline {
    agent any
    tools {
        jdk 'java17'
        maven 'maven3'
    }
    stages {
        stage("Cleanup Workspace"){
            steps {
                cleanWs()
            }
        }
        stage("Checkout from SCM"){
            steps {
                git branch: 'main', credentialsId: 'tejaswankalluri-gh', url: 'https://github.com/tejaswankalluri/BuyLeaf-Backend'
            }
        }
        stage('Inject Secrets') {
            steps {
                withCredentials([file(credentialsId: 'spring-boot-env-file', variable: 'ENV_FILE')]) {
                    sh 'cp $ENV_FILE src/main/resources/secret.properties'
                }
            }
        }
        stage("Build Application"){
            steps {
                sh "mvn clean package"
            }
        }
        stage("Test Aplication"){
            steps {
                sh "mvn test"
            }
        }
        stage("Build Dockerfile"){
            steps {
                sh "docker build -t techtoe/buyleaf-backend ."
            }
        }
        stage('Push image to Hub'){
            steps {
                script {
                   withCredentials([string(credentialsId: 'techtoe-docker-pass', variable: 'dockerhubpwd')]) {
                   sh 'docker login -u techtoe -p ${dockerhubpwd}'
                }
                   sh 'docker push techtoe/buyleaf-backend'
                }
            }
        }
    }
}