define([
  "jquery", "underscore", "backbone",
], function(
  $, _, Backbone,
) {
  return Backbone.View.extend({
    template: _.template($("#wall-tmpl").html()),
    events: {
      "click #new-route": "renderNewRoute",
      "click tr": "renderRoute",
      "click button.edit": "renderEditWall",
      "click button.delete": "deleteWall",
    },
    initialize: function(opt) {
      opt = opt || {}
      this.router = opt.router

      this.wall = this.model

      this.listenTo(this.wall, "change", this.render)
      this.listenTo(this.wall.routes, "reset", this.render)
      this.wall.fetch()
      this.wall.routes.fetch({reset: true})
    },
    render: function() {
      this.$el.html(this.template({wall: this.wall.toJSON(), routes: this.wall.routes.toJSON()}));
      return this;
    },
    renderNewRoute: function() {
      this.router.navigate(this.wall.routes.url()+"/new",{ trigger: true })
    },
    renderRoute: function(e) {
      var can = e.currentTarget.dataset.canonical
      this.router.navigate(this.wall.routes.url()+"/"+can,{ trigger: true })
    },
    renderEditWall: function(e) {
      e.stopPropagation()
      this.router.navigate(this.wall.url()+"/edit", { trigger: true })
    },
    deleteWall: function(e) {
      e.stopPropagation()
      var that = this
      that.wall.destroy({
        wait: true,
        success: function(model, response, options) {
          that.router.navigate(that.wall.gym.url(),{ trigger: true })
        },
        error: function(model, response, options) {
          console.log(response.responseText)
        },
      })
    },
  });
});
