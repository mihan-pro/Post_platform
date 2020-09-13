let router = (function()
{
    let router =
    {
        history: [],
        root: $("#wrapper"),
        change_module: function(th)
        {
            let module_id = $(th).attr("data-module-id");
            this.get_module.call(this, module_id);
        },
        close_popup: function()
        {
            $("#goodCard-curtain").addClass("d-none");
        },
        go_back: function()
        {
            
        },
        get_module: function(module_id)
        {
            loader.send({}, `https://313app.na4u.ru/product/feed/` + module_id, this.render_module.bind(this))
        },
        get_product: function(product_id)
        {
            loader.send({}, `https://313app.na4u.ru/product/` + product_id, this.render_product_single.bind(this))
        },
        get_organisation: function(organisation_id)
        {
            loader.send({}, `https://313app.na4u.ru/organisation/` + organisation_id, this.render_organisation_single.bind(this))
        },
        render_product_single: function(data)
        {
            let products_html = this.get_products_html([data]);
            this.set_html_to_root.call(this, products_html);
        },
        render_organisation_single: function(data)
        {
            let organisation_html = this.get_organisations_html([data]);
            this.set_html_to_root.call(this, organisation_html);
        },
        go_to_product: function(th)
        {
            let product_id = $(th).attr("data-product-id");
            this.get_product.call(this, product_id);
        },
        go_to_organisation: function(th)
        {
            let organisation_id = $(th).attr("data-organisation-id");
            this.get_organisation.call(this, organisation_id);
        },
        render_module: function(data)
        {
            let products_html = this.get_products_html(data.Products);
            let organisations_html = this.get_organisations_html(data.Organisations);

            this.set_html_to_root.call(this, products_html + organisations_html);
        },
        get_organisations_html: function(data)
        {
            //Организации в модуле
            let res = data.map(item =>
                {
                    return `
                        <div class = "module_wrapper" data-organisation-id = "${item.organisation_id}" style = "background: lightgreen; height: 150px; width: 1000px; border: solid 1px black;">
                            <img style = "display: inline-block" src = "${item.organisation_logo}">
                            <div style = "display: inline-block">${item.organisation_name}</div>
                            <div style = "display: inline-block">${item.organisation_description}</div>
                            <div style = "display: inline-block">${item.organisation_address}</div>
                            <div style = "display: inline-block">${item.organisation_rating}</div>
                        </div>               
                    `
                }).join("");
            
            return res;
        },
        get_products_html(data)
        {
            //Продукты в модуле
            let res = data.map(item =>
                {
                    return `
                        <div class = "module_wrapper" data-product-id = "${item.product_id}" style = "background: lightgray; height: 150px; width: 1000px; border: solid 1px black;">
                            <img style = "display: inline-block" src = "${item.product_icon}">
                            <div style = "display: inline-block">${item.product_name}</div>
                            <div style = "display: inline-block">${item.product_description}</div>
                            <div data-organisation-id = "${item.organisation_id}" style = "display: inline-block">${item.organisation_name}</div>
                        </div>               
                    `
                }).join("");

            return res;
        },
        get_default_page: function()
        {
            loader.send({}, "https://313app.na4u.ru/module", this.set_default_page.bind(this));
        },
        set_html_to_root: function(html)
        {
            this.root.html(html);
        },
        set_default_page: function(data)
        {
            let main_page_html = this.create_main_page_str(data);
            this.set_html_to_root.call(this, main_page_html);
        },
        create_main_page_str(data)
        {
            let res = data.map(item =>
            {
                return `
                    <div class = "module_wrapper" data-module-id = "${item.module_id}" style = "height: 150px; width: 1000px; border: solid 1px black;">
                        <img style = "display: inline-block" src = "${item.module_icon}">
                        <div style = "display: inline-block">${item.module_name}</div>
                        <div style = "display: inline-block">${item.module_description}</div>
                    </div>               
                `
            }).join("");

            return res
        },
        wrapper: function (f) 
        {
            return function () 
            {
              f(this);
            };
        }
    }

    
    return router;
})()