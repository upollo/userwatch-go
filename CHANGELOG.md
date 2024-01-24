# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),

## [Unreleased]

### Changed

### Removed

### Added

## [0.2.3]

### Changed

- Updated various dependencies.

### Added

- TRIALED_ON_OTHER_ACCOUNT flag type, which is available for customers with billing
  and subscription data connected.
- COMMERCIAL_USER, PAYMENT_NAME_DIFFERS and LIMITED_DEVICE_INFORMATION flag types.

## [0.2.2]

### Removed

- `userwatch_shared.pb.go` and all its types removed, as they are no longer used.

### Added

- `.proto` files which define much of the API, for reference.

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
