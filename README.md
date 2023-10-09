# ghchk

Verify the validity of GitHub access tokens from the command line.

`ghchk` validates GitHub access tokens against the GitHub API. The input can be from command line arguments, a file, or stdin. Valid tokens, and their associated GitHub usernames, are output to stdout.

## Usage Example

Here a file named `tokens.txt` has a list of potential GitHub access tokens. Use ghchk to validate their authenticity:

```
▶ cat tokens.txt
abcd1234
efgh5678
ijkl9012

▶ ghchk -file=tokens.txt
Token efgh5678: Valid (User: johnDoe)
Token ijkl9012: Valid (User: janeDoe)
Valid tokens, with their associated GitHub usernames, are printed. You can also save this output to another file:

▶ ghchk -file=tokens.txt > valid-tokens.txt

▶ cat valid-tokens.txt
Token efgh5678: Valid (User: johnDoe)
Token ijkl9012: Valid (User: janeDoe)
```

You can also pass input through stdin or only one token via `-token` flag. 

## Flags
`-token=YOUR_ACCESS_TOKEN` : Check a specific token.

`-file=PATH_TO_FILE` : Validate tokens from a file.

`-h` or `-help` : Access help.

## Notes:

To display the results in stdout without writing to a file, redirect or process the output accordingly.
To provide tokens via stdin, pipe them directly into ghchk.

## Installation

### Install using Go
`go install -v github.com/itsmeashim/ghchk@latest`

Or download a [binary release](https://github.com/itsmeashim/ghchk/releases) for your platform.
