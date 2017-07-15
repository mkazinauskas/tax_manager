# tax_manager
This application onsists of two services. `service` and `consumer`
* `service` holds tax management
* `consumer` shows how tax management works

## Api usage
Service server is started on 8080 port

* GET `/` - index page
* GET `/municipalities` - displays all municipalities from database
* POST `/municipalities` - saves new municipality to database
* GET `/municipalities/{id}` - displays specific municipality by id (int64)
* GET `/municipalities/{id}/taxes` - displays all taxes for municipality
* POST `/municipalities/{id}/taxes` - saves new tax to database
* GET `/municipalities/{id}/taxes/{id}` - displays specific municipality by id and tax id (int64)
* GET `/calculate-tax?municipalityName={}&date={date}` - calculates tax for municipality on specified date


## What's done
* It has its own database where municipality taxes are stored `(Data is stored to MySQL database)`
* Taxes should have ability to be scheduled (yearly, monthly, weekly, daily) for each municipality
* Application should have ability to import municipalities data from file (choose one data format
you believe is suitable) `(CSV) format`
* Application should have ability to insert new records for municipality taxes (one record at a
time) `(Implemented via api calls, or just add new value to csv document, which is loaded on application startup)`

## What's pending
* User can ask for a specific municipality tax by entering municipality name and date
* Errors needs to be handled i.e. internal errors should not to be exposed to the end user
You should ensure that application works correctly

## What could be improved
* Database queries could be paginated
* Api could have REST 3rd level with hypermedia
* Dto structs could be use to expose to api, instead of domain structs
* Smarter validation of request objects
