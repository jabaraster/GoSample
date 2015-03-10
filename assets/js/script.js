(function($) {
"use strict";

var initialize = function() {
    $('button').click(function() {
        alert(this);
    });
};

$(function() { initialize(); });

})(jQuery);