define([
  "jquery", "underscore", "backbone",
], function(
  $, _, Backbone,
) {
  return Backbone.View.extend({
    template: _.template($("#route-tmpl").html()),
    events: {
      "click button.edit": "renderEditRoute",
      "click button.delete": "deleteRoute",
    },
    initialize: function(opt) {
      opt = opt || {}
      this.router = opt.router

      this.route = this.model
      this.listenTo(this.route, "change", this.render)

      this.route.fetch()
    },
    render: function() {
      this.$el.html(this.template({route: this.route.toJSON()}));
      return this;
    },
    renderEditRoute: function(e) {
      e.stopPropagation()
      this.router.navigate(this.route.url()+"/edit", { trigger: true })
    },
    deleteRoute: function(e) {
      e.stopPropagation()
      var that = this
      that.route.destroy({
        wait: true,
        success: function(model, response, options) {
          that.router.navigate(that.route.wall.url(),{ trigger: true })
        },
        error: function(model, response, options) {
          console.log(response.responseText)
        },
      })
    },
  });
});
