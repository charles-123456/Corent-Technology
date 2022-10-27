
##

The directory structure of the project development is consistent, easy to understand and easy to manage.

## Directory Structure

-`/APP/`Server work directory

`/app/controller/` routing entrance Action layer

`/APP/Common/` Store public functions and constant definitions

`/app/dao/` database table instance layer table instance layer

`/app/service/` business logic layer

-`/uWeb/`Management work directory

`/uWeb/CUSTOM/` web required static file directory

`/uweb/lib/` web -end framework directory

`/uweb/src/` web development directory

`/uWeb/src/pages/` web all pages of all pages of all pages

`/uweb/src/plugins/` `` Web -end custom plug -in directory

`/uweb/src/app.json` web configuration file

`/uweb/src/app.less` web global style file

`/uweb/dist/` web packed the static file directory

-`/time_series_detector/`time sequence abnormal detection academic directory

`/time_series_detector/model/` model file storage directory

`/time_series_detector/algorithm/` algorithm layer layer

`/time_series_detector/feature/` feature layer

Support the following types of files in the project:
    1. `.json`: configuration file
    2. `.uwx`: uweb view file
    3. `.uw`: uweb logic script
    4. `.js`: ordinary JavaScript logic script
    5. `.TS`: ordinary TypeScript logic script
    6. `.less`: LESS style file
    7. `.css`: CSS style file
    8. `.jsx`: JavaScript React script file that can be used when developing custom plugins
    9. `.TSX`: TypeScript React script file that can be used when developing custom plugins
    10. `.png`,` .jpg`, `.gif`,` .svg`: picture file

-`/docs/`project document storage directory


## call relationship

`uweb` is the management of the management side, the service side interface can be called

`/APP/Controller/` `the inlet of the server -side routing, the service layer can be called

`/APP/Service/` is the service layer of the service.

`/time_series_detector/` academic directory for service layers

