define([
  "backbone",
  "utils",
  "collections/walls"
], function(
  Backbone,
  Utils,
  WallsCollection,
) {
  return Backbone.Model.extend({
    idAttribute: "canonical",
    urlRoot: "/gyms",
    parse: Utils.parseIfData,
    initialize: function() {
      this.walls = new WallsCollection();
      this.walls.gym = this;
    },
  });
})
