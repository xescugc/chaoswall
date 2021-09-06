define([
  "backbone",
  "utils",
  "collections/routes"
], function(
  Backbone,
  Utils,
  RoutesCollection,
) {
  return Backbone.Model.extend({
    idAttribute: "canonical",
    parse: Utils.parseIfData,
    initialize: function() {
      this.routes = new RoutesCollection();
      this.routes.wall = this;
    },
  });
})
