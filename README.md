# ThreeAmigos Corp. (ThAmCo) E-commerce System

## Overview

This repository contains the codebase for an e-commerce platform, consisting of a Customer Web App and a Staff Web App. The system is designed to provide a seamless shopping experience for customers and efficient order and product management for staff members. This document provides an overview of the system architecture, functionality, and guidelines for developers.

## Features

### Customer Web App

- **Public Interface:**
  - Browse and filter products.
  - Loose search within product names and descriptors.

- **User Authentication and Registration:**
  - Secure sign-in and registration.
  - Profile management (name, delivery address, telephone number).

- **Customer Dashboard:**
  - View stock status of products.
  - Real-time funds availability check.
  - Place orders securely.
  - Request account deletion.

- **Email Notifications:**
  - Receive order creation and status update emails.
  - View order history and status.

### Staff Web App

- **Staff Authentication:**
  - Secure sign-in for Staff

- **Staff Dasboard:**
  - Display an overview of key metrics, quick actions, and product-related information for staff

- **Order Management:**
  - View orders that need to be dispatched.
  - Mark orders as dispatched with date and time recording.

- **Customer Management:**
  - View customer profiles, funds, and order history.
  - Delete a customer account with personal data erasure/anonymization.

## Architecture

### Frontend

- Customer Web App: React.js.
- Staff Web App: React.js.
- Containerization: Docker.
- Orchestration: Kubernetes.
- API Gateway: Nginx or API Gateway services.

### Backend

- Microservices Architecture: Golang.
- Databases: MongoDB or PostgreSQL.
- Message Queue: RabbitMQ or Apache Kafka.
- Elasticache: Redis for caching.
- Elasticsearch: For advanced search functionality.

### Tools and Technologies

- Authentication: OAuth 2.0
- Email Service: Integration with third-party email services (e.g., SMTP).
- Testing: Automated verification testing with test doubles.
- CI/CD: GitHub.
- Configuration Management: Kubernetes configurations.

## Security

- User authentication through industry-standard services.
- Data encryption with HTTPS.
- Role-based access control for sensitive operations.

## Testing

- Automated unit tests, integration tests, and end-to-end tests.
- Use of test doubles (mocks, stubs) for thorough testing.

## Continuous Integration and Deployment

- CI/CD pipeline for automated build, test, and deployment.
- Kubernetes for continuous deployment and scaling.

## Acknowledgments

- Special thanks to the developers who made this project possible.
- Inspiration from industry best practices and standards.

