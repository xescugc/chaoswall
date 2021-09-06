define([
  "backbone",
  "utils",
], function(
  Backbone,
  Utils,
) {
  return Backbone.Model.extend({
    idAttribute: "canonical",
    parse: Utils.parseIfData,
  });
});
