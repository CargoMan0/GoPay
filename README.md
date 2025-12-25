# GoPay Service

## Overview
TODO: Write

## Architecture
### 1.1 General Project Architecture
![gopay_arch.png](assets/gopay_arch.png)
### 1.2 Flows description
#### 1.2.1 User registration
![gopay_registration_flow.png](assets/gopay-registration-flow.png)
1. Client initialises an account registration.
2. The request is sent via HTTP v2 to API Gateway.
3. API Gateway validates and convert data to Protobuf, then gRPC request is sent to Auth Service.
## Services
TODO: Describe

### Gateway Service
