# emailVerifier

The emailVerifier project is an application designed to verify validity of email addresses interactively. This application facilitates users with efficient validation of email addresses in real-time.

### Running project locally

#### Prerequisites
- [pkl](https://pkl-lang.org/main/current/pkl-cli/index.html#installation)
- [pkl-gen-go](https://pkl-lang.org/go/current/quickstart.html)

#### Setup

```console
git clone git@github.com:pstano1/email-verifier.git
cd email-verifier
go mod tidy && go mod download
make config
go run ./cmd
```