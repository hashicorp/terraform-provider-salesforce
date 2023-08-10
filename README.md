
<!-- archived-provider -->
>Please note: This Terraform provider is archived, per our [provider archiving process](https://terraform.io/docs/internals/archiving.html). What does this mean?
>1. The code repository and all commit history will still be available.
>1. Existing released binaries will remain available on the releases site.
>1. Issues and pull requests are not being monitored.
>1. New releases will not be published.
>
>If anyone from the community or an interested third party is willing to maintain it, they can fork the repository and [publish it](https://www.terraform.io/docs/registry/providers/publishing.html) to the Terraform Registry. If you are interested in maintaining this provider, please reach out to the [Terraform Provider Development Program](https://www.terraform.io/guides/terraform-provider-development-program.html) at *terraform-provider-dev@hashicorp.com*.

<a href="https://terraform.io">
    <img src=".github/terraform_logo.svg" alt="Terraform logo" title="Terraform" align="right" height="50" />
</a>

# Terraform Provider Salesforce

![Status: Tech Preview](https://img.shields.io/badge/status-experimental-EAAA32) [![Releases](https://img.shields.io/github/release/hashicorp/terraform-provider-salesforce.svg)](https://github.com/hashicorp/terraform-provider-salesforce/releases)
[![LICENSE](https://img.shields.io/github/license/hashicorp/terraform-provider-salesforce.svg)](https://github.com/hashicorp/terraform-provider-salesforce/blob/main/LICENSE)![Tests](https://github.com/hashicorp/terraform-provider-salesforce/workflows/Tests/badge.svg)

This Salesforce provider for Terraform allows you to manage Users, Profiles, and User Roles.

This provider is a technical preview, which means it's a community supported project. It still requires extensive testing and polishing to mature into a HashiCorp officially supported project. Please [file issues](https://github.com/hashicorp/terraform-provider-salesforce/issues/new/choose) generously and detail your experience while using the provider. We welcome your feedback.

## Maintainers

This provider plugin is maintained by the Terraform team at [HashiCorp](https://www.hashicorp.com/)

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) >= 1.0.3
-	[Go](https://golang.org/doc/install) >= 1.19

## Setup

The provider interacts with the Salesforce REST API via a "connected app". Follow the steps described in the [Provider Configuration Reference](https://registry.terraform.io/providers/hashicorp/salesforce/latest/docs) for full details.

## Upgrading the provider

The Salesforce provider doesn't upgrade automatically once you've started using it. After a new release you can run

```bash
terraform init -upgrade
```

to upgrade to the latest stable version of the Salesforce provider. See the [Terraform website](https://www.terraform.io/docs/configuration/providers.html#provider-versions)
for more information on provider upgrades, and how to set version constraints on your provider.

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command or `make build`:
```sh
$ make build
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using The provider

See the [Salesforce Provider documentation](https://registry.terraform.io/providers/hashicorp/salesforce/latest/docs) to get started using the
Salesforce provider.

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).
You can use [goenv](https://github.com/syndbg/goenv) to manage your Go version.
To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `make generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and can cost money to run.

```sh
$ make testacc
```

For guidance on common development practices such as testing changes, see the [contribution guidelines](https://github.com/hashicorp/terraform-provider-salesforce/blob/main/.github/CONTRIBUTING.md).
If you have other development questions we don't cover, please file an issue!

## Experimental Status

By using the software in this repository (the "Software"), you acknowledge that: (1) the Software is still in development, may change, and has not been released as a commercial product by HashiCorp and is not currently supported in any way by HashiCorp; (2) the Software is provided on an "as-is" basis, and may include bugs, errors, or other issues; (3) the Software is NOT INTENDED FOR PRODUCTION USE, use of the Software may result in unexpected results, loss of data, or other unexpected results, and HashiCorp disclaims any and all liability resulting from use of the Software; and (4) HashiCorp reserves all rights to make all decisions about the features, functionality and commercial release (or non-release) of the Software, at any time and without any obligation or liability whatsoever.
