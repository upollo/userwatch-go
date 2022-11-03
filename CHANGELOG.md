# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),

## [0.2.1]

### Changed

- module renamed to `github.com/upollo/userwatch-go`

## [0.2.0]

### Changed

- Package renamed from `userwatchgo` to `upollo`
- `Validate` was renamed to `Verify`.
  - `client.Validate(ctx, request)` becomes
  - `client.Verify(ctx, request)`

### Removed

- the flag type `REPEATED_ACTION` has been removed.

### Added

- A retry was added to prevent intermittent network errors.
