let loader = (function()
{
    let loader = 
    {
        send: function (args, url, callbackSuccess, callbackError, argsSuccess, argsError) 
        {
			$.ajax(
            {
				url: url,
				crossDomain: true,
				dataType: "json",
				type: "GET",
				data: args,
				async: true,
				cashe: false,
                success: function success(data) 
                {
					callbackSuccess(data, argsSuccess);
				},
                error: function error(data) 
                {
					alert("error: ", error);
				}
			});
		}
    };

    return loader;
})()