define([
  "jquery", "underscore", "backbone",

  "views/app/view",

  "views/gym/list", "views/gym/new", "views/gym/show", "views/gym/edit",
  "views/wall/new", "views/wall/show", "views/wall/edit",
  "views/route/new", "views/route/show", "views/route/edit",

  "models/gym", "models/wall", "models/route",
], function(
  $, _, Backbone,

  AppView,

  GymsView, GymNewView, GymView, GymEditView,
  WallNewView, WallView, WallEditView,
  RouteNewView, RouteView, RouteEditView,

  Gym, Wall, Route,
) {
  var Router = Backbone.Router.extend({
    routes: {
      "gyms":                           "gymsRender",
      "gyms/new":                       "gymNewRender",

      "gyms/:gym_canonical":            "gymRender",
      "gyms/:gym_canonical/edit":       "gymEditRender",
      "gyms/:gym_canonical/walls/new":  "wallNewRender",

      "gyms/:gym_canonical/walls/:wall_canonical":             "wallRender",
      "gyms/:gym_canonical/walls/:wall_canonical/edit":        "wallEditRender",
      "gyms/:gym_canonical/walls/:wall_canonical/routes/new":  "routeNewRender",

      "gyms/:gym_canonical/walls/:wall_canonical/routes/:route_canonical":      "routeRender",
      "gyms/:gym_canonical/walls/:wall_canonical/routes/:route_canonical/edit": "routeEditRender",

      "*all":                 "homeRender",
    },
    initialize: function() {
      this.layout = new AppView({
        el: 'body',
      });
    },

    gymsRender: function() {
      var gsv = new GymsView({router: this});
      this.layout.setContent(gsv)
    },
    gymNewRender: function() {
      var gnv = new GymNewView({
        router: this,
      });
      this.layout.setContent(gnv)
    },
    gymEditRender: function(gymCan) {
      var g = new Gym()
      g.set("canonical", gymCan)

      var gev = new GymEditView({model: g, router: this});
      this.layout.setContent(gev)
    },
    gymRender: function(gymCan) {
      var g = new Gym()
      g.set("canonical", gymCan)
      var gv = new GymView({model: g, router: this})
      this.layout.setContent(gv)
    },

    wallNewRender: function(gymCan) {
      var g = new Gym()
      g.set("canonical", gymCan)

      var wnv = new WallNewView({router: this});
      wnv.gym = g

      this.layout.setContent(wnv)
    },
    wallEditRender: function(gymCan,wallCan) {
      var g = new Gym()
      g.set("canonical", gymCan)

      var w = new Wall()
      w.set("canonical", wallCan)
      w.urlRoot = g.walls.url()
      w.gym = g

      var wev = new WallEditView({model: w, router: this});
      this.layout.setContent(wev)
    },
    wallRender: function(gymCan, wallCan) {
      var g = new Gym()
      g.set("canonical", gymCan)

      var w = new Wall()
      w.set("canonical", wallCan)
      w.urlRoot = g.walls.url()
      w.gym = g

      var wv = new WallView({model: w, router: this});
      this.layout.setContent(wv)
    },

    routeNewRender: function(gymCan, wallCan) {
      var g = new Gym()
      g.set("canonical", gymCan)

      var w = new Wall()
      w.set("canonical", wallCan)
      w.urlRoot = g.walls.url()

      var rnv = new RouteNewView({ wall: w, router: this});
      rnv.wall = w

      this.layout.setContent(rnv)
    },
    routeEditRender: function(gymCan, wallCan, routeCan) {
      var g = new Gym()
      g.set("canonical", gymCan)

      var w = new Wall()
      w.set("canonical", wallCan)
      w.urlRoot = g.walls.url()

      var r = new Route()
      r.set("canonical", routeCan)
      r.urlRoot = w.routes.url()

      var rev = new RouteEditView({model: r, router: this});
      this.layout.setContent(rev)
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
      r.wall = w

      var rv = new RouteView({model: r, router: this});
      this.layout.setContent(rv)
    },

    homeRender: function() {
      // For now as we do not have any Home page
      // we'll just redirect to the /gyms
      this.navigate("gyms",{ trigger: true })
    },
  });

  return new Router();
});
