(function()
{
    let args = window.location.search ? utils.parse_url(window.location.search) : {};
    if(!args.module)
    {
        router.get_default_page.call(router);
    }
})()