define({
  // getFormData transforms the $form data into
  // a Object
  getFormData: function($form) {
    var unindexed_array = $form.serializeArray();
    var indexed_array = {};

    $.map(unindexed_array, function(n, i){
      indexed_array[n["name"]] = n["value"];
    });

    return indexed_array;
  },

  // parseIfData returns the response.data
  // if the key data exists
  parseIfData: function(response) {
    if (response === null || response === undefined) {
      return response
    }

    if ("data" in response) {
      return response.data
    }

    return response
  },

  // getPointDistance will calculate the distance
  // between x,y and xx,yy
  getPointDistance: function(x,y,xx,yy){
    var fx = x-xx
    var fy = y-yy

    return Math.sqrt((fx*fx) + (fy*fy))
  },
})


