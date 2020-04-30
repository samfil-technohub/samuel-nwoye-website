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
        checkout([$class:'GitSCM', branches: [[name: '*/master'], [name: '*/develop'], [name: '*/release']], 
        doGenerateSubmoduleConfigurations:false, extensions:[], submoduleCfg:[],
        userRemoteConfigs:[[ url:'https://github.com/samfil-technohub/samuel-nwoye-website.git']]])
        sh('''
            git config user.name 'knoxknot'
            git config user.email 'samuel.nwoye@yahoo.com' 
        ''') 
      }
    }
    stage ('Test') {
      agent { 
        docker { 
          image 'golang' 
          args ' -e GOCACHE=/tmp/.cache -e GO111MODULE=on -e GOOS=linux -e GOARCH=amd64 '
        } 
      }
      steps {
        sh 'go version'
        sh 'go mod download'
        sh 'go test -v'
      }
    }
    stage ('Build and Push') {
      agent { 
        docker { 
          image 'golang'
          args ' -e GOCACHE=/tmp/.cache -e GO111MODULE=on -e GOOS=linux -e GOARCH=amd64 '
        } 
      }
      steps {
        script {
          env.GIT_BRANCH = sh returnStdout: true, script: "git rev-parse --abbrev-ref HEAD"
        }
        sh("git checkout -B ${GIT_BRANCH}")  
        withCredentials([usernamePassword(credentialsId: 'github', passwordVariable: 'github_password', usernameVariable: 'github_username')]) {
          script {
            env.encodedPass=URLEncoder.encode(github_password, "UTF-8")
          }
          sh 'git pull https://${github_username}:${encodedPass}@github.com/samfil-technohub/samuel-nwoye-website.git'
          sh 'go mod download'
          sh 'go build main.go'
          sh 'git push https://${github_username}:${encodedPass}@github.com/samfil-technohub/samuel-nwoye-website.git'
          //sh("git commit -am 'update: build ${env.BUILD_NUMBER} is successful'")
        } 
        sh 'printenv'
      }
    }
    stage('Build AMI') {
      when {
        branch 'master' 
      }
      steps {
        build job: 'samuel-nwoye-website-ami'
      }
    }
    stage ('Clean Workspace'){
      steps {
        cleanWs()
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