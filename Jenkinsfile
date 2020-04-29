pipeline {
  agent any
  options {
    buildDiscarder(logRotator(numToKeepStr:'5', artifactNumToKeepStr:'5', artifactDaysToKeepStr:'7'))
    durabilityHint('PERFORMANCE_OPTIMIZED')
    retry(1)
    skipDefaultCheckout()
    timestamps()
  }
  stages {
    stage ('Checkout') {
      environment {
        def TAG = sh returnStdout: true, script: "git tag -l | tee tags | tail -n1 tags"
      }
      steps {
        checkout([$class:'GitSCM', branches:[[name:'*/*']], doGenerateSubmoduleConfigurations:false, extensions:[], submoduleCfg:[], userRemoteConfigs:[[ url:'https://github.com/samfil-technohub/samuel-nwoye-website.git']]])
        stash(name: 'ws', includes: '**', excludes: '**/.git/**')
        echo "Using Git Tag: ${env.TAG}"    
      }
    }
    stage ('Test') {
      agent { 
        docker { 
          image 'golang' 
          args ' -e GOCACHE=/tmp/.cache -e GO111MODULE=on -e CGO_ENABLED=0 -e GOOS=linux -e GOARCH=amd64 '
        } 
      }
      steps {
        unstash 'ws'
        sh 'go version'
        sh 'go test -v'
      }
    }
    stage ('Build') {
      agent { 
        docker { 
          image 'golang'
          args ' -e GOCACHE=/tmp/.cache -e GO111MODULE=on -e CGO_ENABLED=0 -e GOOS=linux -e GOARCH=amd64 '
        } 
      }
      steps {
        unstash 'ws'
        sh 'go version'
        sh 'go build -o pipeline main.go'
      }
    }
    stage('Deploy') {
      when {
        branch 'master' 
      }
      steps {
        echo "Deploying to Production"
      }
    }
  }
  post {
    success {
      slackSend (channel: '#mymonitor', color: 'good', message: "*${currentBuild.currentResult}:* _Job_ ${env.JOB_NAME}; _Build_ ${env.BUILD_NUMBER}\n *Visit:* ${env.BUILD_URL}")
    }
    failure {
      slackSend (channel: '#mymonitor', color: 'danger', message: "*${currentBuild.currentResult}:* _Job_ ${env.JOB_NAME}; _Build_ ${env.BUILD_NUMBER}\n *Visit:* ${env.BUILD_URL}")
    }
  }
}