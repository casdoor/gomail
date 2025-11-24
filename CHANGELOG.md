# Changelog

All notable changes to this project will be documented in this file. See [Conventional Commits](https://conventionalcommits.org) for commit guidelines.

## [2.3.0](https://github.com/casdoor/gomail/compare/v2.2.0...v2.3.0) (2025-11-24)


### Features

* add semantic release for automated versioning and releases ([#5](https://github.com/casdoor/gomail/issues/5)) ([5707df1](https://github.com/casdoor/gomail/commit/5707df1dd5bca49b5a913594e8fefa757c95a6fe))

# Change Log
All notable changes to this project will be documented in this file.
This project adheres to [Semantic Versioning](http://semver.org/).

## [2.0.0] - 2015-09-02

- Mailer has been removed. It has been replaced by Dialer and Sender.
- `File` type and the `CreateFile` and `OpenFile` functions have been removed.
- `Message.Attach` and `Message.Embed` have a new signature.
- `Message.GetBodyWriter` has been removed. Use `Message.AddAlternativeWriter`
instead.
- `Message.Export` has been removed. `Message.WriteTo` can be used instead.
- `Message.DelHeader` has been removed.
- The `Bcc` header field is no longer sent. It is far more simpler and
efficient: the same message is sent to all recipients instead of sending a
different email to each Bcc address.
- LoginAuth has been removed. `NewPlainDialer` now implements the LOGIN
authentication mechanism when needed.
- Go 1.2 is now required instead of Go 1.3. No external dependency are used when
using Go 1.5.
