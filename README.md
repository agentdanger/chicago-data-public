# chicago-data | full stack reporting solution

## Project Overview

<!--- add badge once we get going - see here -> https://docs.github.com/en/actions/monitoring-and-troubleshooting-workflows/adding-a-workflow-status-badge -->

This repo houses microservices dedicated to ingesting and preparing open Chicago datasets for analysis. This project references data from the City of Chicago's Data Portal. [CDP Here](https://data.cityofchicago.org/)

The project will leverage the Go language for all microservices and Python for some ETL services (including GIS processing.)  Our database choices include Elasticsearch for document storage and PostgreSQL for structured data storage.

## Requirements

- Docker (built using 20.10.11)
- Kubernetes (built using v1.22.3)
- (all other requirements are containerized)

## Data Sources

- [ ] City of Chicago 311 Service Requests [source](https://data.cityofchicago.org/Service-Requests/311-Service-Requests/v6vf-nfxy)
- [ ] City of Chicago Business Licenses [source](https://data.cityofchicago.org/Community-Economic-Development/Business-Licenses/r5kz-chrr)
- [X] City of Chicago Crime 2001-present [source](https://data.cityofchicago.org/Public-Safety/Crimes-2001-to-Present/ijzp-q8t2)
- [ ] City of Chicago Taxi Trips [source](https://data.cityofchicago.org/Transportation/Taxi-Trips/wrvz-psew)
- [ ] City of Chicago Traffic Crashes [source](https://data.cityofchicago.org/Transportation/Traffic-Crashes-Crashes/85ca-t3if)
- [ ] City of Chicago Transporation Network Provider Trips (Uber, Lyft, etc.) [source](https://data.cityofchicago.org/Transportation/Transportation-Network-Providers-Trips/m6dm-c72p)
- [ ] Chicago Transit Authority "L" Station Entries [source](https://data.cityofchicago.org/Transportation/CTA-Ridership-L-Station-Entries-Daily-Totals/5neh-572f)
- [ ] Chicago Transit Authority List of "L" Stops [source](https://data.cityofchicago.org/Transportation/CTA-System-Information-List-of-L-Stops/8pix-ypme)
- [X] Chicago Population by ZIP (batch loaded yearly) [source](https://data.cityofchicago.org/Health-Human-Services/Chicago-Population-Counts/85cm-7uqa) and [source2: data source in link](https://en.wikipedia.org/wiki/Community_areas_in_Chicago)

## Reports Available

- [X] Crime incidents by type, date and zip code [API](http://abcf662db43574fd99b224e0ab41b018-1792580361.us-east-2.elb.amazonaws.com:8080/)
- [ ] Crime trends for each zip code

## Credits

- City of Chicago Data Portal [source](https://data.cityofchicago.org/)
- SODA Developers [source](https://dev.socrata.com/)
- Northwestern University SPS [source](https://sps.northwestern.edu/) *for inspiration for this project*
