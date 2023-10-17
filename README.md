# Investec Programmable Banking CLI

Provided "as is". No warranties or guarantees. Use at your own risk.

CLI to interact with Investec's programmable banking API.

## Setup

First, you'll need some Investec programmable banking credentials. You can either get your own credentials through your Investec Private Banking portal, or you can test the CLI using the sandbox credentials, found [here](https://developer.investec.com/za/api-products/documentation/SA_PB_Account_Information#section/Authentication/Oauth2-Sandbox).

Next, you'll need to create a `investec-cli.yaml` file in your `$HOME` directory:

```yaml
apikey: eUF4elFSRlg5N3ZPY3lRQXdsdUVVNkg2ZVB4TUE1ZVk6YVc1MlpYTjBaV010ZW1FdGNHSXRZV05qYjNWdWRITXRjMkZ1WkdKdmVBPT0=
apiurl: https://openapisandbox.investec.com
clientid: yAxzQRFX97vOcyQAwluEU6H6ePxMA5eY
secret: 4dY0PjEYqoBrZ99r
```

If you would like to connect to your actual Investec account, substitute the sandbox credentials for your actual credentials, and change the `apiurl` to `http//openapi.investec.com`.

This file is also used to cache the active OAuth token, which is valid for 30 minutes and will be refreshed automatically if it expires. If you update any of the credentials, be sure to remove the `token:` entry from the file so that the CLI can refresh the token.

## Usage

The CLI outputs all data in JSON format. You can pipe the output to `jq` for formatting and filtering.

Run `investec-cli --help` for a list of available commands.

Right now, the CLI expects you to choose a default account to execute commands on. First, list all available accounts by running `investec-cli accounts`, and then set the default account by running `investec-cli setAccount [accountID]`.

For information on other commands, simply run `investec-cli [command] --help`.
