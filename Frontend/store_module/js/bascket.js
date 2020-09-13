let basket = (function()
{
    let basket = 
    {
        add_product: function(id)
        {
            let count = Number(store.get("bascket_counter")) || 0;
            let bascket_items = store.get("bascket_items") || {};
            bascket_items[id] = bascket_items[id] || 0;
            bascket_items[id] += 1;
            count += 1;
            store.set({bascket_counter: count, bascket_items: bascket_items});
        },
        minus_product: function(id)
        {
            let count = Number(store.get("bascket_counter")) || 0;
            let bascket_items = store.get("bascket_items") || {};
            bascket_items[id] = bascket_items[id] || 0;
            bascket_items[id] -= 1;
            count -= 1;
            if(bascket_items[id] <= 0)
            {
                bascket_items[id] = 0;
                $(`[data-wrapper-id = "${id}"]`).attr("itemstate", 1);
            } 
            if(count <= 0) count = 0;
            store.set({bascket_counter: count, bascket_items: bascket_items});
        },
        render_bascket_items: function(data)
        {
            for(let item in data)
            {
                $(`[data-basket-value-id = "${item}"]`).html(data[item]);
            }
        },
        render_backet_count: function(num)
        {
            $("#header-shopcartIndicator").html(num);
            if(num <= 0)
            {
                $(`#basket`).attr("data-havegoods", "0");
            } else 
            {
                $(`#basket`).attr("data-havegoods", 1);
            }
        }
    }

    //basket.render_backet_count(store.get("bascket_counter"));
    return basket;
})()