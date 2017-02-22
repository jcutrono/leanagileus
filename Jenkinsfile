#!groovy

import groovy.json.JsonOutput

node {
	sh "echo ${env.BRANCH_NAME}"
	
	stage ('checkout') {
		// Get some code from a GitHub repository		
		git url: 'https://github.com/jcutrono/leanagileus', branch: '${env.BRANCH_NAME}'
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
	
	notifySlack("build succeeded")
	
	if(env.BRANCH_NAME == "develop"){
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
		stage ('deploy production') {
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