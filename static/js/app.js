;(function (w, d) {
  'use strict';

  var App = (function () {

    var _init = function () {
      console.log('init app');
    };

    return {
      init: _init
    }
  })();

  App.init();

})(window, document);