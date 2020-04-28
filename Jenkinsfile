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
      steps {
        checkout([$class:'GitSCM', branches:[[name:'*/*']], doGenerateSubmoduleConfigurations:false, extensions:[], submoduleCfg:[], userRemoteConfigs:[[ url:'https://github.com/knoxknot/samuel-nwoye.git']]])
        script {
          def TAG = sh returnStdout: true, script: "git tag -l | tee tags | tail -n1 tags"
        }
        echo "Using Git Tag: ${TAG}"
        stash(name: 'ws', includes: '**', excludes: '**/.git/**')    
      }
    }
    stage ('Build') {
      agent { 
        docker { 
          image 'golang'
          arg '-v \$(pwd):/app'
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