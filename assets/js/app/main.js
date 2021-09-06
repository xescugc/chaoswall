define([
  "jquery", "underscore", "backbone",
  "router",
], function(
  $, _, Backbone,
  Router,
) {
  Backbone.ajax = function(request) {
    request = _({ contentType: 'application/json' }).defaults(request);
    return Backbone.$.ajax.call(Backbone.$, request);
  };

  Backbone.history.start({pushState: true});
});
