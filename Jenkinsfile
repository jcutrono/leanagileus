#!groovy

import groovy.json.JsonOutput

node {	
	
	stage ('checkout') {			
		git url: 'https://github.com/jcutrono/leanagileus', branch: "${env.BRANCH_NAME}"
		sh 'git clean -fdx; sleep 4;'
	}
	
	def goTool = tool 'go'	
	
	stage ('build') {	
		env.GOPATH="${env.HOME}/go-plugins"
		
		sh "go get github.com/gorilla/mux"
		sh "go get gopkg.in/mgo.v2"
		sh "go build -i"
	}
	
	stage ('test') {		
		sh "go test"
	}
	
	if(env.BRANCH_NAME == "develop"){
		notifySlack("build succeeded")
		
		stage ('merge to master') {
			withCredentials([[$class: 'UsernamePasswordMultiBinding', credentialsId: 'git',
						usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD']]) {
				sh "git checkout master"
				sh "git pull origin master"
				sh "git merge develop"
				sh "git push https://$USERNAME:$PASSWORD@github.com/jcutrono/leanagileus.git master"
			}
		}
	}
	
	if(env.BRANCH_NAME == "master"){
		notifySlack("deploying to production")
		
		stage ('deploy production') {
			def output = sh returnStdout: true, script: 'git remote | grep deploy'
			sh "echo ${output}"
			
			if(output.isEmpty()) {
				sh "echo 'dokku remote does not already exist';"
				sh "git remote add deploy dokku@ec2-54-202-56-172.us-west-2.compute.amazonaws.com:leanagileus"
			}
			else {
				sh "echo 'dokku already exist';"				
			}
			
			sh "git push deploy master"
		}
	}
}


def notifySlack(text) {
    def slackURL = 'https://hooks.slack.com/services/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'
    def payload = JsonOutput.toJson([text      : text,
                                     channel   : "#ci",
                                     username  : "jenkins",
                                     icon_emoji: ":jenkins:"])
    sh "curl -X POST --data-urlencode \'payload=${payload}\' ${slackURL}"
}
