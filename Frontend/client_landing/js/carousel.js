let carousel = (function()
{
	let carousel = 
	{
		imgs: [],				//Адреса картинок
		current_state: 0,
		root: $("#carousel_root"),  //Длинная плашка внутри overflow: hidden
		set_state: function(state)
		{
			this.current_state = state;
			this.render_state(this.root, this.current_state);
		},
		next: function()
		{
			let state = this.current_state;
			if(state >= 5)
			{
				state = 0;
			} else 
			{
				state += 1;
			}
			this.set_state(state);
		},
		prev: function()
		{
			let state = this.current_state;
			if(state <= 0)
			{
				state = 5;
			} else 
			{
				state -= 1;
			}
			this.set_state(state);
		},
		render_state: function(root, state)
		{
			root.attr("current_state", state);
			$(`[state]`).removeClass('active');
			$(`[state = '${state}']`).addClass("active");
		},
		follow_points: function()
		{
			let state = $(this).attr("state");
			carousel.set_state.bind(carousel, state);
		}
	}

	$("#arrow_left").on("click", carousel.prev.bind(carousel));
	$("#arrow_right").on("click", carousel.next.bind(carousel));
	$(".point").on("click", carousel.follow_points);
	return carousel;
})()