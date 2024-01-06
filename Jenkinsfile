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

    }
}