

$(document).on("click", "[data-module-id]", router.wrapper(router.change_module.bind(router)))
$(document).on("click", "[data-product-id]", router.wrapper(router.go_to_product.bind(router)))
$(document).on("click", "[data-organisation-id]", router.wrapper(router.go_to_organisation.bind(router)))
$(document).on("click", ".goodCard-closeBtn", router.wrapper(router.close_popup.bind(router)))