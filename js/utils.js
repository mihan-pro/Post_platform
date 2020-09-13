let utils = (function()
{
    let utils = 
    {
        parse_url: function(url)
        {
            let search = url.split("?")[1].split("&");
            let obj = {};
            search.forEach(item =>
            {
                let s = item.split("=");
                obj[s[0]] = s[1];
            })

            return obj
        },
        copyObj: function(obj)
        {
            return JSON.parse(JSON.stringify(obj));
        }
    };

    return utils;
})()