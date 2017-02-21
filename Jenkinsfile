node {
	// Mark the code checkout 'stage'....
	stage ('checkout') {
		// Get some code from a GitHub repository
		git url: 'https://github.com/jcutrono/leanagileus', branch: 'develop'
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

	stage ('merge to master') {
		sh "git checkout master"
		sh "git merge develop"
	}
	
	stage ('deploy Production') {
		input 'Proceed?'
		sh 'echo "write your deploy code here"; sleep 6;'
		archive 'target/*.jar'
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