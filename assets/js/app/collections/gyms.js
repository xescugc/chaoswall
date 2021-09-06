define([
  "backbone",
  "utils",
  "models/gym",
], function(
  Backbone,
  Utils,
  Gym,
) {
  return Backbone.Collection.extend({
    url: "/gyms",
    model: Gym,
    parse: Utils.parseIfData,
  });
})
