# E2E File Storshare

[![License](https://img.shields.io/github/license/RyanLake6/E2E-file-storshare)](https://github.com/RyanLake6/E2E-file-storshare/blob/main/LICENSE)

## Purpose

CLI tool meant to E2E store and easily share files, built with nextcloud

So why am I making this? Well... nextcloud seems to have server side ecryption but with the default encryption module they provide once you starts encrypting all uploaded data you can't turn it off. Server side encryption takes more time and more space for all files once turned on. So what if you wanted to encrypt some things but not others? Or you don't fully trust their encryption and want to encrypt ontop of theirs as well for very secure documents? This is the solution!

Also what if you want to quickly upload a file or share a file without going through their UI? Well... this tool would allow you to use quick commands to get a shareable link. It would also give easy commands to share the link through third party tools to sms or email to a colleauge.

_TLDR_: I'm making this to provide users with more options and ease of use

## Installation

Clone the repository or download the script.

```bash
git clone https://github.com/RyanLake6/E2E-file-storshare.git
```

Navigate to the directory where the script is located.

```bash
cd E2E-file-storshare
```

Build the repo to make an executable file to run commands

```bash
go build
```

## Quick Use Guide

The following command will give CLI tool usage and full command help list

```bash
./E2E-file-storshare.exe
```

**To view all commands see [here](#commands)**

## Developer Notes:

This was built with a huge thanks to the nextcloud API documentation [here](https://docs.nextcloud.com/server/latest/developer_manual/client_apis/index.html)

### Nextcloud Setup:

I recommend installing as a plugin ontop of TrueNAS, docs can be found [here](https://www.truenas.com/docs/solutions/integrations/nextcloud/) as this provides:

- Centralized storage services within TrueNAS
- On boot startup
- Automatic Updates

### Server-Side Encryption

As I mentioned nextcloud gives you the ability to enable server-side encryption across all your files. I recommend you fully understand your use case as it may NOT be reversable. Please read what nextcloud has to say about this [here](https://docs.nextcloud.com/server/latest/admin_manual/configuration_files/encryption_configuration.html).

I am NOT liable for any encryption shortcomings if you utilize this codebase. Please understand this should not be used in any production level solutions and should be recognized as a non-official nextcloud tool.

### Storing your Secrets?

Don't worry the code doesn't take in any usernames or passwords utilizing nextcloud login flow v2 during the `login` command. Read more [here](https://docs.nextcloud.com/server/latest/developer_manual/client_apis/LoginFlow/index.html#login-flow-v2)

The session token will be stored locally in a config file on your device and is only valid for 20 minutes.

**NOTE**: On windows the file is stored at C:/Users/\<user>/.e2e-file-storshare-cli and this may not work on Linux or Mac. To debug with your OS please look within config/config.go to change where this is stored

## Commands:

### Login

```bash
# Login to Nextcloud with specified credentials
./E2E-file-storshare.exe login --base-url https://<your-nextcloud-instance.com>
```

#### Flags

- `--base-url` https://<your-nextcloud-instance.com> (required): The base url for the Nextcloud instance running.
- `--debug`, `-d` (optional): If supplied, returns debugging information.

### List

```bash
# List all Nextcloud files and folders
./E2E-file-storshare.exe list --remote-path <path-here>
```

#### Flags

- `--remote-path` /\<folder-location> (required): The remote location you wish to view the files from (use `/` for the root folder).
- `--debug`, `-d` (optional): If supplied, returns debugging information.
- `--all-details`, `-a` (optional): If supplied, returns all file information.

### Share

```bash
# List all Nextcloud files and folders
./E2E-file-storshare.exe share --remote-path <path-here> --permissions <permission-int-here>
```

#### Flags

- `--remote-path` /\<folder-location> (required): The remote folder or file you wish to share
- `--permissions` \<integer> (required): 1 = read; 2 = update; 4 = create; 8 = delete; 16 = share; 31 = all (default: 31, for public shares: 1)
- `--debug`, `-d` (optional): If supplied, returns debugging information.

### List-Shares

```bash
# List all Nextcloud files and folders
./E2E-file-storshare.exe list-shares
```

#### Flags

- `--debug`, `-d` (optional): If supplied, returns debugging information.
- `--all-details`, `-a` (optional): If supplied, returns all share information.

### Delete-Share

```bash
# List all Nextcloud files and folders
./E2E-file-storshare.exe delete-share --share-id <share-id-here>
```

#### Flags

- `--share-id` (required): The id of the share you wish to delete
- `--debug`, `-d` (optional): If supplied, returns debugging information.

### Upload

```bash
# Upload the local file to nextcloud
.\E2E-file-storshare.exe upload --remote-path /Documents/notes.txt --local-path "C:\Users\username\MyDocuments\notes"
```

#### Flags

- `--remote-path` (required): remote path to place the local file to in nextcloud
- `--local-path`, (required): local path of the file to upload (absolute path is easiest to avoid issues)