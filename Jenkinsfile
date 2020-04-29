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
        echo "Using Git Tag: ${env.TAG}"
        stash(name: 'ws', includes: '**', excludes: '**/.git/**')    
      }
    }
    stage ('Build') {
      agent { 
        docker { 
          image 'golang'
        } 
      }
      steps {
        unstash 'ws'
        sh 'go build main.go'
      }
    }
    stage ('Test') {
      agent { 
        docker { 
          image 'golang' 
        } 
      }
      steps {
        unstash 'ws'
        sh 'go test -v'
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