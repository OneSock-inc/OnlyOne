from github_webhook import Webhook
from flask import Flask, request, json
import subprocess
import os
from git import Repo

#print('This is error output', file=sys.stderr)
#print('This is standard output', file=sys.stdout)

DEPLOYMENT_BRANCH = "feat_front_Docker-support"

app = Flask(__name__)  # Standard Flask app
webhook = Webhook(app) # Defines '/postreceive' endpoint

@app.route("/")        # Standard Flask endpoint
def root_endpoint():
    return "Nothing here."

@app.route("/",methods = ['POST'])
def on_push():
    ret = "No change detected"
    branch = json.loads(request.data)["ref"].split("/")[-1]
    if branch != DEPLOYMENT_BRANCH:
        return ret
    currdir = os.getcwd().split("/")[-1]
    #dockerps = subprocess.getoutput("docker compose ps --format json")
    #dockerps = json.loads(dockerps)
    #print(check_running(dockerps, "onlyoneFrontend"))
    if currdir == 'OnlyOne':
        repo = Repo()
        repo.git.checkout(DEPLOYMENT_BRANCH)
        currentCommit = repo.head.commit
        repo.remotes.origin.pull()
        if currentCommit != repo.head.commit:
            subprocess.run(["docker-compose","stop"])
            subprocess.run(["docker-compose","build"])
            subprocess.run(["docker-compose","up", "-d"])
            ret = "Docker reloaded"
    else:
        Repo.clone_from("https://github.com/OneSock-inc/OnlyOne.git", "./OnlyOne")
        os.chdir("./OnlyOne")
        subprocess.run(["docker-compose","up", "-d"])
        ret = "Docker started."
    return ret


def check_running(dockerpsOutput, serviceName):
    for service in dockerpsOutput:
        if service["Name"] == serviceName:
            return service["State"] == "running"
    return False

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=9999, debug=True)