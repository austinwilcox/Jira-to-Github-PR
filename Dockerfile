FROM golang:latest

LABEL "com.github.actions.name"="Jira to Github PR"
LABEL "com.github.actions.description"="Automate adding a Jira Tickets Description to Github Pull Requests"
LABEL "com.github.actions.icon"="git-commit"
LABEL "com.github.actions.color"="blue"

LABEL "repository"="https://github.com/austinwilcox/Jira-to-Github-PR
LABEL "homepage"="https://github.com/austinwilcox/Jira-to-Github-PR"
LABEL "maintainer"="Austin Wilcox <me@theaustinwilcox.com>"

workdir /app

COPY . /app

RUN go build -o Jira-to-Github-PR

EXPOSE 8080

CMD ["./Jira-to-Github-PR"]
