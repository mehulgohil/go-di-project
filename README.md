# Go Dependency Injection with Swagger, Gin, and Zap Logger

![Go](https://img.shields.io/badge/Go-1.20-blue.svg)
![Swagger](https://img.shields.io/badge/Swagger-1.16.1-green.svg)
![Gin](https://img.shields.io/badge/Gin-1.9.1-green.svg)
![Zap Logger](https://img.shields.io/badge/Zap%20Logger-1.24.0-blue.svg)

This repository demonstrates a project structure for implementing Dependency Injection in a Golang application, along with the usage of Swagger for API documentation, Gin as the HTTP framework, and Zap Logger for logging.

## Introduction

Dependency Injection is a design pattern that allows for loosely coupled code by removing the direct dependencies between components. It helps improve testability, maintainability, and flexibility of the codebase. This repository demonstrates how to organize your Golang project using the Dependency Injection pattern.

Additionally, the project uses Swagger to automatically generate API documentation, Gin as the HTTP framework for routing and middleware, and Zap Logger for logging the application's events.

## Api Documentation

The API documentation is automatically generated using Swagger. To access the Swagger UI and explore the API endpoints, run the application and open the following URL in your web browser:
```shell
http://localhost:8080/swagger/index.html
```
