define([
  "jquery", "underscore", "backbone",
  "utils",
], function(
  $, _, Backbone,
  Utils,
) {
  return Backbone.View.extend({
    template: _.template($("#gym-edit-tmpl").html()),
    events: {
      "submit form": "submitForm",
    },
    initialize: function(opt) {
      opt = opt || {}
      this.router = opt.router

      this.gym = this.model
      this.listenTo(this.gym, "change", this.render)

      this.gym.fetch()
    },
    render: function() {
      this.$el.html(this.template({gym: this.gym.toJSON()}));
      return this;
    },
    submitForm: function(e) {
      e.preventDefault()
      var that = this
      var data = Utils.getFormData(that.$el.find("form"))
      that.gym.set(data)

      that.gym.save(null, {
        wait: true,
        success: function(model, response, options) {
          that.gym = model
          that.router.navigate(that.gym.url(), { trigger: true })
        },
        error: function(model, response, options) {
          console.log(response.responseText)
        }
      });
    },
  });
});
