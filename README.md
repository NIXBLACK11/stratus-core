# Stratus

Stratus is a monitoring tool that allows users to set up various triggers for their websites and receive email alerts whenever errors or issues occur. The tool is designed to help developers and website administrators stay informed about the health and performance of their sites.

## Features
### Customizable Triggers: 
Set up specific triggers to monitor various aspects of your website, such as response codes, SSL certificate status, server response time, and more.
### Real-time Alerts: 
Receive immediate email notifications when an issue is detected based on the triggers you've configured.
### Detailed Reporting: 
Email alerts include detailed information about the error, such as the project name, site name, site URL, and the specific trigger that was activated.
### Easy Integration: 
Integrate Stratus into your existing workflow with minimal setup and configuration.
### Scalable: 
Suitable for monitoring multiple websites and services with support for numerous triggers.

## Supported trigger types include:

- ```WebsiteUnreachable```
- ```SSLCertificateExpiring```
- ```ServerResponseTimeHigh```
- ```BrokenLinks```
- ```InternalServerError```
- ```UnauthorizedAccess```
- ```DNSResolutionIssue```
- ```ServiceUnavailable```

## How to use the cli:
1. Download the stratus tool
- For linux [link](https://github.com/NIXBLACK11/stratus-cli/releases/tag/v1.0)
- In case of other operating systems build using
    ```
    git clone https://github.com/NIXBLACK11/stratus-cli.git
    cd stratus-cli
    go build -o stratus
    ```
2.  ```bash
    mv stratus /usr/local/bin
    ```
3. Usage:

    ```
    stratus help

    Usage:
        help                         Show available commands and options.
        login                        Login to your account
        signup                       Create a new account
        projects                     List all projects in your account.
        project <project_name>       View and manage urls monitored in a specific project.
        delete <project_name>        Delete the prject specified
        add <config.st>              Add details to an existing or create a new project
    ```

## Layout of configuration file
```js
PROJECTNAME projectName
TRIES 3 

SITENAME test site 11
SITEURL https://www.yurr.com
ALERTTYPE WebsiteUnreachable, BrokenLinks

SITENAME test21 site
SITEURL https://www.tedast2.com
ALERTTYPE BrokenLinks, InternalServerError

SITENAME test2 site
SITEURL https://www.tesat2.com
ALERTTYPE BrokenLinks, ServiceUnavailable, ServerResponseTimeHigh
```
Note: 
- Tries defines the number of times the user is notified in a day in case of trigger.
- The email will be sent to the email with which you sign with.

[stratus-api](https://github.com/NIXBLACK11/stratus-api)
[stratus-cli](https://github.com/NIXBLACK11/stratus-cli)