# UUID CLI 🛠️

A simple and efficient Command-Line Interface (CLI) for generating and managing UUIDs. By default, the tool generates
UUIDs of Version 4, but it supports a wide range of UUID versions and additional functionality for parsing and
working with UUIDs.

## Features

- Generate UUIDs of various versions: V1, V3, V4, V5, V6, and V7
- Parse and validate UUIDs
- Support for generating multiple UUIDs at once

## Why is this tool Useful?

It simplifies tasks across various use cases:

- **Quick Generation**: Instantly create UUIDs for ad-hoc tasks.
- **Automation-Friendly**: Easily integrate into scripts, CI/CD pipelines, and automated workflows.
- **Cross-Platform**: Works consistently across different operating systems and environments.
- **Validation**: Quickly parse and verify UUIDs to ensure correctness and data integrity.

Ideal for developers and system administrators that need a lightweight, efficient, and always accessible tool for
working with UUIDs.

## Installation

// TODO: TBD

## Usage

Run the `uuid` command to generate a single V4 UUID by default:

```bash
$ uuid
```

### Commands

#### UUID Commands

- **`null`**  
  Outputs the null UUID.

  ```bash
  $ uuid null
  ```

- **`parse`**  
  Parses and validates a given UUID string.

  ```bash
  $ uuid parse e4eaaaf2-d142-11e1-b3e4-080027620cdd
  ```

- **`v1`**  
  Generates a Version 1 (timestamp-based) UUID.

  ```bash
  $ uuid v1
  ```

- **`v3`**  
  Generates a Version 3 (namespace-based, MD5 hash) UUID.

  ```bash
  $ uuid v3 "some value"
  ```

- **`v4`**  
  Generates a Version 4 (random) UUID. This is the default behavior.

  ```bash
  $ uuid v4
  ```

- **`v5`**  
  Generates a Version 5 (namespace-based, SHA-1 hash) UUID.

  ```bash
  $ uuid v5 "some value"
  ```

- **`v6`**  
  Generates a Version 6 UUID.

  ```bash
  $ uuid v6
  ```

- **`v7`**  
  Generates a Version 7 UUID. Epoch used is the current time on the machine.

  ```bash
  $ uuid v7
  ```

#### Additional Commands

- **`help`**  
  Displays help information about any command.

  ```bash
  $ uuid help
  ```

- **`version`**  
  Prints the current version of the CLI.

  ```bash
  $ uuid version
  ```

## Examples

### Generate Multiple UUIDs

```bash
$ uuid v7 -n 5
01947961-e155-7a32-82f1-1b2491f301ac
01947961-e155-7a33-ae7b-2a409d7388ab
01947961-e155-7a34-b168-c1a743900b1e
01947961-e155-7a35-b87a-71833dc172fc
01947961-e155-7a36-8374-9ecb5b7c0675
```

### Generate UUID with custom namespace

```bash
$ uuid v5 --namespace 869d7b84-d678-11ef-91a1-426648c33d81 "some value"
e21ac596-de47-5afa-a4c6-009662c4b663
```

### Generate UUID with custom epoch

```bash
$ uuid v7 --epoch 2025-01-18T13:10:05+01:00
01947952-0148-73d4-bca8-095cd1891884
```

### Parse a UUID

```bash
$ uuid parse 2733f45e-d595-11ef-b95f-426648c33d81
version: 1
time: 2025-01-18T13:10:05.633443+01:00
```

