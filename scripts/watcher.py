from github_webhook import Webhook
from flask import Flask
import sys
import subprocess
import os
from git import Repo, Remote

#print('This is error output', file=sys.stderr)
#print('This is standard output', file=sys.stdout)

app = Flask(__name__)  # Standard Flask app
webhook = Webhook(app) # Defines '/postreceive' endpoint

@app.route("/")        # Standard Flask endpoint
def root_endpoint():
    return "Nothing here."

global sub
@app.route("/",methods = ['POST'])
def on_push():
    currdir = os.getcwd()
    currdir = currdir.split("/")
    if currdir[-1] == 'OnlyOne':
        repo = Repo()
        currentCommit = repo.head.commit
        repo.remotes.origin.pull()
        if currentCommit != repo.head.commit:
            subprocess.run(["docker-compose","stop"])
            subprocess.run(["docker-compose","build"])
            subprocess.run(["docker-compose","up", "-d"])
        else:
            print("No change detected")
    else:
        Repo.clone_from("https://github.com/OneSock-inc/OnlyOne.git", "./OnlyOne")
        os.chdir("./OnlyOne")
        subprocess.run(["docker-compose","up", "-d"])
    return "tg post"

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=9999, debug=True)