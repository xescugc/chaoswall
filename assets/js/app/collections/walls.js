define([
  "backbone",
  "utils",
  "models/wall",
], function(
  Backbone,
  Utils,
  Wall,
) {
  return Backbone.Collection.extend({
    model: Wall,
    parse: Utils.parseIfData,
    url: function() {
      return this.gym.url()+"/walls"
    }
  });
});
