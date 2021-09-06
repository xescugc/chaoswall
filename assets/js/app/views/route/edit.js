define([
  "jquery", "underscore", "backbone",
  "utils",
], function(
  $, _, Backbone,
  Utils,
) {
  return Backbone.View.extend({
    template: _.template($("#route-edit-tmpl").html()),
    events: {
      "submit form": "submitForm",
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
    submitForm: function(e) {
      e.preventDefault()
      var that = this
      var data = Utils.getFormData(that.$el.find("form"))
      that.route.set(data)

      that.route.save(null, {
        wait: true,
        success: function(model, response, options) {
          that.route = model
          that.router.navigate(that.route.url(), { trigger: true })
        },
        error: function(model, response, options) {
          console.log(response.responseText)
        }
      });
    },
  });
});
