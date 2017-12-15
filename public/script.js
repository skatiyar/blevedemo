(function($) {
  var searchBtn = $('#search-button');
  var results = $('.Search__Results');

  searchBtn.on('click', function(e) {
    var query = encodeURIComponent($('#search-input').val());

    $.get('/search?q=' + query)
    .then(function(res) {
      var data = JSON.parse(res);
      var resultsHtml = '';
      for (var i = 0; i < data.length; i++) {
        resultsHtml += '<div class="Result"><div class="Id">Doc Id: ' + data[i].id + '</div>';
        var frags = Object.keys(data[i].fragments);
        for (var j = 0; j < frags.length; j++) {
          resultsHtml += '<div class="Field">Doc Field: ' + frags[j] + '</div>';
          for (var x = 0; x < data[i].fragments[frags[j]].length; x++) {
            resultsHtml += '<div class="Frag">' + data[i].fragments[frags[j]][x] + '</div>';
          }
        }
        resultsHtml += '</div>';
      }
      results.empty();
      results.append(resultsHtml);
    })
    .catch(function(err) {
      alert(JSON.stringify(err));
    })
  })
})(window.jQuery)
