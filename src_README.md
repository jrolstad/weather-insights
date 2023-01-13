# Weather Insights
The components in Weather Insights enable the collection and storage of large amounts of weather data so it is available for further analysis. Components are primarily written in Golang and use AWS resources for processing, data storage, and analytical capabilities.

# Requirements
* golang 1.18 or higher

# Components
## Applications

|Name|Location|Purpose|
|---|---|---|
|Command Line Interface|cmd/cli|Command line tool that reads weather data and pushes to the AWS resources for processing|
|Cron Loader|cmd/lambda/cron_loader|Lambda function that runs on a regular basis and pulls weather data from the source api and pushes to an S3 bucket|

## Libraries

|Name|Location|Purpose|
|---|---|---|
|clients|internal/clients|Implementations around interacting with different directory providers and other external APIs / services|
|config|internal/config|Reads configuration data|
|core|internal/core|Extensions to the native types in golang|
|logging|internal/logging|Wrapper around logging code to ensure all items are logged the same|
|models|internal/models|Data models / structs used in the components|
|orchestration|internal/orchestration|The main entry point for all exposed functionality. Contains business logic, processing, and flow implementations|
|repositories|internal/repositories|Data repositories used to read and update application data|

# Infrastructure
The infrastructure used to implement this solution is defined with Terraform HCL.  See the [deployments/infra](deployments/infra/) directory for more details.

A sample deployment pipelien for this can be seen at [terraform-deploy-prd](.github/workflows/deploy-prd.yml)

# Pipelines
The following pipelines are used by this repository to build, test, and deploy a sample instance.  They are:
|Name|Location|Purpose|
|---|---|---|
|application-continuous-integration|[.github/workflows/application-ci.yml](.github/workflows/application-ci.yml)|Application code (golang) CI build|
|terraform-continuous-integration|[.github/workflows/terraform-ci.yml](.github/workflows//terraform-ci.yml)|Infrastructure code (terraform) CI Build|
|terraform-deploy-prd|[.github/workflows/deploy-prd.yml](.github/workflows/deploy-prd.yml)|Sample deployment pipeline to an AWS Account|

# How To Use
To start using Weather Insights on your own resources, use the following steps

## 1. Create Cloud Resources
Using the infrastructure defined via HCL code in the [deployments/infra](deployments/infra/) folder, run this on an AWS Account that you own so the resources are properly configured

## 2. Populate Client Id and Client Secret
For every weather source where there is a Cliend Id and Client Secret, retrieve them from the respective platforms and place them into AWS Secrets Manager

### Example
|Secret Name|Secret Value|
|---|---|
|weather_insights_prd/application_key|the-application-key|
|weather_insights_prd/api_key|superSecretValue!|

## 3. Configure Local AWS Credentials
Once the AWS resources are configured, credentials on the machine that will be running the console app need to be configured.  In your machine's environment variables, ensure there are valid AWS credentials for the Account where the resources created in step 1 are located.

### Example
```
export aws_region=us-west-2
export AWS_ACCESS_KEY_ID=AKIAGDBMSF5KR367TZX
export AWS_SECRET_ACCESS_KEY=somesupersecretvalue
```

## 4. Configure Environment Variables
Additional variables to configure for the application to run locally are below.  Values are for illusstrative purposes only; be sure they match what you have set up in previous steps.

### Client Id Secret Name
Name of the client id / application id to authentiate to weather apis with

```
export weather_application_key_name=weather_insights/prd/application_key
```

### Client Secret Name
Name of the client secret to authenticate to the weather apis with
```
export weather_api_key_name=weather_insights/prd/api_key
```

### S3 Destination Bucket Name
Name of the S3 bucket to send data to
```
export weather_observation_bucket_name=weatherobservationdata
```

# 5. Execute Command Line Application
When the AWS resources are setup and credentials set locally, the command line app can be executed.  This app reads data from one or more api providers and pushes to an S3 bucket for later analysis.  

In the cmd/cli directory, below are examples of commands to run to read directory data.

```
go run main.go
```

## 6. Start Gaining Insights
Using the data exposed via queries and Tableau workbooks, start gaining insights on the data!

# Development Environment Setup
To set up your development environment, follow the steps above in the _How to Use_ section.
Once complete, you should be able to open this solution in an editor of your choice and start making changes / running on your own.