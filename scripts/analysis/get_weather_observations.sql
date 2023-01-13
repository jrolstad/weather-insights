select
    station.name as StationName,
    station.macaddress as StationAddress,
    cast(from_iso8601_timestamp(at) as timestamp) as ObservedAt,
    rain.hourly as RainHourly,
    temperature.fahrenheit as Temperature,
    temperature.feelslike as TemperatureFeelsLike,
    wind.direction as WindDirection,
    wind.speed as WindSpeed,
    wind.gust as WindGust,
    pressure.barometer as BarometerPressure,
    humidity.humidity as Humidity,
    humidity.dewpoint as Dewpoint,
    solar.solarradiation as SolarRadiation,
    solar.uvindex as UvIndex
from observation
where cast(from_iso8601_timestamp(at) as timestamp) > (current_date - interval '7' day)