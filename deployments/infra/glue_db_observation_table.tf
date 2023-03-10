resource "aws_glue_catalog_table" "observeration" {
  name          = "observation"
  database_name = aws_glue_catalog_database.object_db.name

  table_type = "EXTERNAL_TABLE"

  parameters = {
    EXTERNAL              = "TRUE"
  }

  storage_descriptor {
    location      = "s3://${local.observation_bucket_name}/data"
    input_format  = "org.apache.hadoop.mapred.TextInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.HiveIgnoreKeyTextOutputFormat"

    ser_de_info {
      name                  = "json"
      serialization_library = "org.openx.data.jsonserde.JsonSerDe"

      parameters = {
        "serialization.format" = 1
      }
    }

    columns {
      name = "At"
      type = "string"
    }


    columns {
      name = "station"
      type = "struct<macaddress:string,name:string>"
    }

    columns {
      name="pressure"
      type="struct<barometer:float>"
    }

    columns {
      name="humidity"
      type="struct<dewpoint:float,humidity:float,dewpointindoor:float,humidityindoor:float>"
    }

    columns {
      name="temperature"
      type="struct<feelslike:float,feelslikeindoor:float,fahrenheit:float,fahrenheitindoor:float>"
    }

    columns {
      name="airquality"
      type="struct<pm25:float,pm25indoor:float,pm25daily:float,pm25dailyindoor:float>"
    }

    columns {
      name="solar"
      type="struct<solarradiation:float,uvindex:float>"
    }

    columns {
      name="wind"
      type="struct<direction:float,gust:float,gustdirection:float,speed:float,directionaverage2minutes:float,speedaverage2minutes:float,directionaverage10minutes:float,speedaverage10minutes:float,maxdailygust:float>"
    }

    columns {
      name="rain"
      type="struct<daily:float,event:float,hourly:float,monthly:float,total:float,yearly:float,weekly:float>"
    }
  }
}