$(function() {
	$('.js-ping').click(function() {
		var el = $(this);
		var pongEl = el.siblings('.js-pong');
		(function() {
			var pinger = arguments.callee; 
			$.get('/ping', function(data) {
				pongEl.empty().append(data).effect("highlight", {color: '#0092D1'}, 1000);
		    	setTimeout(pinger, 3*1000);			  
			});
		})();		
	});
});
