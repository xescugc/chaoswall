define([
  "jquery", "underscore", "backbone",
  "utils",
], function(
  $, _, Backbone,
  Utils,
) {
  return Backbone.View.extend({
    template: _.template($("#wall-edit-tmpl").html()),
    events: {
      "submit form": "submitForm",
      "change #image": "changePreview",
    },
    initialize: function(opt) {
      opt = opt || {}
      this.router = opt.router

      this.wall = this.model
      this.listenTo(this.wall, "change", this.render)

      this.wall.fetch()
    },
    render: function() {
      this.$el.html(this.template({wall: this.wall.toJSON()}));
      this.beforePreview = this.$el.find("#before-preview")
      this.afterPreview = this.$el.find("#after-preview")

      var pImage = new Image();
      pImage.style.width = "100%";
      pImage.src = this.wall.get("image")
      this.beforePreview.html(pImage)

      return this;
    },
    submitForm: function(e) {
      e.preventDefault()
      var that = this
      var data = Utils.getFormData(that.$el.find("form"))
      var file = that.$el.find('input[name="image"]')[0].files[0]; 
      var reader = new FileReader();
      that.wall.set(data)
      if (file === undefined) {
        that.wall.save(null, {
          wait: true,
          success: function(model, response, options) {
            model.gym = that.gym
            that.router.navigate(model.url(), { trigger: true })
          },
          error: function(model, response, options) {
            console.log(response.responseText)
          }
        });
      } else {
        reader.onloadend = function() {
          data.image = reader.result;
          that.wall.save(null, {
            wait: true,
            success: function(model, response, options) {
              model.gym = that.gym
              that.router.navigate(model.url(), { trigger: true })
            },
            error: function(model, response, options) {
              console.log(response.responseText)
            }
          });
        }

        reader.readAsDataURL(file)
      }
    },
    changePreview: function(e) {
      e.preventDefault()
      var that = this
      var file = that.$el.find('input[name="image"]')[0].files[0]; 
      var reader = new FileReader();
      reader.onloadend = function() {
        var pImage = new Image();
        pImage.style.width = "100%";
        pImage.title = file.name;
        pImage.src = reader.result
        that.beforePreview.html(pImage)
        $.ajax({
          url: that.wall.gym.url()+"/walls-image",
          data: JSON.stringify({ image: reader.result }),
          type: "POST",
          contentType: 'application/json',
          processData: false,
          success: function(data) {
            var vpImage = new Image();
            vpImage.style.width = "100%";
            vpImage.title = file.name;
            vpImage.src = data.data.image;
            that.afterPreview.html(vpImage)
          },
          error: function(data) {
            console.log("error", data)
          },
        })
      }
      reader.readAsDataURL(file)
    },
  });
});
