@(appJs : String, templates : String, constants : play.api.libs.json.JsValue)

@format.raw(appJs)
module.constant('constantData',@format.raw(constants.toString));
module.run(['$templateCache',function(t){@format.raw(templates)}]);
