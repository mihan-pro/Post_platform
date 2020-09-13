let router = (function()
{
    let router =
    {
        history: [],
        root: $(".content"),
        change_module: function(th)
        {
            let module_id = $(th).attr("data-module-id");
            this.get_module.call(this, module_id);
        },
        close_popup: function()
        {
            $("#goodCard-curtain").addClass("d-none");
            basket.render_bascket_items.call(basket)
        },
        show_popup: function()
        {
            $("#goodCard-curtain").removeClass("d-none");
            basket.render_bascket_items.call(basket)
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
           // let products_html = this.get_products_html([data]);
            let products_html = this.create_product_card(data);
            this.set_html_to_popup.call(this, products_html);
            this.show_popup();
        },
        set_html_to_popup: function(html)
        {
            $("#goodCard-curtain").html(html)
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
            let organisations_html = ""; // this.get_organisations_html(data.Organisations);
            this.set_html_to_root.call(this, products_html + organisations_html);
            this.show_subheader();
            $("#subheader-title").html(data.Module.module_name);
        },
        show_subheader: function()
        {
            $("#subheader").removeClass("d-none");
        },
        hide_subheader: function()
        {
            $("#subheader").addClass("d-none");
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
                        <div class="content-item" data-wrapper-id = "${item.product_id}" itemstate="1">
                            <div class="content_itemImg" data-product-id = "${item.product_id}" style = "background-image: url(${item.product_image})">
                                <div class="content-check"></div>
                            </div>
                            <p class="content-titleItem">${item.product_name}</p>
                            <div class="content-itemDelivery">
                                <div class="content-deliveryTime" data-organisation-id = "${item.organisation_id}">
                                    <div class="content-deliveryTimeText">${item.organisation_name}</div>
                                </div>
                                <div class="content-deliveryMark">
                                    <div class="content-deliveryStar"></div>
                                    <div class="content-deliveryText">${item.product_rating}"</div>
                                </div>
                            </div>
                            <div class="content-buyitem">
                                <div class="content-costItem">
                                    <span class="content-itemSumm">${item.product_price}р</span>/1<span>шт</span>
                                </div>
                                <div class="content-buyBtn" data-basket-id = "${item.product_id}">В корзину</div>
                                <div class="content-addTakeOne">
                                    <div class="content-addOneMinus" data-basket-minus-id = "${item.product_id}"></div>
                                    <div class="content-addOneValue" data-basket-value-id = "${item.product_id}">1</div>
                                    <div class="content-addOnePlus" data-basket-plus-id = "${item.product_id}"></div>
                                </div>
                            </div>
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
            this.hide_subheader();
        },
        create_main_page_str(data)
        {
            let res = data.map(item =>
            {
                return `
                <div class="module-item" data-module-id = "${item.module_id}">
                    <div class="module_itemImg" style = "background: url(${item.module_icon})">
                    </div>
                    <p class="module-titleItem">${item.module_name}</p>
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
        },
        buy_from_list: function(th)
        {
            let id = $(th).attr("data-basket-id");
            $(`[data-wrapper-id = "${id}"]`).attr("itemstate", 2);
            basket.add_product(id);
        },
        plus_button: function(th)
        {
            let id = $(th).attr("data-basket-plus-id");
            basket.add_product(id);
        },
        minus_button: function(th)
        {
            let id = $(th).attr("data-basket-minus-id");
            basket.minus_product(id);
        },
        create_product_card: function(item)
        {
            let bascket_items = (store.get("bascket_items") || {})[item.product_id] || 0;
            return `
            <div class="goodCard">
                <div class="goodCard-image" style = "background-image: url(${item.product_image})"></div>
                <div class="goodCard-info">
                    <div class="goodCard-title">${item.product_name}</div>
                    <div class="goodCard-firmInfo">
                        <div class="goodCard-firmName">${item.organisation_name}</div>
                        <div class="goodCard-firmMark">
                            <div class="content-deliveryStar"></div>
                            <div class="goodCard-markValue">${item.product_rating}</div>
                        </div>
                        <div class="goodCard-reports">
                            <div class="goodCard-reportIco"></div>
                            <div class="goodCard-reportValue">1</div>
                        </div>
                    </div>
                    <div class="goodCard-description">
                        • Пункт описания 1<br>
                        • Пункт описания 2<br>
                        • Пункт описания 3 <br>
                        • Пункт описания 4<br>
                        • Пункт описания 5
                    </div>
                    <div class="goodCard-costInfo">
                        <div class="goodCard-price">
                            <span class="content-itemSumm">${item.product_price}р</span>/1<span>шт</span>
                        </div>
                        <div class="goodCard-count">
                            <div class="goodCard-countText">Кол-во</div>
                            <div class="content-addTakeOne">
                            <div class="content-addOneMinus" data-basket-minus-id = "${item.product_id}"></div>
                            <div class="content-addOneValue" data-basket-value-id = "${item.product_id}">${bascket_items}</div>
                            <div class="content-addOnePlus" data-basket-plus-id = "${item.product_id}"></div>
                            </div>
                        </div>
                    </div>
                    <div class="goodCard-buyBlock">
                        <div class="goodCard-buyBtn">В корзину</div>
                        <div class="goodCard-deliveryInfo">
                            <div class="goodCard-deliveryText">Условия доставки</div>
                            <div class="goodCard-infoIco"></div>
                        </div>
                    </div>
                    <div class="goodCard-closeBtn"></div>
                </div>
            </div>            
            `
        }
    }

    
    return router;
})()