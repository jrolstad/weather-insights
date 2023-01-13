# Weather Insights
The weather insights solution is a collection of components that enable people and organizations with weather stations to capture, analyze, and gain insights into weather conditions and trends.  By leveraging the capabilities of big data offerings to quickly and easily analyze large amounts of data, a comprehensive set of insights and trends can be obtained across all supported weather providers.

# Supported Platforms
The following platforms are supported

|Name|Status|
|---|---|
|[Ambient Weather](https://ambientweather.com/)|Available|

# How it Works
For every onboarded platform, data about each weather station is queried in an incremental manner and pushed to a secure Amazon S3 bucket for storage. Once the data is collected, items in the S3 buckets are wrapped by big data tools such as Hive (Amazon Glue Database) and Presto (Amazon Athena) so they can be queried in real time.

Using queries defined in the scripts/analytics directory, this data can then exported to providers such as Tableau to gain insights from the raw data. A sample set of analytical workbooks are available at scripts/tableau/workbooks to showcase what is possible.

# How To Use
The Weather Insights solution is broken down into two pieces - data collection and analytics.

## Components
 ![Component Diagram](/docs/components.png)

## Data Collection
To install, configure, and collect data so it is available for analysis, follow information in the [src_ReadMe](src_README.md)

## Analytics
When data is collected, use the [scripts/analysis](scripts/analysis) to extract data either locally or to a Tableau Server so it is available for the pre-built dashboards at [scripts/tableau/workbooks](scripts/tableau/workbooks/).

# License
This projects is made available under the [MIT License](LICENSE).