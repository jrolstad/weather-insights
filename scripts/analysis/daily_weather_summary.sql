select
    station.name as StationName,
    station.macaddress as StationAddress,
    cast(from_iso8601_timestamp(at) as date) as ObservedAt,
    max(rain.daily) as TotalRain,
    max(temperature.fahrenheit) as MaxTemperature,
    min(temperature.fahrenheit) as MinTemperature,
    max(wind.speed) as MaxWindSpeed,
    min(wind.speed) as MinWindSpeed,
    max(wind.gust) as MaxWindGust,
    max(pressure.barometer) as MaxBarometerPressure,
    min(pressure.barometer) as MinBarometerPressure,
    max(humidity.humidity) as MaxHumidity,
    min(humidity.humidity) as MinHumidity,
    max(humidity.dewpoint) as MaxDewpoint,
    min(humidity.dewpoint) as MinDewpoint,
    max(solar.solarradiation) as MaxSolarRadiation,
    min(solar.solarradiation) as MinSolarRadiation,
    max(solar.uvindex) as MaxUvIndex,
    min(solar.uvindex) as MinUvIndex
from observation
group by cast(from_iso8601_timestamp(at) as date), station.name, station.macaddress