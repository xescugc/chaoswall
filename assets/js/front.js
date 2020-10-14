var app = app || {};

$(function(){
  'use strict';

  // Utils

  // getFormData transforms the $form data into
  // a Object
  var getFormData = function($form) {
    var unindexed_array = $form.serializeArray();
    var indexed_array = {};

    $.map(unindexed_array, function(n, i){
      indexed_array[n["name"]] = n["value"];
    });

    return indexed_array;
  }

  // parseIfData returns the response.data
  // if the key data exists
  var parseIfData = function(response) {
    if (response === null || response === undefined) {
      return response
    }

    if ("data" in response) {
      return response.data
    }

    return response
  }

  // Models

  var Gym = Backbone.Model.extend({
    idAttribute: "canonical",
    urlRoot: "/gyms",
    parse: parseIfData,
    initialize: function() {
      this.walls = new WallsCollection();
      this.walls.gym = this;
    },
  })

  var Wall = Backbone.Model.extend({
    idAttribute: "canonical",
    parse: parseIfData,
    initialize: function() {
      this.routes = new RoutesCollection();
      this.routes.wall = this;
    },
  })

  var Route = Backbone.Model.extend({
    idAttribute: "canonical",
    parse: parseIfData,
  })

  // Collections

  var GymsCollection = Backbone.Collection.extend({
    url: "/gyms",
    model: Gym,
    parse: parseIfData,
  })

  var WallsCollection = Backbone.Collection.extend({
    model: Wall,
    parse: parseIfData,
    url: function() {
      return this.gym.url()+"/walls"
    }
  })

  var RoutesCollection = Backbone.Collection.extend({
    model: Route,
    parse: parseIfData,
    url: function() {
      return this.wall.url()+"/routes"
    }
  })

  // Views

  app.AppView = Backbone.View.extend({
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

  app.GymsView = Backbone.View.extend({
    template: _.template($("#gyms-tmpl").html()),
    collection: new GymsCollection(),
    events: {
      "click #new-gym": "renderNewGym",
      "click tr": "renderGym",
    },
    initialize: function() {
      this.listenTo(this.collection,"reset",this.render);
      this.collection.fetch({reset: true});
    },
    render: function() {
      this.$el.html(this.template({gyms: this.collection.toJSON()}));
      return this;
    },
    renderNewGym: function() {
      app.Router.navigate("gyms/new",{ trigger: true })
    },
    renderGym: function(e) {
      e.stopPropagation()
      var can = e.currentTarget.dataset.canonical
      app.Router.navigate("gyms/"+can,{ trigger: true })
    },
  });

  app.GymNewView = Backbone.View.extend({
    template: _.template($("#gym-new-tmpl").html()),
    events: {
      "submit form": "submitForm",
    },
    render: function() {
      this.$el.html(this.template());
      return this;
    },
    submitForm: function(e) {
      e.preventDefault()
      var data = getFormData(this.$el.find("form"))
      var g = new Gym(data)

      g.save(null, {
        wait: true,
        success: function(model, response, options) {
          app.Router.navigate(model.url(), { trigger: true })
        },
        error: function(model, response, options) {
          console.log(response.responseText)
        }
      });
    },
  });

  app.GymEditView = Backbone.View.extend({
    template: _.template($("#gym-edit-tmpl").html()),
    events: {
      "submit form": "submitForm",
    },
    initialize: function() {
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
      var data = getFormData(this.$el.find("form"))
      that.gym.set(data)

      that.gym.save(null, {
        wait: true,
        success: function(model, response, options) {
          that.gym = model
          app.Router.navigate(that.gym.url(), { trigger: true })
        },
        error: function(model, response, options) {
          console.log(response.responseText)
        }
      });
    },
  });

  app.GymView = Backbone.View.extend({
    template: _.template($("#gym-tmpl").html()),
    events: {
      "click #new-wall": "renderNewWall",
      "click tr": "renderWall",
      "click button.edit": "renderEditGym",
      "click button.delete": "deleteGym",
    },
    initialize: function() {
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
      app.Router.navigate(this.gym.walls.url()+"/new",{ trigger: true })
    },
    renderWall: function(e) {
      var can = e.currentTarget.dataset.canonical
      app.Router.navigate(this.gym.walls.url()+"/"+can,{ trigger: true })
    },
    renderEditGym: function(e) {
      e.stopPropagation()
      app.Router.navigate(this.gym.url()+"/edit", { trigger: true })
    },
    deleteGym: function(e) {
      e.stopPropagation()
      this.gym.destroy({
        wait: true,
        success: function(model, response, options) {
          app.Router.navigate("gyms/",{ trigger: true })
        },
        error: function(model, response, options) {
          console.log(response.responseText)
        },
      })
    },
  });

  app.WallNewView = Backbone.View.extend({
    template: _.template($("#wall-new-tmpl").html()),
    events: {
      "submit form": "submitForm",
    },
    render: function() {
      this.$el.html(this.template());
      return this;
    },
    submitForm: function(e) {
      var that = this
      e.preventDefault()
      var data = getFormData(that.$el.find("form"))
      that.gym.walls.create(data, {
        wait: true,
        success: function(model, response, options) {
          app.Router.navigate(that.gym.url(), { trigger: true })
        },
        error: function(model, response, options) {
          console.log(response.responseText)
        }
      });
    },
  });

  app.WallView = Backbone.View.extend({
    template: _.template($("#wall-tmpl").html()),
    events: {
      "click #new-route": "renderNewRoute",
      "click tr": "renderRoute",
    },
    initialize: function() {
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
      app.Router.navigate(this.wall.routes.url()+"/new",{ trigger: true })
    },
    renderRoute: function(e) {
      var can = e.currentTarget.dataset.canonical
      app.Router.navigate(this.wall.routes.url()+"/"+can,{ trigger: true })
    },
  });

  app.RouteNewView = Backbone.View.extend({
    template: _.template($("#route-new-tmpl").html()),
    events: {
      "submit form": "submitForm",
    },
    render: function() {
      this.$el.html(this.template());
      return this;
    },
    submitForm: function(e) {
      var that = this
      e.preventDefault()
      var data = getFormData(that.$el.find("form"))
      that.wall.routes.create(data, {
        wait: true,
        success: function(model, response, options) {
          app.Router.navigate(that.wall.url(), { trigger: true })
        },
        error: function(model, response, options) {
          console.log(response.responseText)
        }
      });
    },
  });

  app.RouteView = Backbone.View.extend({
    template: _.template($("#route-tmpl").html()),
    initialize: function() {
      this.route = this.model
      this.listenTo(this.route, "change", this.render)

      this.route.fetch()
    },
    render: function() {
      this.$el.html(this.template({route: this.route.toJSON()}));
      return this;
    },
  });

  // Router

  var Router = Backbone.Router.extend({
    routes: {
      "gyms":                           "gymsRender",
      "gyms/new":                       "gymNewRender",

      "gyms/:gym_canonical":            "gymRender",
      "gyms/:gym_canonical/edit":       "gymEditRender",
      "gyms/:gym_canonical/walls/new":  "wallNewRender",

      "gyms/:gym_canonical/walls/:wall_canonical":             "wallRender",
      "gyms/:gym_canonical/walls/:wall_canonical/routes/new":  "routeNewRender",

      "gyms/:gym_canonical/walls/:wall_canonical/routes/:route_canonical":  "routeRender",

      "*all":                 "homeRender",
    },
    initialize: function() {
      this.layout = new app.AppView({
        el: 'body',
      });
    },

    gymsRender: function() {
      var gsv = new app.GymsView();
      this.layout.setContent(gsv)
    },
    gymNewRender: function() {
      var gnv = new app.GymNewView();
      this.layout.setContent(gnv)
    },
    gymEditRender: function(gymCan) {
      var g = new Gym()
      g.set("canonical", gymCan)

      var gnv = new app.GymEditView({model: g});
      this.layout.setContent(gnv)
    },
    gymRender: function(gymCan) {
      var g = new Gym()
      g.set("canonical", gymCan)
      var gv = new app.GymView({model: g})
      this.layout.setContent(gv)
    },

    wallNewRender: function(gymCan) {
      var g = new Gym()
      g.set("canonical", gymCan)
      var wnv = new app.WallNewView();
      wnv.gym = g
      this.layout.setContent(wnv)
    },
    wallRender: function(gymCan, wallCan) {
      var g = new Gym()
      g.set("canonical", gymCan)

      var w = new Wall()
      w.set("canonical", wallCan)
      w.urlRoot = g.walls.url()

      var wnv = new app.WallView({model: w});
      this.layout.setContent(wnv)
    },

    routeNewRender: function(gymCan, wallCan) {
      var g = new Gym()
      g.set("canonical", gymCan)

      var w = new Wall()
      w.set("canonical", wallCan)
      w.urlRoot = g.walls.url()

      var rnv = new app.RouteNewView();
      rnv.wall = w
      this.layout.setContent(rnv)
    },

    routeRender: function(gymCan, wallCan, routeCan) {
      var g = new Gym()
      g.set("canonical", gymCan)

      var w = new Wall()
      w.set("canonical", wallCan)
      w.urlRoot = g.walls.url()

      var r = new Route()
      r.set("canonical", routeCan)
      r.urlRoot = w.routes.url()

      var rv = new app.RouteView({model: r});
      this.layout.setContent(rv)
    },

    homeRender: function() {
      // For now as we do not have any Home page
      // we'll just redirect to the /gyms
      this.navigate("gyms",{ trigger: true })
    },
  });

  Backbone.ajax = function(request) {
    request = _({ contentType: 'application/json' }).defaults(request);
    return Backbone.$.ajax.call(Backbone.$, request);
  };

  app.Router = new Router();
  Backbone.history.start({pushState: true});

})
