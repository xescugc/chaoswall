define([
  "jquery", "underscore", "backbone",
  "utils",
], function(
  $, _, Backbone,
  Utils,
) {
  return Backbone.View.extend({
    template: _.template($("#route-new-tmpl").html()),
    events: {
      "submit form": "submitForm",
      "click canvas#canvas" : "clickCanvas",
    },
    initialize: function(opt) {
      opt = opt || {}
      this.wall = opt.wall
      this.router = opt.router

      this.listenTo(this.wall, "change", this.render)
      this.listenTo(this.wall, "change:hold", this.drawHold)

      this.wall.fetch()
    },
    render: function() {
      this.$el.html(this.template());

      this.canvas = this.$el.find("#canvas")[0]
      var ctx = this.canvas.getContext("2d");

      var that = this;

      var img = new Image();
      img.src = this.wall.get("image")
      img.onload = function() {
        w = img.width;
        h = img.height;
        that.canvas.width = w;
        that.canvas.height = h;
        ctx.drawImage(img, 0, 0);

        //ctx.getImageData(0, 0, w, h);    // some browsers synchronously decode image here
        _.each(that.wall.get("holds"), function(h) {
          that.drawHold(h)
        })
      }

      return this;
    },
    drawHold: function(h) {
      var ctx = this.canvas.getContext("2d");
      ctx.lineWidth = "2";
      if (h.selected) {
        // Green
        ctx.strokeStyle = 'rgb(52, 240, 52)';
      } else {
        // Red
        ctx.strokeStyle = 'rgb(240, 52, 52)';
      }
      // The X and Y are start of the drawing
      ctx.strokeRect(h.x-(h.size/2), h.y-(h.size/2), h.size, h.size);
    },
    clickCanvas: function(e) {
      var rect = this.canvas.getBoundingClientRect();
      var mouseXPos = (e.originalEvent.x - rect.left);
      var mouseYPos = (e.originalEvent.y - rect.top);

      var hold;
      var distance
      _.each(this.wall.get("holds"), function(h) {
        var d = Utils.getPointDistance(h.x, h.y, mouseXPos,mouseYPos)
        // For the first run we set the first distance
        if (distance === undefined) {
          distance = d
          hold = h
        } else if (d < distance) {
          distance = d
          hold = h
        }
      })
      // Trigger an specific event for the wall model
      // that will be listen too by the main function
      // and will just draw again the hold with the 
      // new state
      hold.selected = !hold.selected
      this.wall.trigger("change:hold", hold)
    },
    submitForm: function(e) {
      e.preventDefault()
      var that = this
      var data = Utils.getFormData(that.$el.find("form"))
      that.wall.routes.create(data, {
        wait: true,
        success: function(model, response, options) {
          model.wall = that.wall
          that.router.navigate(model.url(), { trigger: true })
        },
        error: function(model, response, options) {
          console.log(response.responseText)
        }
      });
    },
  });
});
