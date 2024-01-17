# Lucit CLI
A command line tool, written in go, for testing the Lucit API for Digital Out of Home Applications

This tool requires a Lucit Application ID, an Authorization Token and a Secret

# Purpose
This tool allows 3rd party application developers to perform basic API calls to the Lucit API for testing their app id's, tokens, etc.

# API and Developer Documentation

- **Lucit Developer documentation** https://lucit.cc/developers
- **Complete Lucit API reference** https://apidocs.lucit.app

# Notes
This tool is built upon the following open source libraries

- **Cobra** CLI Framework https://github.com/spf13/cobra
- **Cobra CLI** for automating tasks related to generating commands - https://github.com/spf13/cobra-cli
- **Viper** For reading and writing config files - https://github.com/spf13/viper

# Pre-Reqs

**1. Golang**
Lucit CLI is built on **go** and installed using **git**.  You will need to have the following installed

- `git` installation instructions - https://git-scm.com/book/en/v2/Getting-Started-Installing-Git
- `go` installation instructions - https://go.dev/doc/install


**2. Application Id**
To use the Lucit API - You need an Application ID from the Lucit App you created.

For more information on building a Lucit app and getting your Application ID
See https://www.lucit.cc/developers/applications

**3. Authentication Token and Secret**
To access to the API, you need to generate an Authenticaion token and secret.
Tokens are created underneath the "TOKENS" tab of your application

Learn more about creating tokens and secrets at https://www.lucit.cc/post/creating-lucit-application-tokens

# Installation and Setup

**1. Clone the this Repo**

Clone this repo to a local directory on your machine

```
git clone https://github.com/lucit-cc/lucit-cli.git
```

Change to the `lucit-cli` directory

```
cd lucit-cli
```


**2. Initialize your Application**

This sets up your install with your Application ID, Token and Secret

It will also retrieve and store your long-lived Bearer Token.

Under the hood, it is making a call to the [/auth](https://apidocs.lucit.app/#auth-POSTapi-v3-auth) endpoint

```
go run main.go init
```


![init](docs/images/screenshot_init.jpg)



**3.  Test your Settings**

This will validate that your settings are correct and that you can access the API.

Under the hood it is getting status and app information by making a
call to the [/status](https://apidocs.lucit.app/#status-GETapi-v3-status) endpoint

```
go run main.go test
```

![test](docs/images/screenshot_test.jpg)


**4. Make your first API call**


```
go run main.go get /status
```

The response should look something like the following json

```json
{
  "api_version": "v3",
  "app": {
    "application_class": "App\\LuCore\\Applications\\PrivateApplicationClass",
    "application_class_description": "Only you and other members of this application can add it to accounts",
    "created_at": "2024-01-16T20:42:43.000000Z",
    "description": "This is my Lucit Hello World Application",
    "lcuid": "LCUID-LAP-5922ac88-****-*****-*****************",
    "name": "Hello World",
    "options": {
      "allowed_permissions_at_version": {
        "1": [
          "account.view"
        ],
        "2": [
          "account.view",
          "account.createContent"
        ],
        "3": [
          "account.view",
          "account.createContent",
          "account.viewContent"
        ]
      },
      "permissions": null,
      "permissions_version": 3,
      "primary_image_public_url": null
    },
    "organization_name": null,
    "permissions": {
      "allowed": [
        "account.view",
        "account.createContent",
        "account.viewContent"
      ]
    },
    "slug": "HelloWorld",
    "status": 0,
    "updated_at": "2024-01-16T20:43:20.000000Z",
    "website": null
  },
  "app_id_required": true,
  "auth_required": true,
  "auth_user_id_sent": "LCUID-LU-30cbafe4-****-****-****-*****************",
  "message": "LuCore V3 REST API is accessible with an v3_app_id, un-authenticated, and returns json",
  "ok": true,
  "timestamp": 1705438447,
  "user": {
    "lcuid": "LCUID-LU-30cbafe4-****-****-****-*****************",
    "name": "Hello World Token"
  },
  "v3_app_id_sent": "LCUID-LAP-5922ac88-****-*****-*****************"
}
```




**5. Build a binary**

Build and install the `lucit-cli` binary

```
go install
```

**6. Use the Binary**

- `lucit-cli init`
- `lucit-cli test`
- `lucit-cli get /status`


# GET

The `get` command simply accepts the endpoint that you are fetching and will return the JSON response from that endpoint

# POST and PUT (Not Supported)

These commands are NOT currently supported

`lucit-cli` can only be used (currently) to make requests to `GET` endpoints.

# Some quick Examples

Full API documentation of all endpoints is available at https://apidocs.lucit.app/

Here are a few examples GET endpoints to try

**Accounts**
Get a list of accounts that have added your app

```
lucit-cli get /accounts
```

```json
{
  "accounts": [
    {
      "created_at": "2023-11-01T13:17:31.000000Z",
      "description": null,
      "lcuid": "LCUID-LA-****-****-****-****-*****",
      "name": "Coastline Auto (Eric)",
      "options": {
        "primary_image_background_removed_public_url": null,
        "primary_image_public_url": null
      },
      "slug": "CoastlineAutoEric",
      "website": null
    }
  ],
  "success": true
}

```


**Campaigns**
Get a list of campaigns that belong to your accounts

```
lucit-cli get /campaigns
```

```json

{
  "campaigns": [
    {
      "active": true,
      "build_status": 2,
      "campaign_class": "App\\LuCore\\Campaigns\\OperatorContractCampaignClass",
      "campaign_class_description": "Media Owner Contract",
      "campaign_state": 6,
      "campaign_state_changed_at": null,
      "created_at": "2023-11-01T13:22:02.000000Z",
      "hash_id": "lch-4Cd6",
      "last_build_at": "2024-01-16 21:02:07",
      "lcuid": "LCUID-LE-****-****-****-****-*****",
      "name": "Breezy Billboards"
    }
  ],
  "success": true
}
```

# .lucit-cli.yaml

All configuration settings are stored in .lucit-cli.yaml in your home directory.

On linux this is typically in `~/.lucit-cli.yaml` and in Windows this is typically
`C:\Users\{YOURNAME}\.lucit-cli.yaml`

The settings in this file are as follows.

Changing id's, tokens and secrets in this file will result in unpredictable behavior

- `lucit_api_url` : This is the url to the API.  This should always be ` https://api.lucit.app/api/v3`
- `lucit_app_id` : This is your Lucit Application ID
- `lucit_app_token` : This is your Token
- `lucit_app_secret` :  This is your secret
- `lucit_oauth_token` : This is the oauth token that was genenerated when you ran `init`

# More Information

View all of the Lucit developer documentation at https://www.lucit.cc/developers


# About Lucit

Founded in 2019, The Lucit Platform allows customers to view, post, edit, manage,
and schedule their digital billboard creatives in real-time from their desktop or phone,
and brings connectivity to Automotive, Real Estate, and eCommerce systems by automatically generating
dynamic creatives from data for digital signage and digital screens.

Lucit is an open development platform for out-of-home, digital billboards, and digital signage

https://www.lucit.cc/
