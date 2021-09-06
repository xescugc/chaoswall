define([
  "backbone",
  "utils",
  "models/route",
], function(
  Backbone,
  Utils,
  Route,
) {
  return Backbone.Collection.extend({
    model: Route,
    parse: Utils.parseIfData,
    url: function() {
      return this.wall.url()+"/routes"
    }
  });
});
