# Thunderdome CLI

## Join

Your first step is to request access to Thunderdome:

```sh
{{ . }} join --email=<email address> --github=<github username>
```

This will create credentials for your user and store them on your machine.
An administrator will then accept or reject your request.
Once accepted, you will receive an email with a link to verify your identity.

## Status

At any point, you can check your user status:

```sh
{{ . }} status
```

This will display your user's status and your current ticket.

## List

To get a list of claimable tickets, run:

```sh
{{ . }} list --unclaimed
```

This will display a list of all unclaimed tickets with information such as identifiers and description.

## Claim

Once you have found a ticket you would like to work on, you can claim it:

```sh
{{ . }} claim <ticket id>
```

This will set up a branch for you to work on, as well as a PR where you can discuss the ticket.
When you are done with the ticket, you can request a review in the PR.
If someone has worked on the ticket before you, make sure to look at what has already been done on the branch.

## Drop

If the ticket is merged, or you no longer want to work on it, you can drop it:

```sh
{{ . }} drop
```

Note that you can only be assigned one ticket at a time. Similarly, only one person can work on a ticket at a time.

## Leave

If you no longer want to be a part of Thunderdome, you can leave:

```sh
{{ . }} leave
```

This will remove your user and you will have to join again to get access.
If you have claimed a ticket, it will be dropped.
