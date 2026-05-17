# HTTP Server

A lightweight HTTP/1.1 server written in Go, built entirely from scratch without using the standard `net/http` package.

---

## Overview

This project implements the core building blocks of an HTTP server, focusing on understanding how the protocol works under the hood.

---

## Features

- HTTP request parsing
- HTTP response generation
- Concurrent request handling
- Manual connection management (no `net/http` dependency)

---

## Goals

The purpose of this project is to deepen understanding of:

- The HTTP/1.1 protocol
- TCP connection handling in Go
- Concurrency and request lifecycle management
- Low-level server design without abstractions

---

## Notes

This is a learning-focused implementation and is not intended for production use.