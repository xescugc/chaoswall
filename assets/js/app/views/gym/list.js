define([
  "jquery", "underscore", "backbone",
  "collections/gyms",
], function(
  $, _, Backbone,
  GymsCollection,
) {
  return Backbone.View.extend({
    template: _.template($("#gyms-tmpl").html()),
    collection: new GymsCollection(),
    events: {
      "click #new-gym": "renderNewGym",
      "click tr": "renderGym",
    },
    initialize: function(opt) {
      opt = opt || {}
      this.router = opt.router

      this.listenTo(this.collection,"reset",this.render);
      this.collection.fetch({reset: true});
    },
    render: function() {
      this.$el.html(this.template({gyms: this.collection.toJSON()}));
      return this;
    },
    renderNewGym: function() {
      this.router.navigate("gyms/new",{ trigger: true })
    },
    renderGym: function(e) {
      e.stopPropagation()
      var can = e.currentTarget.dataset.canonical
      this.router.navigate("gyms/"+can,{ trigger: true })
    },
  });
});
