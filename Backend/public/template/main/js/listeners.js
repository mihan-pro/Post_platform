

$(document).on("click", "[data-module-id]", router.wrapper(router.change_module.bind(router)))
$(document).on("click", "[data-product-id]", router.wrapper(router.go_to_product.bind(router)))
$(document).on("click", "[data-organisation-id]", router.wrapper(router.go_to_organisation.bind(router)))
$(document).on("click", ".goodCard-closeBtn", router.wrapper(router.close_popup.bind(router)))
$(document).on("click", ".content-buyBtn", router.wrapper(router.buy_from_list.bind(router)))
$(document).on("click", "[data-basket-plus-id]", router.wrapper(router.plus_button.bind(router)))
$(document).on("click", "[data-basket-minus-id]", router.wrapper(router.minus_button.bind(router)))
$(document).on("click", "#goodCard-curtain .goodCard-buyBtn", router.wrapper(router.close_popup.bind(router)))

//события стора
store.addListener("bascket_counter", basket.render_backet_count.bind(basket))
store.addListener("bascket_items", basket.render_bascket_items.bind(basket))
