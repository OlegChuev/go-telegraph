# Telegra.ph Go-lang API Client

## Introduction

This Go-lang API client provides a convenient way to interact with the [Telegra.ph](https://telegra.ph) service programmatically. [Telegra.ph](https://telegra.ph) is a minimalist publishing tool that allows users to create articles with rich formatting and no sign-up required.

Telegra.ph's API docs are located here: [API Docs](https://telegra.ph/api#Node)

## Installation

To use this API client in your Go project, you can simply install it using `go get`:

```bash
  go get github.com/OlegChuev/go-telegraph
```

## Usage

```go
  import telegraph "github.com/OlegChuev/go-telegraph/api/endpoints"

  tc := telegraph.NewClient()

  params := telegraph.CreateAccountParams{
    ShortName: "John Doe",
    AuthorName: "John Doe",
    AuthorUrl: "@my_url",
  }

  account, err := tc.CreateAccount(&params)
```
