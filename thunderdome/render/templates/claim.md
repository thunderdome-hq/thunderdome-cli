{{ .Response.Message }}
You are now ready to work! Next steps:

1. Accept invitations to the GitHub repository and team:
    - https://github.com/settings/organizations
2. Clone the repository and check out the branch:
    - Using https: `git clone {{ .HttpsRepoCloneUrl }} --branch {{ .BranchName }}`
    - Using SSH: `git clone {{ .SshRepoCloneUrl }} --branch {{ .BranchName }}`
    - git@github.com:thunderdome-hq/thunderdome.git
3. Make sure to look at the branch history, you might pick up the work where someone else left off:
    - `{{ .BranchName }}`
4. If you have any questions, comment the pull request:
    -  {{ .PullRequestUrl }}
5. If you want to contact Hummy directly, email:
   {{ range .Ticket.Contacts }}
    - {{ . }}
     {{ else }}
     *None*
     {{ end }}

6. Resolve the ticket and mark the PR as ready for review in GitHub when you are done.
7. When the ticket is resolved, or you want to leave it to someone else, run:
    - `thunderdome drop`

The full ticket is displayed below.

{{ template "ticket" .Ticket }}
