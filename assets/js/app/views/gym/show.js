define([
  "jquery", "underscore", "backbone",
], function(
  $, _, Backbone,
) {
  return Backbone.View.extend({
    template: _.template($("#gym-tmpl").html()),
    events: {
      "click #new-wall": "renderNewWall",
      "click tr": "renderWall",
      "click button.edit": "renderEditGym",
      "click button.delete": "deleteGym",
    },
    initialize: function(opt) {
      opt = opt || {}
      this.router = opt.router

      this.gym = this.model

      this.listenTo(this.gym, "change", this.render)
      this.listenTo(this.gym.walls, "reset", this.render)
      this.gym.fetch()
      this.gym.walls.fetch({reset: true})
    },
    render: function() {
      this.$el.html(this.template({gym: this.gym.toJSON(), walls: this.gym.walls.toJSON()}));
      return this;
    },
    renderNewWall: function() {
      this.router.navigate(this.gym.walls.url()+"/new",{ trigger: true })
    },
    renderWall: function(e) {
      var can = e.currentTarget.dataset.canonical
      this.router.navigate(this.gym.walls.url()+"/"+can,{ trigger: true })
    },
    renderEditGym: function(e) {
      e.stopPropagation()
      this.router.navigate(this.gym.url()+"/edit", { trigger: true })
    },
    deleteGym: function(e) {
      e.stopPropagation()
      var that = this
      that.gym.destroy({
        wait: true,
        success: function(model, response, options) {
          that.router.navigate("gyms/",{ trigger: true })
        },
        error: function(model, response, options) {
          console.log(response.responseText)
        },
      })
    },
  });
});
