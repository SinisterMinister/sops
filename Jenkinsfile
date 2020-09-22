pipeline {
    agent {
        kubernetes {
            // Loads additional containers for jenkins to use
            yamlFile "jenkins-containers.yaml"
        }
    }
    options {
        office365ConnectorWebhooks([[
                startNotification: true,
                notifySuccess: true,
                notifyFailure: true,
                notifyAborted: true,
                url: 'https://outlook.office.com/webhook/fda64394-f023-4b50-9172-d81e9852802f@ef19afba-eaf9-4bb3-9cdb-7fb7a1dd99b8/IncomingWebhook/d0675771e1f8478d89972f68bfd7cfba/a1122964-04e3-4300-a655-b0d77ffdda32'
            ]]
        )
    }
    stages {
        stage('Build Image') {
            steps {
                container('builder') {
                    withCredentials([usernamePassword(credentialsId: "nexus-builder", usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
                        sh "docker login -u $USERNAME -p $PASSWORD hub.flatirons.cloud"
                    }

                    sh "docker build -t hub.flatirons.cloud/docker/sops:latest ."
                    sh "docker push hub.flatirons.cloud/docker/sops:latest"
                }
            }
        }
    }
}