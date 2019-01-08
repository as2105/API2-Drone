pipeline {
  agent {
    node {
      label 'docker-kitchensink-slave'
      customWorkspace '/home/jenkins/workspace/go/src/github.com/SynapticHealthAlliance/fhir-api'
    }
  }

  environment {
    GOLANG_VERSION = '1.11'
    GOPATH = "/home/jenkins/workspace/go"
  }

  stages {
    stage('Test') {
      steps {
        withCredentials([usernamePassword(credentialsId: 'a12d225a-392a-4b01-872a-6038d1487628', usernameVariable: 'GITHUB_USERNAME', passwordVariable: 'GITHUB_TOKEN')]) {
          sh 'echo "machine github.com\n  login $GITHUB_USERNAME\n  password $GITHUB_TOKEN\n" > /home/jenkins/.netrc'
        }
        sh 'dep ensure -v'
        sh 'make sonar'
      }
    }

    stage('Report') {
      steps {
        sh 'sonar-scanner'
      }
    }
  }
}