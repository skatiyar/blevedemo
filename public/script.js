(function($) {
  var searchBtn = $('#search-button');

  searchBtn.on('click', function(e) {
    var query = encodeURIComponent($('#search-input').val());

    $.get('/search?q=' + query)
    .then(function(res) {
      console.log(res);
    })
    .catch(function(err) {
      alert(JSON.stringify(err));
    })
  })
})(window.jQuery)
