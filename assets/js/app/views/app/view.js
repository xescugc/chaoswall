define([
  "backbone",
], function(
  Backbone,
){
  return Backbone.View.extend({
    initialize: function() {
      this.$content = this.$("#content");
    },
    setContent: function(view) {
      var content = this.content;
      if (content) content.remove();
      content = this.content = view;
      this.$content.html(content.render().el);
    },
  });
});
