# GoTasker

GoTasker is a Go library for background task processing, providing a simple and efficient way to manage task execution
asynchronously using goroutines, Redis for task queuing, and logging with Logrus. It supports task cancellation and
timeout features, and provides an easy-to-use function-based API for registering and executing tasks.

> **Reason**: Im lazy to use the existing library, so I create my own library.

## Features

- **Concurrency**: Utilizes Go's goroutines to handle multiple tasks in parallel.
- **Timeout and Cancellation**: Supports cancellation and timeout using context.
- **Redis Integration**: Uses Redis as a task queue for handling background tasks.
- **Logging**: Integrated with [Logrus](https://github.com/sirupsen/logrus) for structured logging.
- **Modular**: Easily extendable for adding more features or custom task handlers.

## Installation

To use **GoTasker** in your Go project, run the following command:

```bash
go get github.com/SyahrulBhudiF/GoTasker
