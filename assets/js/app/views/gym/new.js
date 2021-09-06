define([
  "jquery", "underscore", "backbone",
  "utils",
  "models/gym",
], function(
  $, _, Backbone,
  Utils,
  Gym,
) {
  return Backbone.View.extend({
    template: _.template($("#gym-new-tmpl").html()),
    events: {
      "submit form": "submitForm",
    },
    initialize: function(opt) {
      opt = opt || {}
      this.router = opt.router
    },
    render: function() {
      this.$el.html(this.template());
      return this;
    },
    submitForm: function(e) {
      e.preventDefault()
      var data = Utils.getFormData(this.$el.find("form"))
      var g = new Gym(data)

      var that = this
      g.save(null, {
        wait: true,
        success: function(model, response, options) {
          that.router.navigate(model.url(), { trigger: true })
        },
        error: function(model, response, options) {
          console.log(response.responseText)
        }
      });
    },
  });
});
